package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/cmd/handler"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/blockchain"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/httpin"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/validator"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/config"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/platform/postgres"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const shutdownTimeout = 10 * time.Second

var (
	configPath *string
)

func init() {
	configPath = flag.String("config", "config/config.yaml", "config")

	validator.InitValidator()
	httpin.InitHTTPIn()
}

func main() {
	flag.Parse()

	// shutdownSignal will be used to handle graceful shutdown
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	// loading config file
	cfg, err := config.Load(*configPath, shutdownSignal)
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// initializing logger
	logLevel := slog.LevelInfo
	if cfg.Logger.Debug {
		logLevel = slog.LevelDebug
	}

	opts := &slog.HandlerOptions{
		AddSource: cfg.Logger.AddSource,
		Level:     logLevel,
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, opts)))

	// initializing marketplace Service
	marketplaceService := marketplace.NewService(marketplace.NewPostgresRepository(postgres.New(&cfg.Postgres)))

	server := handler.InitRoutes(cfg, marketplaceService)

	slog.Info(fmt.Sprintf("Listening at port %s", server.Addr))
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http server error", slog.String("error", err.Error()))
		}

		shutdownSignal <- syscall.SIGTERM
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(cfg.Marketplace.Clients) > 0 {
		clients := make(map[int64]*blockchain.ChainClient)

		for chainID, c := range cfg.Marketplace.Clients {
			if !c.Enabled {
				continue
			}

			ethClient, err := ethclient.DialContext(ctx, c.WssURL)
			if err != nil {
				slog.ErrorContext(ctx, "failed to connect to eth client", slog.String("error", err.Error()), slog.Int64("chain_id", chainID), slog.String("collection_factory_address", c.CollectionFactoryAddress))
				os.Exit(1)
			}

			clients[chainID] = blockchain.NewChainClient(ethClient, c.CollectionFactoryAddress)
		}

		if len(clients) > 0 {
			eventListener, err := blockchain.NewChainEventsListener(ctx, clients, marketplaceService)
			if err != nil {
				slog.ErrorContext(ctx, "failed to create new blockchain events listener", slog.String("error", err.Error()))
				os.Exit(1)
			}

			go func() {
				if err := eventListener.Start(ctx); err != nil {
					slog.Error("failed to start event listener", slog.String("error", err.Error()))
				}

				shutdownSignal <- syscall.SIGTERM
			}()
		}
	}

	// Handle graceful shutdown
	<-shutdownSignal

	shutdown(server)
}

func shutdown(server *http.Server) {
	slog.Info("shutting down app")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Shutdown http server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server shutdown error", slog.String("error", err.Error()))
	}
}

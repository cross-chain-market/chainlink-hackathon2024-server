package blockchain

import (
	"context"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/contracts"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"regexp"
)

type ChainEventsListener struct {
	clientsByChainID   map[int64]*ChainClient
	marketplaceService *marketplace.Service
}

type ChainClient struct {
	client                   *ethclient.Client
	collectionFactoryAddress string
}

func NewChainClient(client *ethclient.Client, collectionFactoryAddress string) *ChainClient {
	return &ChainClient{
		client:                   client,
		collectionFactoryAddress: collectionFactoryAddress,
	}
}

func NewChainEventsListener(ctx context.Context, clientsByChainID map[int64]*ChainClient, marketplaceService *marketplace.Service) (*ChainEventsListener, error) {
	for _, client := range clientsByChainID {
		if err := validateContractAddress(ctx, client.client, client.collectionFactoryAddress); err != nil {
			return nil, fmt.Errorf("failed to validate contract address: %w", err)
		}
	}

	return &ChainEventsListener{clientsByChainID, marketplaceService}, nil
}

// Start register to listen blockchain events
func (l *ChainEventsListener) Start(ctx context.Context) error {

	var eg errgroup.Group
	eg.SetLimit(len(l.clientsByChainID))

	for chainID, client := range l.clientsByChainID {
		c := client

		eg.Go(func() error {
			slog.InfoContext(ctx, "start monitoring", slog.Int64("chain_id", chainID), slog.String("collection_factory_address", c.collectionFactoryAddress))

			collectionFactoryContract, err := contracts.NewCollectionFactory(common.HexToAddress(c.collectionFactoryAddress), c.client)
			if err != nil {
				return err
			}

			return l.watchCollectionDeployedLog(ctx, collectionFactoryContract)
		})
	}

	return eg.Wait()
}

func (l *ChainEventsListener) watchCollectionDeployedLog(ctx context.Context, collectionFactoryContract *contracts.CollectionFactory) error {
	events := make(chan *contracts.CollectionFactoryCollectionDeployedLog)

	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}

	subscription, err := collectionFactoryContract.WatchCollectionDeployedLog(opts, events, nil, nil)
	if err != nil {
		return err
	}

	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			return errChan
		case event := <-events:
			slog.InfoContext(ctx, "collection deployed log event received", slog.String("collection_name", event.CollectionName), slog.String("collection_address_hex", event.CollectionAddress.Hex()), slog.String("owner_address_hex", event.Owner.Hex()))

			if err := l.marketplaceService.ProcessCollectionDeployed(ctx, event.CollectionName, event.CollectionAddress, event.Owner); err != nil {
				slog.ErrorContext(ctx, "failed to process collection deployed", slog.String("error", err.Error()), slog.String("collection_name", event.CollectionName), slog.String("collection_address_hex", event.CollectionAddress.Hex()), slog.String("owner_address_hex", event.Owner.Hex()))
				return fmt.Errorf("failed to process collection deployed: %w", err)
			}
		}
	}
}

// validateContractAddress validate the contract address checking if the contract is deployed
func validateContractAddress(ctx context.Context, client *ethclient.Client, address string) error {
	if err := validateAddress(address); err != nil {
		return fmt.Errorf("failed to validate address: %w", err)
	}
	contractAddress := common.HexToAddress(address)
	bytecode, err := client.CodeAt(ctx, contractAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to get code at contract address %s: %w", address, err)
	}

	if len(bytecode) == 0 {
		return ErrInvalidContractAddress
	}

	return nil
}

// validateAddress validate address format
func validateAddress(address string) error {
	regex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if ok := regex.MatchString(address); !ok {
		return ErrInvalidAddress
	}

	return nil
}

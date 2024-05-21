package config

import (
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"syscall"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

type (
	Config struct {
		Profile     Profile
		Logger      Logger
		Postgres    Postgres
		Marketplace Marketplace
	}

	Postgres struct {
		Address      string
		Database     string
		Username     string
		Password     string
		Insecure     bool
		ServiceName  string
		QueryHook    bool
		ReadTimeout  time.Duration
		DialTimeout  time.Duration
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  time.Duration
		MaxLifeTime  time.Duration
	}

	Profile struct {
		Port string
		Flag string
	}

	Logger struct {
		Debug     bool
		Pretty    bool
		AddSource bool
	}

	Marketplace struct {
		Clients map[int64]Client
	}

	Client struct {
		Enabled                  bool
		HttpURL                  string
		WssURL                   string
		CollectionFactoryAddress string
	}
)

func Load(path string, shutdownSignal chan os.Signal) (*Config, error) {
	k := koanf.New(".")

	var config Config

	provider := file.Provider(path)

	if err := k.Load(provider, yaml.Parser()); err != nil {
		return nil, fmt.Errorf("load file error: %w", err)
	}

	if err := k.Unmarshal("", &config); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	if err := provider.Watch(func(event any, err error) {
		slog.Info("config changed: restarting the pod")
		shutdownSignal <- syscall.SIGTERM
	}); err != nil {
		return nil, fmt.Errorf("file watch error: %w", err)
	}

	return &config, nil
}

func (pg *Postgres) DSN() string {
	query := url.Values{"sslmode": {"disable"}}
	return fmt.Sprintf("postgres://%s@%s/%s?%s", url.UserPassword(pg.Username, pg.Password).String(), pg.Address, pg.Database, query.Encode())
}

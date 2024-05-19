package blockchain

import (
	"context"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/envhelper"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/contracts"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log/slog"
	"math/big"
	"os"
)

type ETHClient struct {
	client *ethclient.Client
}

func NewETHClient(client *ethclient.Client) *ETHClient {
	return &ETHClient{client: client}
}

func (c *ETHClient) CreateCollection(ctx context.Context, collection *model.Collection, chainID int64, marketplaceAccountHex string) (string, error) {
	keystoreLocation := envhelper.GetEnv("KEYSTORE")
	collectionID := big.NewInt(collection.ID)

	// FIXME: Replace local keystore with actual user key/signature
	keystore, err := os.Open(keystoreLocation)
	if err != nil {
		slog.ErrorContext(ctx, "failed to load keystore", slog.String("keystore_location", keystoreLocation), slog.String("error", err.Error()))
		return "", err
	}

	defer keystore.Close()

	auth, err := bind.NewTransactorWithChainID(keystore, envhelper.GetEnv("KEYSTOREPASS"), big.NewInt(chainID))
	if err != nil {
		slog.ErrorContext(ctx, "failed to load keystore", slog.String("keystore_location", keystoreLocation), slog.String("error", err.Error()))
		return "", err
	}

	// get ids and amounts
	ids := make([]*big.Int, len(collection.Items))
	totalAmounts := make([]*big.Int, len(collection.Items))
	for i, item := range collection.Items {
		ids[i] = big.NewInt(item.ID)
		totalAmounts[i] = big.NewInt(item.TotalAmount)
	}

	marketplaceAddress := common.HexToAddress(marketplaceAccountHex)

	// deploy collection
	collectionAddress, deployCollectionTx, _, err := contracts.DeployCollection(auth, c.client, collectionID, collection.Name, collection.GetImageURL(), ids, totalAmounts, collection.BaseImagePath, marketplaceAddress)
	if err != nil {
		slog.ErrorContext(ctx, "failed to deploy Collection contract", slog.String("error", err.Error()))
		return "", err
	}

	slog.Info("collection contract deployed. Waiting for tx %s to be confirmed", deployCollectionTx.Hash().Hex())

	// add collection to marketplace
	marketplaceTransactor, err := contracts.NewMarketplaceTransactor(marketplaceAddress, c.client)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create Marketplace Transactor", slog.String("error", err.Error()))
		return "", err
	}

	marketplaceTx, err := marketplaceTransactor.AddCollection(auth, collectionAddress, collectionID)
	if err != nil {
		slog.ErrorContext(ctx, "failed to execute transaction add Collection to marketplace", slog.String("error", err.Error()))
		return "", err
	}

	slog.Info("collection added to marketplace. Waiting for tx %s to be confirmed", marketplaceTx.Hash().Hex())

	// TODO: Do we need to persist transactions?

	// TODO: Do we need to call marketplace addItems function?

	return collectionAddress.Hex(), nil
}

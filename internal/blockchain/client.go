package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func Connect(url string) (*ethclient.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, err
	}
	return client, nil
}

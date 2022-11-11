package ethereumsupport

import (
	"errors"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrorNullRpc = errors.New("the rpc is null")
)

func Provider() (*Provide, error) {
	rpc := os.Getenv("rpc")
	if rpc == "" {
		return nil, ErrorNullRpc
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	return &Provide{Client: client}, nil
}

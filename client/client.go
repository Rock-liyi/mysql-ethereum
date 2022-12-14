package client

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
}

var clientInstance *ClientSingleton

type ClientSingleton struct {
	once sync.Once
}

func (C *ClientSingleton) GetSingleton() *ClientSingleton {
	C.once.Do(func() {
		clientInstance = &ClientSingleton{}
	})
	return clientInstance
}

func (C *ClientSingleton) NewClient() *ethclient.Client {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	//client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	account := common.HexToAddress("0x690B9A9E9aa1C9dB991C7721a92d351Db4FaC990")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
	// fmt.Println
	// fmt
	//	}{}

}

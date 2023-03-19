package main

import (
	"call-contract-in-go/cdd"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

var (
	PrivateKey, _ = crypto.HexToECDSA("your private key")
	FromAddr      = common.HexToAddress("0x89f4931f1E14b3908DBC33C9927d8EfbA3b2fEdf")
	ToAddr        = common.HexToAddress("0xF1C4a341624e4861725B91F78A41ac08CD2Ddd6E")
	ContractCdd   = "0xCbDCE717e290aAdb09fbd009DD14dBCAC30Bf8B1"
	BscTestNet    = "https://data-seed-prebsc-2-s2.binance.org:8545/"
)

func main() {
	//var txHash string
	client, err := ethclient.Dial(BscTestNet)
	if err != nil {
		fmt.Println("eth client: ", err)
		os.Exit(1)
	}

	//txHash, err = cdd.CallContract(client, PrivateKey, FromAddr, ToAddr, ContractCdd)
	//if err != nil {
	//	os.Exit(1)
	//}
	//fmt.Println("tx hash: ", txHash)
	//
	//txHash, err = cdd.CallContractWithAbi(client, PrivateKey, FromAddr, ToAddr, ContractCdd)
	//if err != nil {
	//	os.Exit(1)
	//}
	//fmt.Println("tx hash: ", txHash)

	cdd2, err := cdd.NewCdd2Token(common.HexToAddress(ContractCdd), client)
	if err != nil {
		fmt.Println("new token: ", err)
		os.Exit(1)
	}
	total, err := cdd2.TotalSupply(nil)
	if err != nil {
		fmt.Println("total supply: ", err)
		os.Exit(1)
	}
	fmt.Println("total supply: ", total.String())

	nonce, err := client.NonceAt(context.Background(), FromAddr, nil)
	if err != nil {
		fmt.Println("get nonce: ", err)
		os.Exit(1)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("get gas price: ", err)
		os.Exit(1)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(PrivateKey, big.NewInt(97))
	if err != nil {
		fmt.Println("new transactor: ", err)
		os.Exit(1)
	}
	amount, _ := new(big.Int).SetString("5000000", 10)
	opts.GasLimit = uint64(300000)
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.GasPrice = gasPrice
	tx, err := cdd2.Transfer(opts, ToAddr, amount)
	if err != nil {
		fmt.Println("transfer: ", err)
		os.Exit(1)
	}
	fmt.Println("tx hash: ", tx.Hash().Hex())
}

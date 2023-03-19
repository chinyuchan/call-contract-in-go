package call1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func CallContract(client *ethclient.Client, privKey *ecdsa.PrivateKey, from, to common.Address, contract string) (string, error) {
	// create tx
	// nonce
	nonce, err := client.NonceAt(context.Background(), from, nil)
	if err != nil {
		fmt.Println("get nonce: ", err)
		return "", err
	}

	// gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("gas price: ", err)
		return "", err
	}

	// transfer(address,uint256)： 0xa9059cbb
	data := []byte{}
	method := []byte("transfer(address,uint256)")
	method = crypto.Keccak256(method)[:4]
	fmt.Println("method: ", hexutil.Encode(method))
	paddedAddress := common.LeftPadBytes(to.Bytes(), 32)
	amount, _ := new(big.Int).SetString("10000000000000000000", 10) //10
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	data = append(data, method...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	// new tx
	tx := types.NewTransaction(nonce, common.HexToAddress(contract), big.NewInt(0), uint64(300000), gasPrice, data)

	// sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(97)), privKey)
	if err != nil {
		fmt.Println("sign tx: ", err)
		return "", err
	}

	// send tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("send tx: ", err)
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	nonce := uint64(1)
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice := big.NewInt(5000)

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID := big.NewInt(8080)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0]) //<-- ts.GetRlp undefined
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Printf(rawTxHex)
}

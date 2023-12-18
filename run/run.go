package run

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"inscriptiontool/wallet"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Info struct {
	Account        string
	Rpc            string
	To             string
	Amount         float64
	GF             float64
	GasLimitFactor int
	Data           string
	Count          int
}

func Run(info Info) error {
	now := time.Now()
	defer func() {
		fmt.Printf("发送 %d 笔交易，耗时: %vs\n", info.Count, time.Since(now).Seconds())
	}()
	account, err := readAccount(info.Account)
	if err != nil {
		return errors.Wrap(err, "read Account error")
	}

	ac := common.HexToAddress(account.Address)

	client, err := ethclient.Dial(info.Rpc)
	if err != nil {
		return errors.Wrap(err, "error dialing ethereum client")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "get chain id")
	}

	nonce, err := client.PendingNonceAt(context.Background(), ac)
	if err != nil {
		return errors.Wrap(err, "get nonce")
	}
	gp, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.Wrap(err, "get gas price")
	}
	f, _ := gp.Float64()
	gasPrice := big.NewInt(int64(f * info.GF))

	to := common.HexToAddress(account.Address)
	if info.To != "" {
		to = common.HexToAddress(info.To)
	}
	amount := big.NewInt(int64(info.Amount * 1e18))

	data := []byte(info.Data)
	signer := types.NewLondonSigner(chainId)
	privateKey, err := crypto.HexToECDSA(account.Private)
	if err != nil {
		return errors.Wrap(err, "error parsing private key")
	}
	return loop(info.Count, info, client, signer, privateKey, nonce, to, amount, gasPrice, data)
}

func loop(count int, info Info,
	client *ethclient.Client,
	signer types.Signer,
	privateKey *ecdsa.PrivateKey,
	nonce uint64,
	to common.Address,
	amount *big.Int,
	gasPrice *big.Int,
	data []byte) error {
	var noncechan = make(chan uint64, 100000)
	go func() {
		for i := 0; i < count; i++ {
			noncechan <- nonce
			nonce++
		}
		close(noncechan)
	}()
	var wg sync.WaitGroup
	for nc := range noncechan {
		wg.Add(1)
		go func(ncc uint64) {
			defer wg.Done()
		PP:
			if ncc%10 == 0 {
				gp, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					fmt.Printf("get gas price error, %v\n", err)
				} else {
					f, _ := gp.Float64()
					gasPrice = big.NewInt(int64(f * info.GF))
				}
			}
			unsignedTx := types.NewTransaction(ncc, to, amount, uint64(21000*info.GasLimitFactor), gasPrice, data)
			signedTx, err := types.SignTx(unsignedTx, signer, privateKey)
			if err != nil {
				fmt.Printf("sign transaction error, %v\n", err)
				return
			}
			fmt.Println("=====================================================================================================")
			fmt.Printf("tx hash: %s\n", signedTx.Hash().Hex())
			fmt.Printf("nonce : %d\n", ncc)
			fmt.Printf("gas price : %dgwei\n", gasPrice.Int64()/1e9)
			fmt.Printf("gas limit : %d\n", signedTx.Gas())
			fmt.Printf("from : %s\n", info.Account)
			fmt.Printf("to : %s\n", signedTx.To().Hex())
			fmt.Printf("amount : %d\n", signedTx.Value())
			fmt.Printf("data : %s\n", signedTx.Data())
			fmt.Println("=====================================================================================================")
			// 3. send rawTx
			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				fmt.Printf("send transaction error, %v\n", err)
				goto PP
			}
		}(nc)
	}
	wg.Wait()
	return nil
}

func readAccount(account string) (*wallet.Wallet, error) {
	file, err := os.ReadFile(filepath.Join(wallet.RootDir, fmt.Sprintf("%s.json", account)))
	if err != nil {
		return nil, errors.Wrap(err, "read file error")
	}
	var w wallet.Wallet
	err = json.Unmarshal(file, &w)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal wallet error")
	}
	fmt.Printf("Use Account: %s, address : %s\n", w.Name, w.Address)
	return &w, nil
}

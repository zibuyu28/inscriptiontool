package wallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/tyler-smith/go-bip39"
	"os"
	"path/filepath"
)

type Wallet struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Private  string `json:"private"`
	Mnemonic string `json:"mnemonic"`
}

var RootDir = ""

func init() {
	dir, _ := os.UserHomeDir()
	RootDir = filepath.Join(dir, ".inscription_tool", "wallets")
}

func Create(name string) error {

	if name == "" {
		return errors.New("name is empty")
	}
	res, err := checkWalletExist(name)
	if err != nil {
		return errors.Wrap(err, "check wallet exist error")
	}
	if res {
		return errors.Errorf("wallet [%s] already exist", name)
	}
	// 生成助记词
	entropy, err := bip39.NewEntropy(128) // 126位的随机数生成助记词
	if err != nil {
		return errors.Wrap(err, "create entropy error")
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return errors.Wrap(err, "create mnemonic error")
	}

	// 从助记词派生出私钥
	seed := bip39.NewSeed(mnemonic, "")          // 传入空字符串作为密码
	privateKey, err := crypto.ToECDSA(seed[:32]) // 使用前32字节作为私钥
	if err != nil {
		return errors.Wrap(err, "create private key error")
	}

	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	wallet := Wallet{
		Name:     name,
		Address:  address.Hex(),
		Private:  hexutil.Encode(crypto.FromECDSA(privateKey))[2:],
		Mnemonic: mnemonic,
	}
	err = writeWallet(wallet)
	if err != nil {
		return errors.Wrap(err, "write wallet error")
	}
	return nil
}

func checkWalletExist(name string) (bool, error) {
	_ = os.MkdirAll(RootDir, os.ModePerm)
	walletName := fmt.Sprintf("%s.json", name)
	_, err := os.Stat(filepath.Join(RootDir, walletName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, errors.Wrap(err, "stat file error")
	}
	return true, nil
}

func writeWallet(wallet Wallet) error {
	_ = os.MkdirAll(RootDir, os.ModePerm)
	walletName := fmt.Sprintf("%s.json", wallet.Name)
	file, err := os.OpenFile(filepath.Join(RootDir, walletName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, "open file error")
	}
	defer file.Close()

	marshal, err := json.MarshalIndent(wallet, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshal wallet error")
	}
	_, err = file.Write(marshal)
	if err != nil {
		return errors.Wrap(err, "write file error")
	}
	fmt.Println(string(marshal))
	return nil
}

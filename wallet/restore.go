package wallet

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

func Restore(name, pk string) error {
	if name == "" || pk == "" {
		return errors.Errorf("name [%s] or pk [%s] is empty", name, pk)
	}
	exist, err := checkWalletExist(name)
	if err != nil {
		return errors.Wrap(err, "check wallet exist error")
	}
	if exist {
		return errors.Errorf("wallet [%s] already exist", name)
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return errors.Wrap(err, "error parsing private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}
	toAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	var wallet Wallet
	wallet.Name = name
	wallet.Address = toAddress.Hex()
	wallet.Private = pk
	err = writeWallet(wallet)
	if err != nil {
		return errors.Wrap(err, "write wallet error")
	}
	return nil
}

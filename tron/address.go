package tron

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
	"github.com/tyler-smith/go-bip39"
	"github.com/weibi/tron-protocol/common/base58"
)

func CreateMnemonic() (mnemonic string, err error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	return
}

func IsMnemonicValid(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

func AddressListFromMnemonic(mnemonic string) (addresses []string, err error) {
	for i := 0; i < 100; i++ {
		address, errA := AddressFromMnemonic(mnemonic, i)
		if errA != nil {
			return nil, errA
		}

		addresses = append(addresses, address)
	}
	return
}

func AddressFromMnemonic(mnemonic string, idx int) (address string, err error) {
	private, err := PrivateFromMnemonic(mnemonic, idx)
	if err != nil {
		return
	}

	address = AddressFromEsdaPrivateKey(private)

	return
}

// https://likefacai.com/archives/golangbo-chang-qian-bao-de-jian-dan-demo
func PrivateFromMnemonic(mnemonic string, idx int) (privateKey *ecdsa.PrivateKey, err error) {
	if !bip39.IsMnemonicValid(mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic/助记词错误")
	}
	private, _ := keys.FromMnemonicSeedAndPassphrase(mnemonic, "", idx)
	privateKey = private.ToECDSA()
	return
}

const AddressLength = 21

type Address [AddressLength]byte

func (a Address) Bytes() []byte {
	return a[:]
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func AddressFromEsdaPrivateKey(key *ecdsa.PrivateKey) (addr string) {
	//address := PubkeyToAddress(key.PublicKey)
	addrBytes := PubkeyToAddress(key.PublicKey).Bytes()
	addr = base58.EncodeCheck(addrBytes)
	return
}

func PubkeyToAddress(p ecdsa.PublicKey) Address {
	address := crypto.PubkeyToAddress(p)
	addressTron := append([]byte{0x41}, address.Bytes()...)
	return BytesToAddress(addressTron)
}

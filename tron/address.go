package tron

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
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

func IsAddressValid(addr string) bool {
	_, err := common.DecodeCheck(addr)
	return err == nil
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

type Account struct {
	OwnerPermission struct {
		Keys []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		PermissionName string `json:"permission_name"`
	} `json:"owner_permission"`
	FreeAssetNetUsageV2 []struct {
		Value int    `json:"value"`
		Key   string `json:"key"`
	} `json:"free_asset_net_usageV2"`
	AccountResource struct {
	} `json:"account_resource"`
	ActivePermission []struct {
		Operations string `json:"operations"`
		Keys       []struct {
			Address string `json:"address"`
			Weight  int    `json:"weight"`
		} `json:"keys"`
		Threshold      int    `json:"threshold"`
		ID             int    `json:"id"`
		Type           string `json:"type"`
		PermissionName string `json:"permission_name"`
	} `json:"active_permission"`
	AssetV2 []struct {
		Value int    `json:"value"`
		Key   string `json:"key"`
	} `json:"assetV2"`
	Address    string              `json:"address"`
	Balance    uint64              `json:"balance"`
	CreateTime int64               `json:"create_time"`
	Trc20      []map[string]string `json:"trc20"`
}

// {"freeNetLimit": 5000,"TotalNetLimit": 43200000000,"TotalNetWeight": 30110876345,"TotalEnergyLimit": 90000000000,"TotalEnergyWeight": 13036858}
type AccountResource struct {
	NetUsed           uint64 `json:"NetUsed"`
	NetLimit          uint64 `json:"NetLimit"`
	EnergyLimit       uint64 `json:"EnergyLimit"`
	EnergyUsed        uint64 `json:"EnergyUsed"`
	FreeNetLimit      uint64 `json:"freeNetLimit"`
	TotalNetLimit     uint64 `json:"TotalNetLimit"`
	TotalNetWeight    uint64 `json:"TotalNetWeight"`
	TotalEnergyLimit  uint64 `json:"TotalEnergyLimit"`
	TotalEnergyWeight uint64 `json:"TotalEnergyWeight"`
	FreeNetUsed       uint64 `json:"freeNetUsed"`
}

func GetWalletAccount(addr string) (account Account, err error) {
	url := fmt.Sprintf("%s/wallet/getaccount", fullApi)
	_, err = restyClient.R().
		ForceContentType("application/json").
		SetResult(&account).
		SetBody(map[string]interface{}{
			"address": addr,
			"visible": true,
		}).Post(url)

	if err != nil {
		fmt.Printf("GetWalletAccount1  failed:%v\n", err)
		return
	}
	//  {"address": "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX","balance": 2000000,"create_time": 1586426514000,"account_resource": {},"owner_permission": {"permission_name": "owner","threshold": 1,"keys": [{"address": "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX","weight": 1}]},"active_permission": [{"type": "Active","id": 2,"permission_name": "active","threshold": 1,"operations": "7fff1fc0033e0b00000000000000000000000000000000000000000000000000","keys": [{"address": "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX","weight": 1}]}],"assetV2": [{"key": "1000026","value": 28817511379}],"free_asset_net_usageV2": [{"key": "1000026","value": 0}]}
	//logger.InfoSimple("resp: %s", resp.Body())
	return
}

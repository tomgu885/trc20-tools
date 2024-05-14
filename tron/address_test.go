package tron

import (
	"fmt"
	"testing"
)

func TestIsAddressValid(t *testing.T) {
	addrList := []string{
		"fefez",
		"TBh9ReKAXSiK7SmkPziMRx1PKYqWKUHyP6", // ok
		"TTZ3vBn7oZhck8CQC1XfrzbDZWor5pDoPv", // ok
		"TXfG1sQiJM3nhhkUEgNFBCn9PFbfLGgcPn", // not
		"TXfG1sQiJM3nhhkUEgNFBCn9PFbfLGgcPa", // not
		"TMQr8fFigT7Z25bkdGeGKSsjdkPkrHJgnL", // ok
		"TMQr8fFigT7Z25bkdGeGKSsjdkPkrHJgnn", // not ok
	}

	for _, addr := range addrList {
		valid := IsAddressValid(addr)
		fmt.Printf("%s is valid: %t\n", addr, valid)
	}
}

func TestAddressFromMnemonic(t *testing.T) {
	mnemonic := "flush impose belt soft volume grit begin afford educate tray pigeon pluck"

	for i := 0; i < 10; i++ {
		address, err := AddressFromMnemonic(mnemonic, i)
		if err != nil {
			fmt.Println("err:", err.Error())
			continue
		}

		fmt.Printf("%d: %s\n", i, address)
	}
}

func TestGetWalletAccount(t *testing.T) {
	addr := "TBh9ReKAXSiK7SmkPziMRx1PKYqWKUHyP6"
	info, err := GetWalletAccount(addr)
	if err != nil {
		return
	}

	fmt.Printf("trx balance: %d\n", info.Balance)
}

package cmd

import (
	"fmt"
	"strings"
	"trc20-readline/tron"
)

func CreateMnemonic() string {

	return ""
}

func GenerateAddress(mnemonic string) (addresses []string, err error) {
	words := strings.Split(mnemonic, " ")
	if len(words) != 12 {
		fmt.Println("len:", len(words))
		return nil, fmt.Errorf("invalid mnemonic,助记词必须是12个")
	}
	addresses, err = tron.AddressListFromMnemonic(mnemonic)
	return
	//return []string{"addrr1", "addr2"}, nil
}

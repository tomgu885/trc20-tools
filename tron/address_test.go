package tron

import (
	"fmt"
	"testing"
)

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

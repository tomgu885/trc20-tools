package cmd

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	mnemonic := "flush impose belt soft volume grit begin afford educate tray pigeon pluck"

	balances, err := Query(mnemonic)
	if err != nil {
		fmt.Println("failed:", err.Error())
		return
	}

	for _, balance := range balances {
		fmt.Println("balance:", balance)
	}
}

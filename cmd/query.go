package cmd

import (
	"fmt"
	"trc20-readline/tron"
)

func Query(mnemonic string) (balances []string, err error) {

	addresses, err := GenerateAddress(mnemonic)
	if err != nil {
		return
	}

	for idx, addr := range addresses {
		trxBalance, errT := tron.BalanceOf(addr)
		if errT != nil {
			return []string{}, errT
		}

		fmt.Printf("trx %d: %s\n", idx, trxBalance)
	}

	return
}

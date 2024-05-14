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

		//fmt.Printf("trx %d %s: %s\n", idx, addr, tron.Balance2Trx(trxBalance))

		trc20Balance, errT := tron.BalanceOfTrc20(addr)
		if errT != nil {
			continue
		}

		if trc20Balance.Int64() != 0 || trxBalance > 0 {
			line := fmt.Sprintf("%2d  %s\t %s\t %s", idx, addr, tron.Balance2Trx(trxBalance), tron.BigToken2Usdt(trc20Balance))
			balances = append(balances, line)
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf("\n")
	return
}

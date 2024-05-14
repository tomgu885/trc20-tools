package cmd

import (
	"errors"
	"fmt"
	"time"
	"trc20-readline/tron"
)

func Send(mnemonic, to, amount string, idx int) (txID string, err error) {
	amountInt, err := tron.String2int64(amount)
	if err != nil {
		err = errors.New("数字不正确")
	}

	from, err := tron.AddressFromMnemonic(mnemonic, idx)
	if err != nil {
		return
	}

	trc20Balance, err := tron.BalanceOfTrc20(from)

	if err != nil {
		return
	}

	if amountInt > trc20Balance.Int64() {
		err = fmt.Errorf("%s 余额不足", from)
		return
	}

	trxBalance, err := tron.BalanceOf(from)
	if err != nil {
		return
	}

	fmt.Printf("from: %s\n", from)

	if trxBalance < 20_000_000 { // 小于 20 TRX
		// 查询主账号的 TRX 余额
		if 0 == idx {
			err = errors.New("主账号也没有 TRX")
			return
		}
		fmt.Printf("%s TRX 不足，从主账号转 20 TRX过去...\n")
		blockNum, trxID, errTrx := tron.SendTrx(mnemonic, 0, from, "20")
		if errTrx != nil {
			err = errors.New("TRX加油失败")
			return
		}
		fmt.Printf("trx 加油: txID:%s @ block: %d %s\n", trxID, blockNum)
		fmt.Println("等待5秒钟")
		time.Sleep(5 * time.Second)
	}

	txID, err = tron.SendTrc20(mnemonic, idx, to, amount)

	return
}

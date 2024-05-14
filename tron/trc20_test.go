package tron

import (
	"fmt"
	"testing"
)

func TestBalanceOf(t *testing.T) {
	addr := "TBh9ReKAXSiK7SmkPziMRx1PKYqWKUHyP6"
	balance, err := BalanceOf(addr)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Printf("balance: %d\n", balance)
}

func TestBalanceOfTrc20(t *testing.T) {
	//Init(false)
	addr := "TBGgZLwcEH3qCy1ApNG4CfLfpyCX7EQozd"
	//tokenAddr := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	balance, err := BalanceOfTrc20(addr)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Printf("Balance:%s\n", balance.String())
}

func TestSendTrx(t *testing.T) {
	mnemonic := "flush impose belt soft volume grit begin afford educate tray pigeon pluck"
	toAddr := "TQSoNPL2Q4nM77aMbVg5KwNpUtAP6qiLSS"
	SendTrx(mnemonic, 0, toAddr, "10.4")
}

func TestSendTrc20(t *testing.T) {
	mnemonic := "flush impose belt soft volume grit begin afford educate tray pigeon pluck"
	toAddrList := []string{
		"TMQr8fFigT7Z25bkdGeGKSsjdkPkrHJgnL",
		"TTZ3vBn7oZhck8CQC1XfrzbDZWor5pDoPv",
		"TWk2aPcVS5BJFCzezJim55Rzw6Deh1y3TJ",
		"TGFSWEn4Sj3AbQMbNGPdLBuixvma7pQk6y",
	}
	for _, toAddr := range toAddrList {

		tx, err := SendTrc20(mnemonic, 1, toAddr, "19.301")
		if err != nil {
			fmt.Println("SendTrc20 failed: ", err.Error())
			return
		}

		fmt.Printf("to %s tx:%s \n", toAddr, tx)
	}
}

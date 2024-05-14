package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"slices"
	"strconv"
	"strings"
	"trc20-readline/cmd"
	"trc20-readline/tron"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:      "\033[31m»\033[0m ",
		HistoryFile: "/tmp/trc20_readline.tmp",
	})

	if err != nil {
		fmt.Println("错误:")
		panic(err)
	}

	defer l.Close()

	l.CaptureExitSignal()

	mnemmonic := ""
	modes := []string{"gen", "query", "send", "help"}
	env := "mainnet"
	mode := ""
	fromIndex := int(-1)
	toAddr := ""
	amount := ""
	//confirm := false
	if env == "" {
		cmd.EnvPrompt(l) //  设置环境
	} else {
		cmd.MnemonicPrompt(l)
		tron.Init(env == "shasta")
	}
	// 循环
	for {
		line, errL := l.Readline()
		if errL == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		}

		if len(line) == 0 {
			continue // 输入空行, 则要求用户再输入
		}

		line = strings.TrimSpace(line)
		//fmt.Println("...line:", line)
		switch {
		case "quit" == line || "bye" == line: // 再见
			fmt.Println("bye 再见")
			os.Exit(0)
		case "" == env:
			if slices.Contains([]string{"prod", "shasta"}, line) {
				env = line
				tron.Init(env == "shasta")
				cmd.MnemonicPrompt(l)
			} else {
				fmt.Println("请输入 prod 或者shasta")
			}

		case "new" == line:
			mnemmonic, err = tron.CreateMnemonic()
			if err != nil {
				fmt.Println("发生错误:", err.Error())
				continue
			}
			fmt.Println("助记词: ", mnemmonic)
			cmd.ModePrompt(l)
		case "env" == line:
			env = ""
			fmt.Println("请重新输入环境")
			cmd.EnvPrompt(l)
			continue
		case "mn" == line:
			mnemmonic = ""
			fmt.Println("助记词消除，输入新的助记词")
			cmd.MnemonicPrompt(l)
			continue
		case "gen" == line: // start of gen
			mode = ""
			if "" == mnemmonic {
				fmt.Println("未输入助记词")
				cmd.MnemonicPrompt(l)
				continue
			}
			fmt.Println("gen....")
			addresses, errGen := cmd.GenerateAddress(mnemmonic)
			if errGen != nil {
				fmt.Println("发生错误:", errGen.Error())
			} else {
				fmt.Println("生成的地址如下")
				for idx, addr := range addresses {
					fmt.Printf("%3d: %s\n", idx, addr)
				}

			}
			cmd.ModePrompt(l)
			continue
			// end of gen:
		case "query" == line:
			mode = ""
			if "" == mnemmonic {
				fmt.Println("未输入助记词")
				cmd.ModePrompt(l)
				continue
			}
			fmt.Println("querying")
			balances, errB := cmd.Query(mnemmonic)
			if errB != nil {
				fmt.Println("发生错误:", errB.Error())
				continue
			}

			fmt.Printf("id  地址\t\t\t\tTRX\t\t USDT \t| %d个地址有TRX或USDT\n", len(balances))
			for _, balance := range balances {
				fmt.Println(balance)
			}
			cmd.ModePrompt(l)
			continue
			// end of query
		case slices.Contains(modes, line): // set mode
			mode = line
			//fmt.Println("mode:", mode)
			cmd.SetPrompt(mode, l)
		case "" == mnemmonic:
			if !tron.IsMnemonicValid(line) {
				fmt.Println("助记词不正确")
				cmd.SetPrompt(mode, l)
				continue
			}
			mnemmonic = line
			cmd.ModePrompt(l)
		case "" == mode: // no mode setted
			fmt.Printf("\n")
			cmd.MnemonicPrompt(l)
			fromIndex = -1
			toAddr = ""
			amount = ""
			continue
		case mode == "send":
			if "clear" == line {
				fromIndex = -1
				toAddr = ""
				amount = ""

				fmt.Println("请输入接收地址")
				l.SetPrompt("address\u001B[31m»\u001B[0m ")
				continue
			}
			if toAddr == "" {
				tron.IsAddressValid(line)
				toAddr = line
				l.SetPrompt("发送地址 id\u001B[31m»\u001B[0m ")
				continue
			}

			if fromIndex < 0 {
				idx, errI := strconv.Atoi(line)
				if errI != nil {
					fmt.Println("输入的id 不是数字")
					l.SetPrompt("发送地址 id\u001B[31m»\u001B[0m ")
					continue
				}
				fromIndex = idx
				l.SetPrompt("转账金额:")
				continue
			}

			if tron.IsNumeric(line) {
				amount = line
				from, _ := tron.AddressFromMnemonic(mnemmonic, fromIndex)
				//trc20balance, errB := tron.BalanceOfTrc20(from)
				enough, thisBalance, errB := tron.HasEnoughTrc20(from, amount)
				if errB != nil {
					fmt.Println("发生错误:", errB.Error())
				}

				if !enough {
					fmt.Printf("%s (%s) 金额不足以转出(%s)，请重新输入\n", from, tron.BigToken2Usdt(thisBalance), amount)
					mode = ""
					fromIndex = -1
					toAddr = ""
					amount = ""
					cmd.ModePrompt(l)
					continue
				}

				amount = line
				fmt.Printf("从 %s(余额 %s USDT) 转到 %s 金额 %s\n", from, tron.BigToken2Usdt(thisBalance), toAddr, amount)
				fmt.Println("输入 yes 或者 y 确认| 输入 clear 重新输入")
				l.SetPrompt("\u001B[31m»\u001B[0m ")

				continue
			}

			if slices.Contains([]string{"y", "yes"}, line) && fromIndex >= 0 && toAddr != "" && amount != "" {
				txHash, errSend := cmd.Send(mnemmonic, toAddr, amount, fromIndex)
				if errSend != nil {
					fmt.Println("发生错误:", errSend.Error())
					mode = ""
					fromIndex = -1
					toAddr = ""
					amount = ""
					cmd.ModePrompt(l)
					continue
				} else {
					fmt.Println("USDT 发送成功: ", txHash)
				}

				mode = ""
				fromIndex = -1
				toAddr = ""
				amount = ""
				cmd.ModePrompt(l)
			}
			// end of send
		}

	}

	fmt.Println("end")
}

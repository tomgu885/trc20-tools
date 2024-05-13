package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"slices"
	"strings"
	"trc20-readline/cmd"
	"trc20-readline/tron"
)

func main() {
	fmt.Println("start")
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
	modes := []string{"gen", "query", "transfer", "help"}
	env := ""
	mode := ""

	cmd.EnvPrompt(l) //  设置环境
	for {
		line, errL := l.Readline()
		if errL == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		}

		line = strings.TrimSpace(line)
		fmt.Println("...line:", line)
		switch {
		case "quit" == line || "bye" == line: // 再见
			fmt.Println("bye 再见")
			os.Exit(0)
		case "" == env:
			if slices.Contains([]string{"prod", "shasta"}, line) {
				env = line
				tron.Init(env == "shasta")
			}
			cmd.MnemonicPrompt(l)
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
			continue
		case "gen" == line:
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
					fmt.Printf("%03d: %s\n", idx, addr)
				}
				mode = ""
			}
			continue
		case "query" == line:
			if "" == mnemmonic {
				fmt.Println("未输入助记词")
				cmd.ModePrompt(l)
				continue
			}
			fmt.Println("querying")
		case slices.Contains(modes, line):
			mode = line
			l.SetPrompt(fmt.Sprintf("act: %s", mode))
			cmd.SetPrompt(mode, l)
		case "" == mnemmonic:
			if !tron.IsMnemonicValid(line) {
				fmt.Println("助记词不正确1")
				cmd.SetPrompt(mode, l)
				continue
			}
			mnemmonic = line
			cmd.ModePrompt(l)
		case "" == mode: // no mode setted
			fmt.Printf("\n")
			cmd.MnemonicPrompt(l)
			//cmd.Help(l)

		case mode == "transfer":

		}

	}

	fmt.Println("end")
}

package cmd

import (
	"fmt"
	"github.com/chzyer/readline"
)

func EnvPrompt(l *readline.Instance) {
	fmt.Println("请输入网络, prod 正式 shasta 测试")
	l.SetPrompt("env\u001B[31m»\u001B[0m ")
}

func MnemonicPrompt(l *readline.Instance) {
	//fmt.Println("输入你要进行的操作 gen 生成地址, query 查询余额, transfer 转账:\n")
	//l.SetPrompt("mode\u001B[31m»\u001B[0m ")
	fmt.Println("input mnemonic/输入助记词 (不需分行) 或者 输入new 生成新的, quit(ctrl+c) 退出:\n")
	l.SetPrompt("助记词\u001B[31m»\u001B[0m ")
}

func ModePrompt(l *readline.Instance) {
	fmt.Println("输入你要进行的操作 new 新助记词, gen 生成地址, query 查询余额, transfer 转账,  mn 输入新的助记词, env:环境 quit(ctrl+c) 退出:\n")
	l.SetPrompt("mode\u001B[31m»\u001B[0m ")
}

// "gen", "query", "transfer", "help"
func SetPrompt(mode string, l *readline.Instance) {
	switch mode {
	case "gen":
		fmt.Println("input mnemonic/输入助记词 (不需分行):\n")
		l.SetPrompt("助记词\u001B[31m»\u001B[0m ")
		//l.SetPrompt("input mnemonic/输入助记词:\n")
	case "query":
		l.SetPrompt("input mnemonic/输入助记词")
	case "transfer":
		l.SetPrompt("input mnemonic/输入助记词")

	}
}

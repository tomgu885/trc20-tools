package cmd

import "github.com/chzyer/readline"

func Help(l *readline.Instance) {
	l.SetPrompt("帮助文本\n第二行")
}

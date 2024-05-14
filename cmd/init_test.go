package cmd

import (
	"testing"
	"trc20-readline/tron"
)

func TestMain(m *testing.M) {
	tron.Init(true)
	m.Run()
}

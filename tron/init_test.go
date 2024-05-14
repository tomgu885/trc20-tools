package tron

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	Init(true)
	m.Run()
}

func TestBlockId(t *testing.T) {
	fmt.Println("g.Address:", grpcClient.Address)
	//err := grpcClient.Start()
	//if err != nil {
	//	fmt.Println("failed: ", err.Error())
	//	return
	//}
	block, err := grpcClient.GetNowBlock()
	if err != nil {
		fmt.Println("failed: ", err.Error())
		return
	}

	fmt.Printf("blockId:%d\n", block.BlockHeader.GetRawData().Number)
}

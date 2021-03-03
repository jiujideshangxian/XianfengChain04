package main

import (
	"XianfengChain04/chain"
	"fmt"
)

func main() {
	fmt.Println("HelloWorld")

	/*block0 := chain.Block{
		Height: 0,
		Version: 0x00,
		TimeStamp: time.Now().Unix(),
		None: 0,
		PrevHash: [32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	}
	fmt.Println(block0)*/

	block0 :=chain.CreateGenesis([]byte("hi"))
	block1 := chain.NewBlock(block0.Height,[32]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},[]byte("helloword"))
	fmt.Println(block0)
	fmt.Println(block1)
}

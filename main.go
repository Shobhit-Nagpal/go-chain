package main

import (
	"fmt"
	"github.com/Shobhit-Nagpal/go-chain/internal/block"
	"github.com/Shobhit-Nagpal/go-chain/internal/hash"
)

func main() {
  fmt.Println("Hello world")
  genesisBlockHash := hash.CreateNewHash()
  genesisBlock := block.CreateBlock(genesisBlockHash)
  fmt.Println(genesisBlock)
  fmt.Println(genesisBlock.GetPreviousHash())
}

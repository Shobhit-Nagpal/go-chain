package main

import (
  "fmt"
  "github.com/Shobhit-Nagpal/go-chain/internal/block"
)

func main() {
  fmt.Println("Hello world")
  block := block.CreateBlock("lmfaoooo")
  fmt.Println(block)
}

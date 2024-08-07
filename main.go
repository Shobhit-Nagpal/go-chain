package main

import (
	"fmt"

	"github.com/Shobhit-Nagpal/go-chain/internal/account"
	"github.com/Shobhit-Nagpal/go-chain/internal/block"
	"github.com/Shobhit-Nagpal/go-chain/internal/chain"
	"github.com/Shobhit-Nagpal/go-chain/internal/hash"
	"github.com/Shobhit-Nagpal/go-chain/internal/txn"
)

var accounts map[string]*account.Account = make(map[string]*account.Account)

func main() {
  lmaoChain := chain.CreateChain("lmaoChain")
  deez := account.CreateAccount()
  nuts := account.CreateAccount()
  accounts[deez.Id.String()] = deez
  accounts[nuts.Id.String()] = nuts

  transaction, err := txn.CreateTransaction(deez, nuts, 69)
  if err != nil {
    fmt.Println(err.Error())
  }

  genesisHash := hash.CreateNewHash()

  genesisBlock := block.CreateBlock(genesisHash)
  err = genesisBlock.AddTransaction(transaction)
  if err != nil {
    fmt.Println(err.Error())
  }

  //Add consesus mechanism here
  err = lmaoChain.AddBlock(genesisBlock)
  if err != nil {
    fmt.Println(err.Error())
  }

  txn2, err := txn.CreateTransaction(nuts, deez, 39)
  if err != nil {
    fmt.Println(err.Error())
  }

  newBlock := block.CreateBlock([]byte(genesisBlock.GetHash()))
  err = newBlock.AddTransaction(txn2)
  if err != nil {
    fmt.Println(err.Error())
  }

  err = lmaoChain.AddBlock(newBlock)
  if err != nil {
    fmt.Println(err.Error())
  }
}

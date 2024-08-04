package main

import (
	"fmt"

	"github.com/Shobhit-Nagpal/go-chain/internal/account"
	"github.com/Shobhit-Nagpal/go-chain/internal/block"
	"github.com/Shobhit-Nagpal/go-chain/internal/chain"
	"github.com/Shobhit-Nagpal/go-chain/internal/hash"
	"github.com/Shobhit-Nagpal/go-chain/internal/txn"
)

var accounts map[string]account.Account = make(map[string]account.Account)

func main() {
  lmaoChain := chain.CreateChain("lmaoChain")
  deez := account.CreateAccount()
  nuts := account.CreateAccount()
  accounts[deez.Id.String()] = *deez
  accounts[nuts.Id.String()] = *nuts

  transaction, err := txn.CreateTransaction(deez, nuts, 69)
  if err != nil {
    fmt.Println(err.Error())
  }

  genesisHash := hash.CreateNewHash()

  genesisBlock := block.CreateBlock(genesisHash)
  genesisBlock.AddTransaction(transaction)

  //Add consesus mechanism here
  lmaoChain.AddBlock(genesisBlock)

  fmt.Printf("Deez's balance: %f\n", deez.GetBalance())
  fmt.Printf("Nuts's balance: %f\n", nuts.GetBalance())

  txn2, err := txn.CreateTransaction(deez, nuts, 69)
  if err != nil {
    fmt.Println(err.Error())
  }

  fmt.Println(txn2.Status)

  newBlock := block.CreateBlock([]byte(genesisBlock.GetHash()))
  newBlock.AddTransaction(txn2)
  lmaoChain.AddBlock(newBlock)
}

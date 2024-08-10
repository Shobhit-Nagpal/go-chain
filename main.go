package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/Shobhit-Nagpal/go-chain/internal/account"
	"github.com/Shobhit-Nagpal/go-chain/internal/block"
	"github.com/Shobhit-Nagpal/go-chain/internal/chain"
	"github.com/Shobhit-Nagpal/go-chain/internal/hash"
	"github.com/Shobhit-Nagpal/go-chain/internal/txn"
)

var accounts map[string]*account.Account = make(map[string]*account.Account)

func main() {

	lmaoChain := chain.CreateChain("lmaoChain")
  txnPool := initialize()

	genesisHash := hash.CreateNewHash([]byte(time.Now().String()))

  //Batch transactions to add in blocks
	genesisBlock := block.CreateBlock(string(genesisHash), []*txn.Txn{transaction})
	err = genesisBlock.AddTransaction(transaction)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Add consesus mechanism here
	err = lmaoChain.AddBlock(genesisBlock)
	if err != nil {
		fmt.Println(err.Error())
	}

	txn2, err := txn.CreateTransaction(nuts, deez, 420)
	if err != nil {
		fmt.Println(err.Error())
	}

	newBlock := block.CreateBlock(genesisBlock.GetHash(), []*txn.Txn{txn2})
	err = newBlock.AddTransaction(txn2)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = lmaoChain.AddBlock(newBlock)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(lmaoChain.GetLength())
}

func initialize() []*txn.Txn {
	accountIds := []string{}

	deez := account.CreateAccount()
	nuts := account.CreateAccount()

	spidey := account.CreateAccount()
	flash := account.CreateAccount()

	accounts[deez.Id.String()] = deez
	accounts[nuts.Id.String()] = nuts
	accounts[spidey.Id.String()] = flash
	accounts[flash.Id.String()] = flash

	accountIds = append(accountIds, deez.Id.String())
	accountIds = append(accountIds, nuts.Id.String())
	accountIds = append(accountIds, spidey.Id.String())
	accountIds = append(accountIds, flash.Id.String())

	txnPool := []*txn.Txn{}
	totalAccounts := len(accounts)

	for i := 0; i < 1000; i++ {
		senderIdx := rand.IntN(totalAccounts)
		receiverIdx := rand.IntN(totalAccounts)
		if senderIdx == receiverIdx {
			continue
		} else {
			transaction, err := txn.CreateTransaction(accounts[accountIds[senderIdx]], accounts[accountIds[receiverIdx]], float32(rand.IntN(100)))
			if err != nil {
				fmt.Println(err)
			}
			txnPool = txn.EnqueueTxn(txnPool, transaction)
		}
	}

  return txnPool
}

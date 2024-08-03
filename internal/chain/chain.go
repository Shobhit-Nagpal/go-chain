package chain

import (
	"errors"

	"github.com/Shobhit-Nagpal/go-chain/internal/block"
	"github.com/Shobhit-Nagpal/go-chain/internal/err"
)

type Metadata struct {
	Name         string
	GenesisBlock block.Block
}

type Chain struct {
	Metadata Metadata
	Blocks   map[string]block.Block
	Length   int
}

func CreateChain(name string) Chain {
	return Chain{
		Metadata: Metadata{
			Name: name,
		},
		Blocks: map[string]block.Block{},
		Length: 0,
	}
}

func (c *Chain) GetName() string {
	return c.Metadata.Name
}

func (c *Chain) GetLength() int {
	return c.Length
}

func (c *Chain) GetBlock(blockHash string) (block.Block, error) {
	block, exists := c.Blocks[blockHash]
	if !exists {
		return block, errors.New(err.BLOCK_NONEXISTENT)
	}

	return block, nil
}

func (c *Chain) GetGenesisBlock() block.Block {
	return c.Metadata.GenesisBlock
}

func (c *Chain) AddBlock(block block.Block) error {
	_, exists := c.Blocks[block.GetHash()]
	if exists {
		return errors.New(err.BLOCK_EXISTS_IN_CHAIN)
	}

	if c.Length == 0 {
		c.Metadata.GenesisBlock = block
	}

	c.Blocks[block.GetHash()] = block
	c.Length++

	return nil
}

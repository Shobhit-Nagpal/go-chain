package chain

import "github.com/Shobhit-Nagpal/go-chain/internal/block"

type Chain struct {
  Blocks []block.Block
  Length int
}

package hash

import (
	"crypto/sha256"
)

func CreateNewHash(data []byte) []byte {
  h := sha256.New()

  h.Write(data)
  blockHash := h.Sum(nil)

  return blockHash
}

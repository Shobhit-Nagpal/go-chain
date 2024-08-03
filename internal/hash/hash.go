package hash

import "crypto/sha256"

func CreateNewHash() []byte {
  data := "This is supposed to be some previous data ka hash"
  h := sha256.New()

  h.Write([]byte(data))
  blockHash := h.Sum(nil)

  return blockHash
}

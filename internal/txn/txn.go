package txn

const (
	SUCCESS = iota
	FAILURE
	PENDING
)

type Txn struct {
  Id       string
	Sender   string
	Receiver string
	Status   int
}

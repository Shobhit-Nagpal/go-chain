package txn

const (
	SUCCESS = iota
	FAILURE
	PENDING
)

type Txn struct {
	Sender   string
	Receiver string
	Status   int
}

package types

type Money int64

type Category string

type Currency string

type PAN string

type Status string

const (
	StatusOK Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPRPGRESS"
)

type Card struct {
	ID    int
	PAN   PAN
	Balance Money
	MinBalance Money
	Currency   Currency
	Color  string
	Name    string
	Active   bool
}

type Payment struct {
	ID    int
	Amount Money
	Category Category
	Status Status
}
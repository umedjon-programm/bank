package types

type Money int64

type Category string

type Status string
const(
	StatusOK Status="OK"
	StatusFail Status="FAIL"
	StatusInProgress Status="INPROGRESS"
)
type Payment struct {
	ID     int
	Amount Money
	Category Category
}

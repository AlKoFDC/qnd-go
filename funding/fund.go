package funding

import "errors"

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

type fund struct {
	A string
	b string
}
type fund2 struct {
	A int
}

type Fund2 struct {
	fund
	fund2
	B int
}

func NewFund2() *Fund2 {
	return &Fund2{fund: fund{A: "hooray"}}
}

func NewFund(initBalance int) *Fund {
	return &Fund{
		balance: initBalance,
	}
}

// A regular function returning a pointer to a fund
func NewFund100() (*Fund, error) {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: 100,
	}, errors.New("error")
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}

func getBalance(bp BalanceProvider) int {
	return bp.Balance()
}

type BalanceProvider interface {
	Balance() int
}

type Withdrawer interface {
	Withdraw(int)
}

type FundManager interface {
	BalanceProvider
	Withdrawer
	//foo() error
}

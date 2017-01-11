package funding

import (
	"sync"
	"testing"
)

func TestBalance(t *testing.T) {
	f := NewFund(2)
	t.Log(f.Balance())
	t.Log(f.balance)
	t.Log(f.Balance())
}

func BenchmarkFund(b *testing.B) {
	// Add as many dollars as we have iterations this run
	fund := NewFund(b.N)

	// Burn through them one at a time until they are all gone
	for i := 0; i < b.N; i++ {
		fund.Withdraw(1)
	}

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}

const WORKERS = 10

func BenchmarkWithdrawals(b *testing.B) {
	b.Skip("This will most probably fail. See BenchmarkWithdrawalsService")
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	// Add as many dollars as we have iterations this run
	fund := NewFund(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		// Let the waitgroup know we're adding a goroutine
		wg.Add(1)

		// Spawn off a founder worker, as a closure
		go func() {
			// Mark this worker done when the function finishes
			defer wg.Done()

			for i := 0; i < dollarsPerFounder; i++ {
				fund.Withdraw(1)
			}

		}() // Remember to call the closure!
	}

	// Wait for all the workers to finish
	wg.Wait()

	if fund.Balance() != 0 {
		b.Error("Balance wasn't zero:", fund.Balance())
	}
}

func TestInterface(t *testing.T) {
	var bf balanceFunc = func() int {
		return 7
	}
	t.Log(getBalance(&bf))
}

type balanceFunc func() int

var _ Withdrawer = (*Fund)(nil)
var _ FundManager = (*Fund)(nil)

func (bf balanceFunc) Balance() int {
	return bf()
}

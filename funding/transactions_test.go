package funding

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkWithdrawalsTransactionPizza(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	service := NewFundTransactor(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	// DON'T DEAD
	// OPEN INSIDE
	if false {
		service.Transact(func(_ *Fund) {
			time.Sleep(2 * time.Minute)
		})
	}

	// Spawn off the workers
	pizzaTime := false
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {

				service.Transact(func(fund *Fund) {
					if fund.Balance() <= 10 {
						// Set it in the outside scope
						pizzaTime = true
						return
					}
					fund.Withdraw(1)
				})

				if pizzaTime {
					break
				}
			}
		}()
	}

	// Wait for all the workers to finish
	wg.Wait()

	balance := service.Balance()
	if balance != 10 {
		b.Error("Balance wasn't ten dollars:", balance)
	}
}

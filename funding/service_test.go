package funding

import (
	"sync"
	"testing"
)

func BenchmarkWithdrawalsService(b *testing.B) {
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	service := NewFundService(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	// Spawn off the workers
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				service.Withdraw(1)
			}
		}()
	}

	// Wait for all the workers to finish
	wg.Wait()

	if balance := service.Balance(); balance != 0 {
		b.Error("Balance wasn't zero:", balance)
	}
}

func BenchmarkWithdrawalsServicePizza(b *testing.B) {
	b.Skip("This will most probably fail. See BenchmarkWithdrawalsTransactionPizza")
	// Skip N = 1
	if b.N < WORKERS {
		return
	}

	service := NewFundService(b.N)

	// Casually assume b.N divides cleanly
	dollarsPerFounder := b.N / WORKERS

	// WaitGroup structs don't need to be initialized
	// (their "zero value" is ready to use).
	// So, we just declare one and then use it.
	var wg sync.WaitGroup

	// Spawn off the workers
	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < dollarsPerFounder; i++ {
				// Stop when we're down to pizza money
				if service.Balance() <= 10 {
					break
				}
				service.Withdraw(1)
			}
		}()
	}

	// Wait for all the workers to finish
	wg.Wait()

	if balance := service.Balance(); balance != 10 {
		b.Error("Balance wasn't ten dollars:", balance)
	}
}

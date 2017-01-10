package funding

type FundTransactor struct {
	commands chan TransactionCommand
	fund     *Fund
}

func NewFundTransactor(initialBalance int) *FundTransactor {
	server := &FundTransactor{
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundTransactor) loop() {
	for transaction := range s.commands {
		// Now we don't need any type-switch mess
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
}

func (s *FundTransactor) Balance() int {
	var balance int
	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})
	return balance
}

func (s *FundTransactor) Withdraw(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

// Typedef the callback for readability
type Transactor func(fund *Fund)

// Add a new command type with a callback and a semaphore channel
type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

// Wrap it up neatly in an API method, like the other commands
func (s *FundTransactor) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

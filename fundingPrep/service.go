package funding

import "fmt"

type FundService struct {
	commands chan interface{}
	fund     *Fund
}

func NewFundService(initialBalance int) *FundService {
	server := &FundService{
		// make() creates builtins like channels, maps, and slices
		commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundService) loop() {
	for command := range s.commands {

		// command is just an interface{}, but we can check its real type
		switch command.(type) {

		case WithdrawCommand:
			// And then use a "type assertion" to convert it
			withdrawal := command.(WithdrawCommand)
			s.fund.Withdraw(withdrawal.Amount)

		case BalanceCommand:
			getBalance := command.(BalanceCommand)
			balance := s.fund.Balance()
			getBalance.Response <- balance

		default:
			panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}

func (s *FundService) Balance() int {
	responseChan := make(chan int)
	s.commands <- BalanceCommand{Response: responseChan}
	return <-responseChan
}

func (s *FundService) Withdraw(amount int) {
	s.commands <- WithdrawCommand{Amount: amount}
}

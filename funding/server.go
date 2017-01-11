package funding

import "fmt"

type FundServer struct {
	Commands chan interface{}
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make() creates builtins like channels, maps, and slices
		Commands: make(chan interface{}),
		fund:     NewFund(initialBalance),
	}

	// Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	for command := range s.Commands {

		// command is just an interface{}, but we can check its real type
		switch command := command.(type) {
		case WithdrawCommand:
			s.fund.Withdraw(command.Amount)

		case BalanceCommand:
			balance := s.fund.Balance()
			command.Response <- balance
			close(command.Response)
		default:
			panic(fmt.Sprintf("Unrecognized command: %v %T", command, command))
		}
	}
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}

type validator interface {
	valid() bool
}

func (w WithdrawCommand) valid() bool {
	return w.Amount > 0
}
func (b BalanceCommand) valid() bool {
	return b.Response != nil
}

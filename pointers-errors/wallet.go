package pointerserrors

import "fmt"

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Lowercase fields are considered private outside of this package.

type Wallet struct {
	balance Bitcoin
}

// Calling a function or method creates a copy of the arguments.
// Without defining a pointer to the `Wallet` struct using `*`, we are modifying this copy rather than the argument passed.
// The balance field will then never be changed, and always return the default value.
// Removing the `*` and running `go test` will show that the memory addresses do not match.

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// We don't technically need the pointer reciever here as we aren't modifiying the value on the field.
// However, it's good convention to keep your method receiver types the same for consistency.

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

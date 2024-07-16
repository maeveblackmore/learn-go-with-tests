package pointerserrors

// Lowercase fields are considered private outside of this package.

type Wallet struct {
	balance int
}

// Calling a function or method creates a copy of the arguments.
// Without defining a pointer to the `Wallet` struct using `*`, we are modifying this copy rather than the argument passed.
// The balance field will then never be changed, and always return the default value.

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}

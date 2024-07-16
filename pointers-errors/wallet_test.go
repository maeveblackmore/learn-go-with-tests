package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()

	// `&` obtains the memory address where the variable `wallet.balance` is stored.
	// `%p` prints memory addresses in base 16 notation with a leading `0x` to represent hexadecimal.
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d want %d... %#v", got, want, wallet.balance)
	}
}

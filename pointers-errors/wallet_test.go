package pointerserrors

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// `&` obtains the memory address where the variable `wallet.balance` is stored.
		// `%p` prints memory addresses in base 16 notation with a leading `0x` to represent hexadecimal.
		fmt.Printf("address of balance in test is %p \n", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
	})
}

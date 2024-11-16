package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	// assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
	// 	t.Helper()
	// 	got := wallet.Balance()
	//
	// 	if got != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// }
	//
	// assertError := func(t testing.TB, got error, want string) {
	// 	t.Helper()
	//
	// 	if got == nil {
	// 		t.Fatal("didn't get an error but wanted one")
	// 	}
	//
	// 	if got.Error() != want {
	// 		t.Errorf("got %s, want %s", got, want)
	// 	}
	// }

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		assertBalance(t, wallet, Bitcoin(20))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insuficient funds", func(t *testing.T) {
		startingBallance := Bitcoin(20)
		wallet := Wallet{startingBallance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBallance)
		assertError(t, err, ErrInsuficientFunds)
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didn't want one!")
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one!")
	}

	if got != ErrInsuficientFunds {
		t.Errorf("got %q, want %q", got, want)
	}
}

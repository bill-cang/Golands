// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"sync"
	"testing"
	"gopl.io/ch9/bank2"
)

func TestBank(t *testing.T) {
	// Save [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := bank.GetMoney(), (1000+1)*1000/2; got != want {
		t.Errorf("GetMoney = %d, want %d", got, want)
	}
}

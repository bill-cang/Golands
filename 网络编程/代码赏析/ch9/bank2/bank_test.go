// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"sync"
	"testing"
)

func TestBank(t *testing.T) {
	// Save [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 103; i++ {
		n.Add(1)
		go func(amount int) {
			Save(amount)
			n.Done()
		}(i)
	}
	n.Wait()

	if got, want := GetMoney(), (100+1)*100/2; got != want {
		t.Errorf("GetMoney = %d, want %d", got, want)
	}
}

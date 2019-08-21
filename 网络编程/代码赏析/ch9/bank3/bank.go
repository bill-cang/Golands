// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 263.

// Package bank provides a concurrency-safe single-account bank.
package bank

//!+
import "sync"

var (
	mu         sync.Mutex // guards totalMoney
	totalMoney int
)

func Deposit(amount int) {
	mu.Lock()
	totalMoney = totalMoney + amount
	mu.Unlock()
}

func GetMoney() int {
	mu.Lock()
	b := totalMoney
	mu.Unlock()
	return b
}

//!-

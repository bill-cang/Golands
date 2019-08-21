package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"one", "two", "three"}
	dom1(strs)
	fmt.Println("================")
	dom2(strs)
}

func dom1(strs []string) {
	for _, s := range strs {
		go func(s string) {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}(s)
	}
	time.Sleep(3 * time.Second)
}

func dom2(strs []string) {
	for _, s := range strs {
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("%s ", s)
		}()
	}
	time.Sleep(3 * time.Second)
}

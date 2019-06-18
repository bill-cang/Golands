package main

import (
	"fmt"
	"testing"
)

func TestGetKeyIdioms(t *testing.T) {
	str := "一树梨花压海棠"
	for _, v := range str {
		spf := fmt.Sprintf("%c", v)
		GetKeyIdioms(spf)
		//fmt.Println(idioms)

	}
}



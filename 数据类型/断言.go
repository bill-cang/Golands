package main

import "fmt"

func main() {
	var name interface{}
	name="ckx"
	fmt.Println(name.(string))
}

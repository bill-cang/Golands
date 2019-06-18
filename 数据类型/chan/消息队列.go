package main

import "fmt"

type handle interface {
	eat(foot string) bool
	breed()
}

func newhandle() handle {
	return &gdog{
		kind: "demu",
		age:  12,
		name:"hua",
	}
}

type dog struct {
	kind string
	age  int
}

type gdog struct {
	dog
	name string
}

func (d *gdog) eat(foot string) bool {
	res := false
	switch foot {
	case "鸡肉":
		return true
	case "树叶":
		return false
	}
	return res
}

func (d *gdog) breed() {
	fmt.Println("make the little dog")
}

func main() {
	hd := newhandle()
	bo := hd.eat("shuye")
	fmt.Println(bo)


}

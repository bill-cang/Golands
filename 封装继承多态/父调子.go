package main

import (
	"fmt"
	"reflect"
)

type A struct {

}

func (self A) say() {
	println(self.Name())
}

func (self A) sayReal(child interface{}) {
	ref := reflect.ValueOf(child)
	method := ref.MethodByName("Name")
	if (method.IsValid()) {
		r := method.Call(make([]reflect.Value, 0))
		fmt.Println(r[0].String())
	} else {
		// 错误处理
	}
}

func (self A) Name() string {
	return "I'm A"
}

type B struct {
	A
}

func (self B) Name() string {
	return "I'm B"
}


type C struct {
	A
}

func main() {
	b := B{}
	b.say()         //I'm A
	b.sayReal(b)    //I'm B

	c := C{}
	c.say()         //I'm A
	b.sayReal(b)    //I'm A
}
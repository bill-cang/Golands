//https://studygolang.com/articles/2762
package main

import (
	"fmt"
	"reflect"
)

type boy struct {
	Name string
	age  int
}

type human interface {
	SayName()
	SayAge()
}

func (this *boy) SayName() {
	fmt.Println(this.Name)
}

func (this *boy) SayAge() {
	fmt.Println(this.age)
}

func main() {
	// 定义接口变量
	var i human
	// 初始化对象，jown持有对象指针。
	jown := &boy{
		Name: "jown",
		age: 15,
	}
	// 因为boy实现了human中的方法，所以它实现了human接口。
	// 这时，i就指向jown对象。
	i = jown

	// 通过反射获取接口i 的类型和所持有的值。
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// ... 后续操作
	fmt.Println(t,"\t",v,"\t")

	// 获取i所指向的对象的类型
	structType := t.Elem()
	// 获取对象的名字
	structName := structType.Name()
	fmt.Println(structName)

	//利用t的MethodByName() Method方法:
	method, _ := t.MethodByName("SayAge")
	fmt.Println(method)
}

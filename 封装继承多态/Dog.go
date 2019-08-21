package main

import "fmt"

type Dog struct {
	name string
	sex  int
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog *Dog) GetName() string {
	return dog.name
}

func (dog *Dog) SetSex(sex int) {
	dog.sex = sex
}

func (dog *Dog) GetSex() int {
	return dog.sex
}

func (dog *Dog) doBark(bark string) string {
	return bark
}

func main() {
	dog4 := new(Dog)

	fmt.Printf("%T",dog4)

}

//==============创建对象的四种方法===================
//4通过new得到结构体【类】的对象的指针。然后通过指针设置对象的属性、方法
func createDog4() {
	dog4 := new(Dog)
	dog4.name = "德牧"
	dog4.SetSex(0)
	bark := dog4.doBark("汪汪……汪汪汪……")
	fmt.Println(dog4.GetName(), dog4.sex, bark)
}
//3创建对象时选择性的给对象赋值，然后同过set方法设置其属性
func createDog03() {
	dog3 := Dog{name: "阿拉斯"}
	dog3.SetName("阿拉斯加")
	dog3.SetSex(1)
	bark := dog3.doBark("汪汪汪")
	fmt.Println(dog3.name, dog3.sex, bark)
}
//2创建对象时就直接给对象参数赋值
func createDog02() {
	dog2 := Dog{"萨摩", 0}
	bark := dog2.doBark("汪汪")
	fmt.Println(dog2.name, dog2.sex, bark)
}

//1创建一个空对象，然后再给对象属性、方法赋值
func createDog01() {
	dog1 := Dog{}
	dog1.name = "哈士奇"
	dog1.sex = 1
	bark := dog1.doBark("呜呜……")
	fmt.Println(dog1.name, dog1.sex, bark)
}

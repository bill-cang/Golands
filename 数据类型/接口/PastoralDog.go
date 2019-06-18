package main

import "fmt"

type PasDog struct {
	DogEntity
	dfc *DogFc
}

type student struct {
	name string
	age int
}

func (this *PasDog)Eat(foot string) bool  {
	res:=false
	switch foot {
	case "树叶":
		res=false
	case "鸡腿":
		res=true
	default:
		res=false
	}
	return res
}

func main() {
	var stu = student{
		name:"张报纸",
		age :12,
	}

	fmt.Println(stu)

	pdg := new(PasDog)
	pdg.Name="八公"
	pdg.Age=13
	eat := pdg.Eat("树叶")

	fmt.Println(eat)

}

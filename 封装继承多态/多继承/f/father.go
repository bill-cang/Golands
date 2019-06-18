package f

import "fmt"

type Father struct {
	Name string
	Age int
	hobby []string
}

func (f *Father)Eat(some string)  {
	fmt.Println("Father's func",some)
}

func (f *Father)Dringk(some string)  {
	fmt.Println("Father's func",some)
}

func (f *Father)Card(){
	fmt.Println("Father's func",)
}

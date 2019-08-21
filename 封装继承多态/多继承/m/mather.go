package m

import "fmt"

type Mather struct {
	Name string
	Age int
	hobby []string
}

func (m *Mather)Eat(some string)  {
	fmt.Println("Mather's func",some)

}

func (m *Mather)Dringk(some string)  {
	fmt.Println("Mather's func",some)

}

func (m *Mather)Cook(){
	fmt.Println("Mather's func")

}

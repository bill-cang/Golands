package main

import "fmt"

type PolicDog struct {
	//继承父类属性和方法
	Dog
	//独有属性
	poNum int
}

//重写父类方法
func (pd *PolicDog) doBark(bark string) string {
	return pd.name + bark
}

//独有方法
func (pd *PolicDog) doJob(job string) {
	fmt.Println("警犬", pd.poNum, "执行", job, "任务！")
}

func (pd *PolicDog)SetPoNum(pn int){
	pd.poNum=pn
}

func (pd *PolicDog)GetPoNum()int{
	return pd.poNum
}

func main() {
	poDog := PolicDog{}
	poDog.name="啸天"
	poDog.sex=0
	poDog.SetPoNum(001)
	poDog.doBark("汪")
	poDog.doBark("追捕")
	fmt.Print(poDog)

}

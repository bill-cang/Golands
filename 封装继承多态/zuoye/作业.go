package mian

import "fmt"

/*
#接口->父类实现->多种子类实现->多态
·定义接口IPerson，定义吃喝睡三个抽象方法；
·定义一个IPerson的实现类Worker即劳动者，拥有劳动方法Work()(output string)其中output是其工作产出，和休息方法Rest()；
·继承Worker实现三个不同职业的子类：程序员Coder、老师Teacher、农民Farmer，并创建一个Worker的集合；
·从命令行循环接收【今天星期几】的输入，如果是周一到周五全体工作，如果是周六日程序员工作其他人休息；
*/
type IPerson interface {
	eat(foot string)
	drink()
	sleep()
}

type DoThing interface {
	
}

type Person struct {
	zhiye string
}

func (w *Person)SetZhiYe(zy string){
	w.zhiye=zy
}

func (w *Person)GetZhiYe() string{
	return w.zhiye
}

func (w *Person)eat(foot string){
	fmt.Printf("%T的午餐%s",w.zhiye,foot)
}

func (w *Person)drink(water string){
	fmt.Printf("%T的饮料%s",w.zhiye,water)
}

func (w *Person)sleep(){
	fmt.Println("十二点啦，休息！")
}

func main() {

}


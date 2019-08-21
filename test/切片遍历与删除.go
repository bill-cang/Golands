package main

import "fmt"

type stud struct {
	Name string
	Age  int
}

func main() {
	pase_student()
}
func pase_student() {
	m := make(map[string]*stud)
	stus := []stud{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//错误 stu 永远指向一个地址
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	//正确
	for k, stu := range stus {
		m[stu.Name] = &stus[k]
	}
	fmt.Println(m)
}

//切片删除
func Remove(s []string, value interface{}) error {
	for i, v := range s {
		if value == v {
			s = append(s[:i], s[i+1:]...)
			return nil
		}
	}
	return nil
}

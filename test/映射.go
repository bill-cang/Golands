//go中Map的键竟然可以重复！！
package main

type Goods struct {
	id    string
	cid   int
	name  string
	price float32
}

var cidmap map[int][]*Goods
var pamp map[int][]string

func init()  {
	pamp := make(map[int][]string)

}
func main() {


}

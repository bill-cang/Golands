package main
/*交替打奇偶数
*/
import (
	"fmt"
	"sync"
)

var (
	odd  []  int
	even [] int

	chA = make(chan int, 1)
	chB = make(chan int, 1)

	wgx =sync.WaitGroup{}
)


func init() {
	odd = make([]int, 0)
	even = make([]int, 0)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even = append(even, i)
			continue
		}
		odd = append(odd, i)
	}
}

func main() {
	chA<-0
	wgx.Add(2)
	go pritEven();
	go pritOdd();
	wgx.Wait()
}

func pritEven() {
	for _, v := range even {
		<-chA
		fmt.Println(v)
		chB<-v
	}
	close(chB)
	wgx.Done()
}

func pritOdd() {
	for _, v := range odd {
		<-chB
		fmt.Println(v)
		chA<-v
	}
	close(chA)
	wgx.Done()
}

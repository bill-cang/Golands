// 07.多任务资源竞争问题
package main


import (
	"fmt"
	"time"
)


func PrinterVII(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1VII() {
	PrinterVII("今生注定我爱你")
}

func person2VII() {
	PrinterVII("FUCKOFF")
}

func main() {
	go person1VII()
	go person2VII()

	for {
		time.Sleep(time.Second)
	}

}




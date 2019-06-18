package A基础概念

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func CheckErrorC(err error){
	if  err !=nil{
		fmt.Println("网络错误",err.Error())
		os.Exit(1)
	}
}

func main() {
	conn,err :=net.Dial("tcp","127.0.0.1:8898")//建立TCP服务器
	defer conn.Close() //延迟关闭网络连接
	CheckErrorC(err)//检查错误

	//创建一个黑窗口（标准输入）的读取器
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for{
		lineBytes, _, err := reader.ReadLine()
		CheckErrorC(err)
		conn.Write(lineBytes)
		fmt.Println("发送消息")

		/*接收消息*/
		n, err := conn.Read(buffer)
		CheckErrorC(err)

		msg := string(buffer[:n])
		fmt.Println("收到服务端消息：",msg)
	}


}

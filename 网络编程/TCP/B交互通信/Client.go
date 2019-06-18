package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func CheckErrorC(err error) {
	if err != nil {
		fmt.Println("网络错误", err.Error())
		os.Exit(1)
	}
}

func MessageSend(conn net.Conn) {
	var msg string
	reader := bufio.NewReader(os.Stdin) //读取键盘输入

	for {
		lineBytes, _, _ := reader.ReadLine() //读取一行
		msg = string(lineBytes)              //键盘输入转化为字符串

		if msg == "exit" {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

		_, err := conn.Write([]byte(msg)) //输入写入字符串
		if err != nil {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8898") //建立TCP服务器
	CheckErrorC(err)                               //检查错误
	defer conn.Close()

	//发送消息中有阻塞读取标准输入的代码
	//为了避免阻塞住消息的接收，所以把它独立的协程中
	go MessageSend(conn)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		CheckErrorC(err)

		msg := string(buffer[:n])
		fmt.Println("收到服务器消息", msg)

		if msg=="fuckoff"{
			break
		}

	}

	fmt.Println("连接已断开")
}

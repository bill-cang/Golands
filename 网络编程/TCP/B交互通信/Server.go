package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckErrorS(err error) {
	if err != nil {
		fmt.Println("网络错误", err.Error())
		os.Exit(1)
	}
}

func Processinfo(conn net.Conn) {
	buffer := make([]byte, 1024) //开创缓冲区
	defer conn.Close()           //关闭连接

	for {
		n, err := conn.Read(buffer) //读取数据
		CheckErrorS(err)

		if n != 0 {
			//拿到客户端地址
			remoteAddr := conn.RemoteAddr()
			msg := string(buffer[:n])
			fmt.Println("收到消息",msg, "来自", remoteAddr)

			if strings.Contains(msg,"钱") {
				conn.Write([]byte("fuckoff"))
				break
			}
			conn.Write([]byte("已阅："+msg))
		}
	}

}

func main() {
	//建立TCP服务器
	listener, err := net.Listen("tcp", "127.0.0.1:8898")
	CheckErrorS(err)
	defer listener.Close() //关闭网络

	fmt.Println("服务器正在等待")

	for {
		conn, err := listener.Accept() //新的客户端连接
		CheckErrorS(err)

		//处理每一个客户端
		go Processinfo(conn)
	}

}

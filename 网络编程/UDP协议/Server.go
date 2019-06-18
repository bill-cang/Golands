package main

import (
	"os"
	"fmt"
	"net"
)

//处理错误信息
func checkError(err error) {
	if err != nil { //指针不为空
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
}

func receiveUDPMsg(udpConn *net.UDPConn) {
	//声明30字节的缓冲区
	buffer := make([]byte, 30)

	//从udpConn读取客户端发送过来的数据，放在缓冲区中(阻塞方法)
	//返回值：n=读到的字节长度,remoteAddr=客户端地址,err=错误
	n, remoteAddr, err := udpConn.ReadFromUDP(buffer) //从udp接收数据
	if err != nil {
		fmt.Println("Error", err.Error()) //打印错误信息
		return
	}

	fmt.Printf("接收到来自%v的消息：%s", remoteAddr,string(buffer[0:n]))

	//向远程地址中写入数据
	_, err = udpConn.WriteToUDP([]byte("hao nimei"), remoteAddr)
	checkError(err)
}


func main() {

	//解析IP和端口得到UDP地址
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8848")
	checkError(err)

	//在解析得到的地址上建立UDP监听
	udpConn, err := net.ListenUDP("udp", udp_addr)
	defer udpConn.Close() //关闭链接
	checkError(err)

	//从udpConn中接收UDP消息
	receiveUDPMsg(udpConn) //收消息
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

var (
	onlinAdrrMap    = make(map[string]net.Conn)
	chServerTurnOff = make(chan string, 0)
)

func sHandleErr(err error, where string) {
	if err != nil {
		fmt.Printf("%v:%v", where, err)
		os.Exit(1)
	}
}

func main() {
	//建立监听
	listener, e := net.Listen("tcp", "127.0.0.1:8848")
	sHandleErr(e, "net.Listen")

	defer func() {
		for _, conn := range onlinAdrrMap {
			conn.Write([]byte("【server】！！！！注意：服务即将在二十秒后关闭，进行维护！"))
		}
		for i := 0; i < 21; i++ {
			fmt.Println("【系统即将关闭】",20-i)
			time.Sleep(1*time.Second)
		}
		//<-time.NewTicker(20 * time.Second).C
		listener.Close()
	}()
	fmt.Println("服务器已开启============================================")

	go func() {
		//上线提醒
		for {
			//获取链接监听的客服端
			conn, e := listener.Accept()
			sHandleErr(e, "listener.Accept()")
			addr := conn.RemoteAddr().String()

			//给已在线客户端发送新上线通知
			for _, conn := range onlinAdrrMap {
				conn.Write([]byte("***" + addr + "上线啦"))
			}

			//将新上线客户端存入在线集合
			onlinAdrrMap[addr] = conn

			//服务器监控打印
			fmt.Printf("%v已上线,在线数%d\n", addr, len(onlinAdrrMap))

			//开辟一个独立的协程获取客户端消息
			go getMassageForClent(conn)

		}
	}()

	//开辟一条协程关闭服务器
	go func() {
		for {
			getServerInput()
		}
	}()

	<-chServerTurnOff
}

func getServerInput() {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	str := string(bytes)
	if strings.Compare(str, "turnoff") == 0 {
		chServerTurnOff <- str
	}
}

func getMassageForClent(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != io.EOF {
			sHandleErr(err, "conn.Read(buffer)")
		} else {
			//从在线Map中删除下线客户端
			delete(onlinAdrrMap, addr)
			//给在线客户端发送下线客户端消息
			for _, conn := range onlinAdrrMap {
				conn.Write([]byte("***"+addr + "已下线。"))
			}
			//conn.Close()
			break
		}

		msage := string(buffer[:n])
		if strings.Index(msage, "@@") == 0 {
			//群发
			msage=msage[2:len(msage)]
			for key, conn := range onlinAdrrMap {
				if key != addr {
					conn.Write([]byte("***"+addr+":"+msage))
				}
			}
		} else if strings.Compare("$$$", msage) == 0 {
			fmt.Println("【Client故障】" + msage)
		} else {
			//点对点发消息
			targetAddr := msage[:15]
			sendMage := msage[16:len(msage)]
			//targetAddr.(net.Addr)
			sendConn := onlinAdrrMap[targetAddr]
			sendConn.Write([]byte("***" + addr + ":" + sendMage))
		}

	}

}

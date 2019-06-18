package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	chClientQuit = make(chan string)
	conn         net.Conn
	e            error
)

func cHandleErr(conn net.Conn, err error, where string) {
	if err != nil {
		fmt.Printf("【系统错误】%v:%v", where, err)
		conn.Write([]byte("$$$" + where + err.Error()))
		/*这里的非正常关闭会引起服务器的异常：
			read tcp 127.0.0.1:8848->127.0.0.1:58625: wsarecv: An existing connection was forcibly closed by the remote host.
		*/
		//os.Exit(1)
	}
}
func main() {
	//拨号
	conn, e = net.Dial("tcp", "127.0.0.1:8848")
	cHandleErr(conn, e, "net.Dial")
	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + "您已上线")

	//开辟一个协程接收服务器消息
	go getMassageForServer(conn)

	//开辟一条协程发消息
	go sendMsgToTarget(conn)
	//阻塞主程序，待到给出推出信号
	<-chClientQuit
}

func sendMsgToTarget(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		//cHandleErr(err,"reader.ReadLine()")

		msg := string(bytes)
		msg = strings.TrimSpace(msg)
		if len(msg) != 0 {
			if strings.Contains(msg, "over") {
				chClientQuit <- "over"
				return
			}

			//防止地址后面紧跟的“:”中文【待优化】
			msg = strings.Replace(msg, "：", ":", 1)
			_, err := conn.Write([]byte(msg))
			cHandleErr(conn, err, "onn.Write")
		} else {
			fmt.Println("***消息不能为空！")
		}

	}
}

func getMassageForServer(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		//读取消息
		n, e := conn.Read(buffer)
		cHandleErr(conn, e, "conn.Read")
		msg := string(buffer[:n])
		fmt.Println(msg)
		//开辟一条协程去处理，关机期间还能通信。
		go ServerWaring(msg)

	}
}

func ServerWaring(msg string) {
	if strings.Compare(msg, "【server】！！！！注意：服务即将在二十秒后关闭，进行维护！") == 0 {
		<-time.NewTicker(20*time.Second).C
		chClientQuit<-msg
	}
}

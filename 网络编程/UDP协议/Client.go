package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	chSend = make(chan string, 10)
	chBack = make(chan string, 10)
	chOver = make(chan string,1)
	conn   net.Conn
	star int = 0 //首次链接成功
)

func main() {

	go func() {
		for {
			conn = GetConForNet("udp", "127.0.0.1:8848")
			select {
			case sendStr := <-chSend:
				if strings.Compare(sendStr,"ClientStar")==0 {
					if star==0{
						fmt.Println("服务链接成功！")
						fmt.Print("Client：")
					}else{
						fmt.Print("Client：")
					}
				}
				var massage string
				fmt.Scanf("%s", massage)
				chSend <- massage
				SendMassage()
				star++
				<-chBack

			case massage := <-chBack:
				GetServerMassage(conn, massage)
			}

		}
	}()

	<-time.NewTicker(3*time.Second).C
	<-chOver //fatal error: all goroutines are asleep - deadlock!
}

//请求服务器链接
func GetConForNet(netWork string, address string) (net.Conn) {
	//请求连接服务器，得到连接对象
	conn, err := net.Dial(netWork, address)
	defer conn.Close()
	if err != nil {
		fmt.Println("网络连接出错")
		os.Exit(1)
	}
	chSend<-"ClientStar"
	return conn
}

func GetServerMassage(conn net.Conn, meaasge string) {

	buffer := make([]byte, 30)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取消息错误：err=", err)
		os.Exit(1)
	}
	fmt.Println("Server：", string(buffer[:n]))

}

func SendMassage() {
	//向连接中写入消息
	massage :=<-chSend
	if strings.Compare(massage,"over") ==0{
		chOver<-massage
	}
	conn.Write([]byte(massage))
	fmt.Println( massage)
}

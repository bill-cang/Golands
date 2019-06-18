package A基础概念

import (
	"fmt"
	"net"
	"os"
)

func CheckErrorS(err error){
	if  err !=nil{
		fmt.Println("网络错误",err.Error())
		os.Exit(1)
	}
}

func  Processinfo(conn net.Conn){
	buf :=make([]byte,1024)//开创缓冲区
	defer conn.Close() //关闭连接

	//与客户端源源不断地IO
	for{
		n,err:=conn.Read(buf) //读取数据
		if err !=nil{
			break
		}

		if n !=0{
			msg := string(buf[:n])
			fmt.Println("收到消息",msg)
			conn.Write([]byte("已阅："+msg))

			if msg=="分手吧"{
				break
			}
		}
	}

}


func main() {
	//建立TCP服务器
	listener,err:=net.Listen("tcp","127.0.0.1:8898")
	defer listener.Close() //关闭网络
	CheckErrorS(err)

	//循环接收，来者不拒
	for  {
		//接入一个客户端
		conn,err:= listener.Accept() //新的客户端连接
		CheckErrorS(err)

		//为每一个客户端开一条独立的协程与其IO
		go  Processinfo(conn)
	}


}

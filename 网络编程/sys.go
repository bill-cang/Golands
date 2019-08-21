package main

import (
	"fmt"
	"net"
)

func main() {


	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", addrs)
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	/*	var sysIn string
		fmt.Println("请输入：")
		fmt.Scanf("%s",&sysIn)
		fmt.Println(sysIn)*/

}

package main

import "fmt"

func main() {
	rds:=InstanceCache()
	fmt.Printf("**%T\t%v\n",rds,rds)

	//string
	StrFactory(rds)


}



package main

import "fmt"

func HandleError(err error,where string)  {
	if err != nil {
		fmt.Println("***",where,":",err)
	}
}

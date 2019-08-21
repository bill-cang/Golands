package main

import "fmt"

/*
map[balance:0.00000000 blocks:385499 connections:16 difficulty:5036.518453240553 errors: keypoololdest:1563517841 keypoolsize:101 paytxfee:0.00000000 protocolversion:170013 proxy: relayfee:0.00000100 testnet:false timeoffset:0 version:3030150 walletversion:60000]
*/

func main() {
	mmp := make(map[string]interface{}, 0)
	mmp["balance"] = 0.00000000
	mmp["blocks"] = 385499
	mmp["difficulty"] = 5036.518453240553
	mmp["errors"] = 16

	fmt.Printf("%T\n",mmp)


	for k, v := range mmp {
		fmt.Println(k, v)
	}
}

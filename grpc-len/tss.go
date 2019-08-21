package main

import (
	"flag"
	"fmt"
)

func main() {

	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	if *tls {
		fmt.Println(tls)
	}

}

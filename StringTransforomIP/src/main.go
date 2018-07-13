package main

import (
	"os"
	"fmt"
	"net"
)

func main(){
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"usage: %s ip-adde \n",os.Args[0])
		os.Exit(0)
	}
	name :=os.Args[1]
	addr :=net.ParseIP(name)
	if addr == nil{
		fmt.Println("invalid address")
	}else {
		fmt.Println("the address is" , addr.String())
	}
	os.Exit(0)
}
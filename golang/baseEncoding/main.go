package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	hello:="你好世界"
	debyte:=base64.StdEncoding.EncodeToString([]byte(hello))
	fmt.Println(debyte)
	enbyte,err:=base64.StdEncoding.DecodeString(string(debyte))
	if err!=nil {
		fmt.Println(err.Error())
	}

	if hello!=string(enbyte){
		fmt.Println("hello is not equal to enbyte")
	}
	fmt.Println(string(enbyte))
}

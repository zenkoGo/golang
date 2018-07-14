package main

import (
	"crypto/md5"
	"io"
	"fmt"
)

func main(){
	h := md5.New()
	io.WriteString(h,"xuyaojiamidemima")
	pwmd5:=fmt.Sprintf("%x",h.Sum(nil))
	fmt.Printf(pwmd5+"\n")
	salt1:="@#$%"
	salt2:="^&*("
	io.WriteString(h,salt1)
	io.WriteString(h,"abc")
	io.WriteString(h,salt2)
	io.WriteString(h,pwmd5)
	last:=fmt.Sprintf("%x",h.Sum(nil))
	fmt.Printf(last)
}


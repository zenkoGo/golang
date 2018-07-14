package main

import (
	"crypto/sha256"
	"io"
	"fmt"
	"crypto/sha1"
	"crypto/md5"
)

func main(){
	h:=sha256.New()
	io.WriteString(h,"his money is twice tained : tain your money and taind mine ")
	fmt.Printf("% x\n",h.Sum(nil))
	h=sha1.New()
	io.WriteString(h,"his money is twice tained : tain your money and taind mine")
	fmt.Printf("% x\n",h.Sum(nil))
	h=md5.New()
	io.WriteString(h,"his money is twice tained : tain your money and taind mine")
	fmt.Printf("% x\n",h.Sum(nil))
}
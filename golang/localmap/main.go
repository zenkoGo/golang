package main

import (
	"fmt"
)

var locales map[string]map[string]string

func main(){
	locales=make(map[string]map[string]string,2)
	en:=make(map[string]string,10)
	en["pea"]="pea"
	en["bean"]="bean"
	locales["en"]=en
	cn:=make(map[string]string)
	cn["pea"]="豌豆"
	cn["bean"]="毛豆"
	cn["how old"]="我今年%d 岁了\n"
	locales["zh-cn"]=cn
	lang:="zh-cn"
	fmt.Println(msg(lang,"pea"))
	fmt.Println(msg(lang,"bean"))
	fmt.Printf(msg(lang,"how old"),30)

}


func msg(local,key string)string{
	if v,ok:=locales[local];ok {
		if v2,ok:=v[key];ok{
			return v2
		}
	}
	return ""
}

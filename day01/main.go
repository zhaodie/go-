package main

import (
	"fmt"
	"strconv"
	"strings"
)

var str1 string
var a1 int
var a2 int

func main() {
	fmt.Print("请输入(格式为192.1.1.1/192.168.1.1/24):")
	fmt.Scanf("%s",&str1)
	a1 = strings.LastIndex(str1, ".")
	a2 = strings.Index(str1, "/")
	str2 := []rune(str1)
	zw := string(str2[a2+1:])
	if zw == "24"{
		for i := 0; i < 255; i++{
			fmt.Println(string(str2[:a1+1])+strconv.Itoa(i))
		}
	}else{
		fmt.Println(str1)
	}
}

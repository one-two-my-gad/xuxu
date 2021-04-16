package main

import (
	"fmt"
)

//整数

func main() {

	//十进制
	var i1 int = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) //把十进制数转换至二进制
	fmt.Printf("%o\n", i1) //把十进制转换至八进制
	fmt.Printf("%x\n", i1) //把十进制转换至 十六进制

	//八进制 以0开头
	i2 := 077
	fmt.Printf("%o\n", i2)

	//十六进制 以0x开头
	i3 := 0x1234567
	fmt.Printf("%x\n", i3)

	fmt.Printf("%T\n", i3) //%T为打印变量类型

	i4 := int8(9) //明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}

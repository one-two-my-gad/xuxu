package main

import (
	"fmt"
)

//浮点数

func main() {
	f1 := 1.23456
	//	math.MaxFloat32 //float32最大值
	fmt.Printf("%t\n", f1) //默认go语言中小数都是float64类型

	f2 := float32(1.23456)
	fmt.Printf("%t\n", f2) //显示声明float32类型
	// f1 = f2   //float32类型的值不能直接赋值给float64类型的变量

}

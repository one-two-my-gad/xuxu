package main

import "fmt"

//常量 作用于全局 定义后常量不能修改

const pi = 3.1415926

//批量声明变量
const (
	status   = 200
	notfount = 404
)

// 批量声明常量时 如果某一行声明后没有赋值 默认就和上一行声明 一致
const (
	n1 = 100
	n2
	n3
)

//iota 作为枚举值
const (
	a1 = iota //0
	a2        //1
	a3        //2
)


//插队

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

// 多个常量声明在一行

const (
	d1, d2 = iota + 1, iota + 2 //d1:1  d2:2

	d3, d4 = iota + 1, iota + 2 //d3:2 d4:3
)



//定义数量级

const (
	_ = iota 
	KB = 1 << (10 * iota) // 1左移10位：100000000000 二进制转换为十进制 为2的十次方=1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//pi =12345
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2", a2)
	fmt.Println("a3", a3)

	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)

}

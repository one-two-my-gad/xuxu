package main

import "fmt"

//批量声明变量
var (
	name string
	age  int
	isOk bool
	s3   string
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	s3 = "哈哈哈"
	fmt.Println(isOk)
	fmt.Println(age)
	fmt.Printf("name:%s", name)
	fmt.Println(s3)
}

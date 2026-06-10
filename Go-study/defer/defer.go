package main

import "fmt"

func main() {
	//defere有点类似于cpp里的析构函数
	defer fmt.Println("main end1")
	defer fmt.Println("main end2") //以压栈形式，先进后出，所以是end2先出

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

}

//defer比return还晚执行
//在当前函数的生命周期完全结束之后才出栈

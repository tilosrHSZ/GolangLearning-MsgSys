package main

import "fmt"

const (
	BEIJING = iota //第一行初始值是0，第二行初始值是1，第三行初始值是2，以此类推
	SHANGHAI
	GUANGDONG
)

func main() {

	//常量
	const length int = 10
	fmt.Println("length =", length)

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("GUANGDONG =", GUANGDONG)

}

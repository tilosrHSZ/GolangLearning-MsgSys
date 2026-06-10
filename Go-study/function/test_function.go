package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("----foo1----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//匿名返回多个返回值
func foo2(a string, b int) (int, int) {
	fmt.Println("----foo2----")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

//有形参名返回多个返回值
func foo3(a string, b int) (r1 int, r2 int) { //此处int可以放在后面r1, r2 int
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	r1 = 1000
	r2 = 2000

	return
}

func main() {

	c := foo1("abc", 200)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("qawed", 123)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	ret1, ret2 = foo3("1asd", 3123)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)
}

package main

import "fmt"

func main() {

	//默认值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of a = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s type of bb = %T\n", bb, bb)

	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of a = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s type of cc =%T\n", cc, cc)

	//最常用的构建变量的方法,只能够用在函数内部
	d := 100
	fmt.Println("d = ", d)
	fmt.Printf("type of d = %T\n", d)

	//声明多个变量
	var e, f, g int = 100, 200, 300
	fmt.Println("e = ", e, "f = ", f, "g = ", g)

	var kk, ll = 100, "abcd"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}

package main

import "fmt"

func main() {
	s := []int{1, 2, 3} //len = 3, cap = 3

	s1 := s[0:2] //左闭右开，取s里的1，2

	fmt.Println(s1)

	s1[0] = 100
	//会发现两个都被改变了，相当于是拷贝了内存地址
	fmt.Println(s)
	fmt.Println(s1)

	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) //s2 = [0,0,0]

	copy(s2, s)
	//深拷贝，开辟了一块新的地址空间进行拷贝，这样修改s2就不会改变s
	fmt.Println(s2)

	s2[0] = 200

	fmt.Println(s)
	fmt.Println(s2)

}

package main

import (
	"fmt"
	"time"
)

func main1() {
	/*
		//go创建匿名进程
		go func() {
			fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				//rumtime.Goexit()//退出当前进程
				fmt.Println("B")
			}()

			fmt.Println("A")
		}()
	*/
	//带参进程
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)
	//此时外主进程main不会收到副进程的返回值
	//若要返回，使用channel

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}

}

package main

import "fmt"

func main2() {
	defer fmt.Println("main goroutine end")
	c := make(chan int) //create channel

	go func() {
		defer fmt.Println("sub goroutine end")

		fmt.Println("sub goroutine run")

		c <- 123 //channel赋值
	}()
	//channel会控制两个进程并行
	//无缓冲的channel哪边先跑到和c有关的语句就阻塞等待另一边传输
	fmt.Println("main goroutine run")
	num := <-c //channel使用
	fmt.Println(num)

}

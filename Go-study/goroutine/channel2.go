package main

import (
	"fmt"
	"time"
)

func main3() {
	c := make(chan int, 3) //channel with cache

	fmt.Println("len(c) = ", len(c), "cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub goroutine end")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("sub goroutine is sent", i, " len(c) = ", len(c), "cap(c) = ", cap(c))

		}

	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	fmt.Println("main goroutine end")
}

//4已经大于了channel的范围，在输入第四个数的时候会阻塞，直到后面main里取出了一个值才会在sub里填入

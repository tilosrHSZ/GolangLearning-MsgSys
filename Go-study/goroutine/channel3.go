package main

import "fmt"

func main4() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		close(c) //通过close关闭一个channel
	}()

	/*
		//ok为true就打印channel中的内容，如果为false就跳出当前循环
		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	//可以使用range来迭代不断操作的channel(上面那段可以用这个代替)
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main Finished..")

}

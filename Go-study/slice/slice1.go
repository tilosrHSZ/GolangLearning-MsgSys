package main

import "fmt"

func printArray(myArray []int) {
	//slice动态数组的时候是个引用传递
	//此处使用_来表示匿名变量，这里用于不需要index
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}

	myArray[0] = 100
}

func main1() {
	myArray := []int{1, 2, 3, 4} //这是一个动态数组

	fmt.Printf("myArray's type is %T\n", myArray)

	printArray(myArray)

	fmt.Println("-----")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}

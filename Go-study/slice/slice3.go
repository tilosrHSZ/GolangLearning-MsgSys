package main

import "fmt"

func main3() {
	var numbers = make([]int, 3, 5) //第二个数值是容量，3，4位已开辟但不可访问

	fmt.Printf("len = %d. cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1) //用append为numbers扩展

	fmt.Printf("len = %d. cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2) //用append为numbers扩展

	fmt.Printf("len = %d. cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//此时已经满了，再追加会再开辟容量，2倍扩容
	numbers = append(numbers, 3) //用append为numbers扩展

	fmt.Printf("len = %d. cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

}

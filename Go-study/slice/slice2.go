package main

import "fmt"

func main2() {

	//1
	slice1 := []int{1, 2, 3}

	//2
	var slice2 []int
	slice2 = make([]int, 3)

	//3
	var slice3 []int = make([]int, 3)
	slice3[0] = 100

	//4
	slice4 := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	if slice1 == nil {
		fmt.Println("slice1 is empty")
	} else {
		fmt.Println("slice1 is not empty")
	}

}

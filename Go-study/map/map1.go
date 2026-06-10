package main

import "fmt"

func main1() {
	//第一种声明方式
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is empty")
	}

	//map开辟地址空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "java"
	myMap1["two"] = "cpp"
	myMap1["three"] = "python"
	//输出的时候是乱序是因为map是一个哈希，不是按插入顺序排序的
	fmt.Println(myMap1)

	//第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "cpp"
	myMap2[3] = "python"

	fmt.Println(myMap2)

	//第三种声明方式
	myMap3 := map[string]string{
		"one":   "php",
		"two":   "cpp",
		"three": "py",
	}
	fmt.Println(myMap3)

}

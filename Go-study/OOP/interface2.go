package main

import "fmt"

func myfunc(arg interface{}) { //万能类型
	fmt.Println("-----")
	fmt.Println("myfunc called")
	fmt.Println(arg)

	//interface{}"断言"
	/*
		value, ok := arg.(string)
		if !ok {
			fmt.Println("arg is not a string")
			fmt.Printf("value type is %T\n", value)
		} else {
			fmt.Println("arg is a string, value = ", value)
		}
	*/
	//这样断言出来会有个问题，如果是string，value就是arg的值，反之会是string类型的空值

	//用type+swich...case断言
	switch v := arg.(type) {
	case string:
		fmt.Println("arg is string, value = ", v)

	case int:
		fmt.Println("arg is int, value = ", v)

	case float64:
		fmt.Println("arg is float64, value = ", v)

	case Book1:
		fmt.Println("arg is Book1, value = ", v)

	}
}

type Book1 struct {
	auth string
}

func main() {
	book := Book1{auth: "go"}

	myfunc(book)
	myfunc(10)
	myfunc("abc")
	myfunc(10.5)

}

package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("Value: ", reflect.ValueOf(arg))
}

func main4() {
	var num float64 = 3.14

	reflectNum(num)
}

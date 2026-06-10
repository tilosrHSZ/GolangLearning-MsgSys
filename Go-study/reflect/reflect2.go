package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is call")
	fmt.Printf("%v\n", this)
}

func main5() {
	user := User{1, "tilo", 20}

	DoFileAndMethod(user)
}

func DoFileAndMethod(input interface{}) {

	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过type
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		//方法也有类型，它的类型是func(inputtype)outputtype
		//所以这里会输出一个func(main.User)，main.User就是在main里面定义的User类型
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

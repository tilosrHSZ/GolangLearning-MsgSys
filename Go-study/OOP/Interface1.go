package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

//interface本质上是一个指针，后续传入具体的接口的时候需要注意以指针形式传入

//具体的类1
type Cat struct {
	color string
}

//自动连接到interface
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类2
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

//对interface父类操作的一个函数
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("type = ", animal.GetType())
	fmt.Println("color = ", animal.GetColor())
}

func main3() {
	//不同的类
	cat := Cat{"Black"}
	dog := Dog{"Yellow"}

	//相同的函数，传入不同的对象产生不一样的结果
	//实现多态
	showAnimal(&cat) //注意传入的是指针，如上文所说的接口本质上是一个指针类型
	showAnimal(&dog)

}

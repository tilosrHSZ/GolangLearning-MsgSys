package main

import (
	"fmt"
)

type Level struct {
	Number int
}

func (self *Level) LevelUP() {
	self.Number += 1
}

type Hero struct {
	Name  string
	Ad    int
	Level //继承了Level
}

/*
func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Number)
}

func (this Hero) GetName() string {
	return this.Name
}
//此处会是值传递而导致不会改变
func (this Hero) ChangeName(newName string) {
	this.Name = newName
}
*/

// 正确的写法
func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Number)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) ChangeName(newName string) {
	this.Name = newName
}

func main2() {
	//create
	hero := Hero{Name: "z3", Ad: 100, Level: Level{Number: 1}}

	hero.Show()

	hero.ChangeName("l4")

	hero.LevelUP()
	hero.Show()
}

//go语言没有class关键字，是通过struct再绑定方法来实现OOP编程

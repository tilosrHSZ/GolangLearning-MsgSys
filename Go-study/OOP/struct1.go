package main

import "fmt"

//type myint int

type Book struct {
	title string
	auth  string
}

//值传递不会改变外面的值
func changebook1(book Book) {
	book.auth = "1231"
}

//手动传个指针进来把它变成引用传递
func changebook2(book *Book) {
	book.auth = "3213" //这里也可以用(*book).auth = "3213"，但由于go的逆天语法糖可以自动解引用
}

func main1() {

	/*var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)
	*/

	var book1 Book
	book1.title = "fucv"
	book1.auth = "me"

	fmt.Printf("%v\n", book1)

	fmt.Println("-----")
	changebook1(book1)
	fmt.Printf("%v\n", book1)

	fmt.Println("-----")
	changebook2(&book1)
	fmt.Printf("%v\n", book1)

}

package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("write a book")
}

func main3() {
	//b:pair<type:Book, value:Book{}地址
	b := &Book{}

	//r:pair<type: , value: (用空格实际上就是nil)
	var r Reader
	//r:pair<type:Book, value:Book{}地址
	r = b
	r.ReadBook()

	//w:pair<type:Book, value:Book{}地址
	var w Writer
	w = r.(Writer)

	w.WriteBook()
}

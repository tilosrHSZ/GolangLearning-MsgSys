package main

import (
	"fmt"
	"io"
	"os"
)

func main2() {
	//tty:pair<type:*os.File, value:"/dev/tty"文件描述符
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open the file error", err)
		return
	}

	//r:pair<type: , value:
	var r io.Reader
	//r:pair<type:*os.File, value:"/dev/tty"文件描述符
	r = tty
	//pair的传递

	//w:pair<type: ,value: >
	var w io.Writer
	//w: pair<type:*os.File, value:"/dev/tty"文件描述符
	w = r.(io.Writer)

	w.Write([]byte("Test\n"))
}

package main

import (
	_ "golang-studying/init/lib1"      //匿名导入的方式：下划线+空格，后面即使不用这个包的方法也不会报错
	mylib2 "golang-studying/init/lib2" //重命名导入
	//. "golang-studying/init/lib2" //静态导入，之后使用方法不用写包名
)

func main() {
	//lib1.Lib1Test()
	mylib2.Lib2Test()
}

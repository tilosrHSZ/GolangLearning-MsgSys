package lib1

import "fmt"

//lib1 api,一定要大写才能对外开放
func Lib1Test() {
	fmt.Println("Lib1Test()...")
}

func init() {
	fmt.Println("lib1, init() ...")
}

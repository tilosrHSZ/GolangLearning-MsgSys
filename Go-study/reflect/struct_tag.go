package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"My name"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem() //Elem用于解引用并取出所有值

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo)
		fmt.Println("doc: ", tagdoc)
	}
}

func main6() {
	var re resume

	findTag(&re)
}

package main

import "fmt"

func PrintMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeMap(cityMap map[string]string, key string, value string) {
	cityMap[key] = value
}

func main() {
	cityMap := make(map[string]string)

	//insert
	cityMap["China"] = "Beijing"
	cityMap["JP"] = "KD"
	cityMap["USA"] = "LD"

	//show
	PrintMap(cityMap)

	//delete
	delete(cityMap, "USA")

	//change
	cityMap["JP"] = "TKY"

	fmt.Println("------")
	//show
	PrintMap(cityMap)
}

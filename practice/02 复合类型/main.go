package main

import "fmt"

func main() {

	/*
		slice 引用数据类型，改变变量副本值的时候，会改变变量本身的值
	*/

	var slice1 []int
	fmt.Printf("%v - %T - %v/n", slice1, slice1, len(slice1))

	var makeSlice = make([]int, 4, 8)
	fmt.Printf("%d -- %d/n", len(makeSlice), cap(makeSlice))

	/*
		map
	*/

	var userInfo = make(map[string]string)

	userInfo["name"] = "lele"
	userInfo["age"] = "20"

	var userInfo2 = map[string]string{
		"name": "sanshi",
		"age":  "25",
	}

	fmt.Println(userInfo)
	fmt.Println(userInfo2["name"], userInfo2["age"])

}

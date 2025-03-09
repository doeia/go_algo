package main

import "../01primeryType/calc"

import (
	"encoding/json"
	"fmt"
)

func getUserInfo() (string, int) {
	return "lele", 10
}

func main() {
	/*
		数据类型
	*/

	//基本： 整型，浮点，布尔，字符
	//复合： 数组，切片，结构体，map, 通道， 接口

	/*
		变量
	*/

	// var a int = 10
	// const myName string = "lele"
	// 类型推导
	b := 20

	// fmt.Println(a)
	fmt.Printf("%v, %T\n", b, b)
	// fmt.Println(myName)
	// fmt.Println("hello")

	/*
		变量2
	*/

	var userName, age = getUserInfo()
	//匿名变量 _
	var userName2, _ = getUserInfo()

	fmt.Println(userName, age)
	fmt.Println(userName2)

	/*
		array
	*/

	// 推导数组长度 [...]
	// var arr1 = [...]int{1, 2, 3, 4}
	// fmt.Println(len(arr1))

	// 指定下标的值，没指定的0
	var arr2 = [...]int{0: 1, 1: 2, 2: 3, 6: 4}
	fmt.Println(arr2)

	/*
		struct
	*/

	// 自定义类型
	type myInt int
	type myFn func(int, int) int

	type Address struct {
		Name  string
		Phone string
		City  string
	}

	type Person struct {
		ID      int `json:"id"` // label
		name    string
		Age     int // 大写表示共有
		gender  string
		Hobby   []string
		map1    map[string]string
		Address Address
	}

	var p1 Person // new(Person), &Person{}
	p1.name = "lele"
	p1.gender = "female"
	p1.Age = 20

	p1.Hobby = make([]string, 3, 6)
	p1.Hobby[0] = "coding"
	p1.Hobby[1] = "writing"

	p1.map1 = make(map[string]string)
	p1.map1["longtitue"] = "124.1131"
	p1.map1["latitude"] = "23.42134"

	p1.Address.Name = "chaoyang"
	p1.Address.Phone = "111"

	// jsonfy
	jsonByte, _ := json.Marshal(p1)
	jsonStr := string(jsonByte)

	fmt.Println("jsonStr", jsonStr)

	// 直接赋值
	var p2 = Person{
		name:   "sanshi",
		gender: "male",
		Age:    25,
	}

	fmt.Printf("%v %T", p1, p1)
	fmt.Printf("%v %T", p2, p2)

	sum := calc.Add(10, 2)

	fmt.Println(sum)
}

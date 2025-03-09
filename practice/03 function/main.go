package main

import "fmt"

/*
	type 自定义方法
*/
func add(x, y int) int {
	return x + y
}
func subsract(x, y int) int {
	return x - y
}

type calcType func(int, int) int

func calc(x, y int, ct calcType) int {
	return ct(x, y)
}

/*
	panic recover
*/

func funPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}() // 匿名自执行
	panic("a panic")
}

func main() {
	sum := calc(10, 5, add)
	fmt.Println(sum)
	sub := calc(10, 5, subsract)
	fmt.Println(sub)

	/*
		type 自定义方法
	*/
	fmt.Println("start")
	defer fmt.Println(0)
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println("end")

	funPanic()

	/*
		条件 循环
	*/
	if sum > 10 {
		fmt.Println(sum, " is lager than 10")
	}

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i, " is now")
	// }

	// var str = "hello go lang"

	// for _, v := range str {
	// 	fmt.Printf("%c\n", v)
	// }
}

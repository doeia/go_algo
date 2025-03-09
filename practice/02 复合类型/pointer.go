package pointer

import "fmt"

func pointer() {
	var a int = 10
	var p = &a // type *int
	fmt.Printf("%v %T %p", a, a, &a)
	fmt.Printf("%v %T", p, p)
	fmt.Println(*p) // print value of a
}

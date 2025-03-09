package main

import "github.com/shospring/decimal"

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	unixTime := timeObj.Unix()

	fmt.Println(timeObj)
	fmt.Println(unixTime)

	/*
		package:
		method 1: go mod download // global
		method 2: go mod vendor // copy to local project
	*/

	quantity := decimal.NewFromInt(3)
	fmt.Println(quantity)
}

package main

import (
	"fmt"
	"workspace_go/service"
)

func main() {
	user := User{
		Name: "John",
		Age:  25,
	}

	marshal, err := user.Marshal(user)
	if err != nil {
		panic(err)
	}

	newUser := &service.User{}
	err = newUser.Unmarshal(marshal, newUser)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.String())

}

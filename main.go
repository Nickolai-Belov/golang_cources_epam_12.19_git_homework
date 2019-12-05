package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("please enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello brave new ", name)
}

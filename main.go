package main

import "fmt"

func main() {
	const Brave, New string = " brave ", "new "
	var name string
	fmt.Print("please enter your name: ")
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Println("Hello", Brave, New, name)
	fmt.Println("How is the weather today?")
}

package main

import "fmt"

func main() {
	var name string
	fmt.Print("please enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
	fmt.Println("How is the weather today?")
}

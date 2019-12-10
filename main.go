package main

import "fmt"

func main() {
	const br = "brave"
	const ne = "new"
	fmt.Println("Hello ", br, " ", ne, " world")
	var hell = "Hello "
	var name string
	fmt.Println("please enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println("How is the weather today?")
}

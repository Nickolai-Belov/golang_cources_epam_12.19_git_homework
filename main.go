package main

import "fmt"

func main() {
	const b, n = "brave", "new"
	name := ""
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Printf("Hello %s %s %s\n", b, n, name)
	fmt.Println("How is the weather today?")
}

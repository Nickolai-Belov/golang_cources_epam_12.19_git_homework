package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Print("Please enter your name: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("Hello", input)
	fmt.Println("How is the weather today?")
}

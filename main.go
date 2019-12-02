package main

import (
	"fmt"
	"os"
)

const (
	BraveAdj = "brave"
	NewAdj = "new"
)

func main() {
	var input string
	fmt.Print("Please enter your name: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("Hello", BraveAdj, NewAdj, input)
	fmt.Println("How is the weather today?")

}

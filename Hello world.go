package main

import (
	"fmt"
	"os"
)

func main()  {
	var name string
	fmt.Printf("Please enter your name: ")
	_, _ = fmt.Fscan(os.Stdin, &name)
	fmt.Println("Hello, ", name)
	fmt.Printf("How is the weather today?")
	//weather asking
}
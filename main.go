package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("please enter your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("%s", text)
	fmt.Printf("Hello brave new%s", text)
	fmt.Println("How is the weather today?")
}

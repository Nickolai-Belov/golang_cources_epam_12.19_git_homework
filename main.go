package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter yout name: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Hello, brave new", text)
	fmt.Println("How is the weather today?")
}


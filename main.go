package main

  import (
	"bufio"
	"fmt"
	"os"
)

func main() {
 	fmt.Println("hello world")		 	
	reader := bufio.NewReader(os.Stdin)
     	fmt.Print("please enter your name: ")
     	text, _ := reader.ReadString('\n')
     	fmt.Println("hello brave new " + text)

      	fmt.Print("How is the weather today?")
}		 

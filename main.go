package golanghome

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("please enter your name:")
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	name := inputScanner.Text()

	fmt.Println("Hello", name)
	fmt.Println("How is the weather today?")
	inputScanner.Scan()
	weather := inputScanner.Text()
	fmt.Println("Weather today:", weather)
	const Brave = "brave"
	const New = "new"

	fmt.Println("Hello", Brave, New, "world")
}

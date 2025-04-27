package switchPractice

import (
	"bufio"
	"fmt"
	"os"
)

func SwitchTest() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your favorite Adventure Time character: ")
	scanner.Scan()
	var answer = scanner.Text()
	switch {
	case answer == "ice king":
		fmt.Println("simple simon")
	case answer == "finn":
		fmt.Println("he has a hero heart")
	case answer == "jake":
		fmt.Println("sucking at making CLI tools is the first step ")
		fmt.Println("towards being sort of good at making CLI tools")
	case answer == "lsp":
		fmt.Println("oh my glob")
	default:
		fmt.Println("thats a good choice")
	}
}

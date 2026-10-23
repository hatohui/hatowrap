package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your full name: ")

	if scanner.Scan() {
		input := scanner.Text()
		cleanInput := strings.TrimSpace(input)
		fmt.Printf("Hello, %s!\n", cleanInput)
	}
}

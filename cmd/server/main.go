package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// could use fmt.Scan instead, but it divides words
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hello, db! \".exit\" to quit")
	for {
		fmt.Print("db> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		if input == ".exit" {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}

		fmt.Printf("Command: '%s'\n", input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Unexpected error:", err)
		os.Exit(1)
	}
}
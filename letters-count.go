package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var (
		n int
		s string
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	for _, letter := range s {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			n++
		}
	}
	fmt.Print(n)
}

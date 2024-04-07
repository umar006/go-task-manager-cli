package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(`Menu:
	1. Add task
	2. List task
	3. Complete task
	4. Delete task
	`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	choice := scanner.Text()

	fmt.Println(choice)
}

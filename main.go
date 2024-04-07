package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var tasks []string

func main() {
	fmt.Println(`Menu:
	1. Add task
	2. List task
	3. Complete task
	4. Delete task
	`)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	menu := scanner.Text()

	switch menu {
	case "1":
		for {
			fmt.Print("what's your task? ")
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			newTask := scanner.Text()
			if len(newTask) == 0 {
				break
			}

			tasks = append(tasks, newTask)
		}
	default:
		fmt.Println("wrong menu")
	}
}

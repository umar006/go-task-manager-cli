package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var tasks []string

func main() {
	for {
		fmt.Println("Menu")
		fmt.Println("1. Add task")
		fmt.Println("2. List task")
		fmt.Println("3. Complete task")
		fmt.Println("4. Delete task")
		fmt.Println("0. Exit")
		fmt.Println()

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		menu := scanner.Text()

		switch menu {
		case "1":
			for {
				fmt.Print("What's your task? ")
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
			fmt.Println()
		case "2":
			fmt.Println("Task list:")
			for _, task := range tasks {
				fmt.Println("\t- " + task)
			}
			fmt.Println()
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Wrong menu!")
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	task      string
	completed bool
}

var (
	tasks          []Task
	completedTasks []Task
)

func main() {
	for {
		MainMenu()
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

				inputTask := scanner.Text()
				if len(inputTask) == 0 {
					break
				}

				newTask := Task{task: inputTask}

				tasks = append(tasks, newTask)
			}
			fmt.Println()
		case "2":
			TaskList()
			fmt.Println()
		case "3":
			if len(tasks) == 0 {
				fmt.Println("No in progress tasks!")
				fmt.Println()
				break
			}

			for {
				TaskList()
				fmt.Println()

				fmt.Print("Complete a task? ")
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				completeTask := scanner.Text()
				if len(completeTask) == 0 {
					break
				}

				var afterComplete []Task
				for idx, task := range tasks {
					toInt, err := strconv.Atoi(completeTask)
					if err != nil {
						log.Fatal(err)
					}

					if idx+1 != toInt {
						afterComplete = append(afterComplete, task)
					} else {
						completedTasks = append(completedTasks, task)
					}
				}
				tasks = afterComplete
				fmt.Println()
			}
			fmt.Println()
		case "4":
			if len(tasks) == 0 {
				fmt.Println("Do not have tasks!")
				fmt.Println()
				break
			}

			for {
				TaskList()
				fmt.Println()

				fmt.Print("Delete a task? ")
				scanner.Scan()
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}

				deleteTask := scanner.Text()
				if len(deleteTask) == 0 {
					break
				}

				var afterRemove []Task
				for idx, task := range tasks {
					toInt, err := strconv.Atoi(deleteTask)
					if err != nil {
						log.Fatal(err)
					}

					if idx+1 != toInt {
						afterRemove = append(afterRemove, task)
					}
				}

				tasks = afterRemove
				fmt.Println()
			}
			fmt.Println()
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Wrong menu!")
		}
	}
}

func MainMenu() {
	fmt.Println("Menu")
	fmt.Println("1. Add task")
	fmt.Println("2. List task")
	fmt.Println("3. Complete task")
	fmt.Println("4. Delete task")
	fmt.Println("0. Exit")
}

func TaskList() {
	fmt.Println("Task list:")
	for idx, task := range tasks {
		fmt.Printf("\t%d. %s\n", idx+1, task.task)
	}

	fmt.Println("Completed task list:")
	for idx, task := range completedTasks {
		fmt.Printf("\t%d. %s\n", idx+1, task.task)
	}
}

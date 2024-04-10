package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Task struct {
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var tasks []Task

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ReadTaskFromJsonFile()

	for {
		MainMenu()
		fmt.Println()

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		menu := scanner.Text()

		switch menu {
		case "1":
			AddTask(scanner)
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

			CompleteTask(scanner)
			fmt.Println()
		case "4":
			if len(tasks) == 0 {
				fmt.Println("Do not have tasks!")
				fmt.Println()
				break
			}

			DeleteTask(scanner)
			fmt.Println()
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Wrong menu!")
		}
	}
}

func ReadTaskFromJsonFile() {
	tasksRead, err := os.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(tasksRead, &tasks)
	if err != nil {
		log.Fatal(err)
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
		if !task.Completed {
			fmt.Printf("\t%d. %s\n", idx+1, task.Task)
		} else {
			fmt.Printf("\t%d. %s\t\t\u2713 Complete\n", idx+1, task.Task)
		}
	}
}

func AddTask(scanner *bufio.Scanner) {
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

		newTask := Task{Task: inputTask}

		tasks = append(tasks, newTask)
	}

	tasksByte, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("tasks.json", tasksByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func CompleteTask(scanner *bufio.Scanner) {
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

			if idx+1 == toInt {
				task.Completed = true
			}
			afterComplete = append(afterComplete, task)
		}
		tasks = afterComplete
		fmt.Println()
	}
}

func DeleteTask(scanner *bufio.Scanner) {
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
}

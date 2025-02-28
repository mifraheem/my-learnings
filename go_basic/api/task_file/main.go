package main

import (
	"fmt"
	"os"
)

var filename string = "./tasks.txt"

type Task struct {
	Name string
	Done bool
	Due  string
}

func main() {
	fmt.Println("Welcome to task tracker system")
	fmt.Println("X-----------------------------X")
	create_file()
	all_tasks := load_tasks()
	fmt.Println(all_tasks)
	addTaskToFile("ifraheem", "2024-02-20")

}

func create_file() {
	_, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(filename)
			fmt.Println("File created")
		}
	}
}

func load_tasks() string {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(file)
}

func addTaskToFile(name, due string) {

	os.WriteFile(filename, []byte(name+" "+due+"\n"), 0644)
}

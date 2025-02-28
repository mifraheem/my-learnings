package main

import "fmt"

type Task struct {
	title   string
	done    bool
	dueDate string
}

func main() {
	fmt.Println("Welcome to Task Tracking System")
	fmt.Println("--------------------------------")
	addTask("ifraheem has to go", "23 june")
	addTask("ifraheem o go", "2 june")
	addTask("ifraheem o go", "25 june")

	displayTasks()
	completeTask(1)
	displayTasks()
}

var tasks []Task

func addTask(title, due string) {
	newTask := Task{
		title:   title,
		dueDate: due,
		done:    false,
	}
	tasks = append(tasks, newTask)
	fmt.Println("Task added successfully")
}

func deleteTask(index int) {
	if index > 0 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task has been deleted Successfully")
	}
}

func displayTasks() {
	for i, task := range tasks {
		status := "Pending"
		if task.done {
			status = "Done"
		}
		fmt.Printf("Task %d: %s - %s - %s\n", i+1, task.title, task.dueDate, status)
		// fmt.Println(i, task)
	}

}

func completeTask(index int) {
	if index > 0 && index <= len(tasks) {

		// target_task := tasks[index]
		// target_task.done = true
		tasks[index].done = true
		fmt.Println("Your Task has been completed")
	}
}

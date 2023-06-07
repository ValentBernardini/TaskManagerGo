package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name string
	desc string
	done bool
}

type TaskList struct {
	tasks []Task
}

// Add Task
func (t *TaskList) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}

// Done
func (t *TaskList) DoneTask(index int) {
	t.tasks[index].done = true
}

// Edit Task
func (t *TaskList) EditTask(index int, task Task) {
	t.tasks[index] = task
}

// Remove Task
func (t *TaskList) RemoveTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	list := TaskList{}

	//Instacy of bufio for read the input
	read := bufio.NewReader(os.Stdin)

	for {
		var option int
		fmt.Println("select an option: \n",
			"1. Add Task\n",
			"2. Done Task\n",
			"3. Edit Task\n",
			"4. Remove Task\n",
			"5. Exit",
		)
		fmt.Print("Insert option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var task Task
			fmt.Print("Task Name: ")
			task.name, _ = read.ReadString('\n') //Read the input
			fmt.Print("Task Description: ")
			task.desc, _ = read.ReadString('\n')
			list.AddTask(task)
			fmt.Println("Task Added")
		case 2:
			var index int
			fmt.Print("Task Index: ")
			fmt.Scanln(&index)
			list.DoneTask(index)
			fmt.Println("Task Done")
		case 3:
			var index int
			var task Task
			fmt.Print("Task Index for editing the task: ")
			fmt.Scanln(&index)
			fmt.Print("Task Name: ")
			task.name, _ = read.ReadString('\n') //Read the input
			fmt.Print("Task Description: ")
			task.desc, _ = read.ReadString('\n')
			list.EditTask(index, task)
			fmt.Println("Task Edited")
		case 4:
			var index int
			fmt.Print("Task Index for remove the task: ")
			fmt.Scanln(&index)
			list.RemoveTask(index)
			fmt.Println("Task Removed")
		case 5:
			fmt.Println("Bye....")
			return
		default:
			fmt.Println("Invalid Option")
		}
		//Print the tasks
		fmt.Println("Task List: ")
		fmt.Println("====================")
		for i, task := range list.tasks {
			fmt.Printf("%d. %s - %s - Done: %t\n", i, task.name, task.desc, task.done)
		}
		fmt.Println("====================")

	}

}

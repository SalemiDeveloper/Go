package main 

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"taskcli/task"
)

const taskFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: taskcli [add|list|done|delete|edit]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Uso: taskcli add \"descrição da tarefa\"")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		_ = task.Add(taskFile, title)

	case "list":
		_ = task.List(taskFile)

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Uso: taskcli done [id]")
			return
		}
		id := mustAtoi(os.Args[2])
		_ = task.Done(taskFile, id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Uso: taskcli delete [id]")
			return
		}

		id := mustAtoi(os.Args[2])
		_ = task.Delete(taskFile, id)

	case "edit":
		if len(os.Args) < 4 {
			fmt.Println("Uso: taskcli edit [id] \"novo título\"")
			return
		}
		id := mustAtoi(os.Args[2])
		newTitle := strings.Join(os.Args[3:], " ")
		_ = task.Edit(taskFile, id, newTitle)

	default:
		fmt.Println("Comando desconhecido:", os.Args[1])
	}
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ID inválido:", s)
		os.Exit(1)
	}
	return n
}
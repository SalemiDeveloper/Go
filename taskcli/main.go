package main 

// importação de bibliotecas
import (
	"fmt"				// impressão de mensagens
	"os"				// interação com o sistema operacional
	"strconv"			// conversão de strings para números
	"strings"			// manipulação de strings
	"taskcli/task"		// importação do pacote task, que contém as funções para gerenciar as tarefas
)

const taskFile = "tasks.json" // arquivo JSON que armazena as tarefas

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

/* 
Função para converter uma string para um número inteiro, retornando um erro se a conversão falhar.
Essa função é usada para converter o ID da tarefa fornecido como um argumento de linha de comando.

. Tenta converter a string para um número inteiro
. Se a conversão falhar, imprime uma mensagem de erro e sai do programa.
. Retorna o número inteiro convertido.
*/
func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ID inválido:", s)
		os.Exit(1)
	}
	return n
}

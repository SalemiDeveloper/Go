package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"created_at"`
	DoneAt    *time.Time `json:"done_at,omitempty"`
}

/* Função para carregar uma lista de tarefas de um arquivo JSON.

. Abre o arquivo especificado pelo parâmetro filename
	. Se o arquivo não existir, retorna uma lista vazia e um erro nil.
	. Se ocorreu um arro ao abrir o arquivo, retorna um erro.
. Cria uma variável para armazenar a lista de tarefas.
. Deserializa o conteúdo do arquivo JSON em uma lista de tarefas.
	. Se ocorrer um erro ao deserializar o JSON, retorna um erro.
. Retorna a lista de tarefas carregadas com sucesso.



file, err := os.ReadFile(filename) -> Esta linha linha abre o arquivo especificado pelo parâmetro 'filename' 
                                   e armazena o conteúdo do arquivo em uma variável 'file'. O parâmetro 'err' 
								   é usado para armazenar qualquer erro que ocorra durante a abertura 
								   do arquivo.

if err := nil {...}                -> Esta seção verifica se ocorreu um erro ao abrir o arquivo. Se o erro for 'nil', 
          			               significa que o arquivo foi aberto com sucesso.

if os.IsNotExist(err) {...}        -> Esta seção verifica se o erro é um erro de arquivo não encontrado. Se for, retorna
							       uma lista vazia e um erro 'nil', indicando que o arquivo não existe.

return nil, err                    -> Se ocorreu um erro ao abrir o arquivo e não é um erro de arquivo não encontrado, retorna um erro.

var tasks []Task                   -> Cria uma variável para armazenar a lista de tarefas.

err = json.Unmarshal(file, &tasks) -> Deserializa o conteúdo do arquivo JSON em uma lista de tarefas. O parâmetro
                                      '&tasks' é uma referência à variável 'tasks' que foi criada anteriomente.

2° if err != nil {...}             -> Verifica se ocorreu um erro ao deserializar o JSON. Se for, retorna um erro.

return tasks, nil                  -> Retorna a lista de tarefas carregadas com sucesso.
*/

func LoadTasks(filename string) ([]Task, error) {	
	file, err := os.ReadFile(filename)
	if err != nil {		
		if os.IsNotExist(err) {
			return []Task{}, nil
		}		
		return nil, err
	}
	
	var tasks []Task	
	err = json.Unmarshal(file, &tasks)
	if err != nil {		
		return nil, err
	}
	
	return tasks, nil
}

/* Função para salvar uma lista de tarefas em um arquivo JSON.

. Serializa a lista de tarefas em um JSON.
	. Se ocorrer um erro ao serializar o JSON, retorna um erro.
. Salva o JSON no arquivo especificado pelo parâmetro filename.



data, err := json.MarshalIndent(tasks, "", "  ") -> Esta linha serializa a lista de tarefas em um JSON. O parâmetro 'tasks' é a lista de tarefas
                                                 a ser serializada. O parâmetro 'json.MarshalIndent' é usado para formatar o JSON com indentação.
												 parâmetro 'err' é usado para armazenar qualquer erro que ocorrer durante a serialização.

ir err := nil {...}	                             -> Verifica se ocorreu um erro ao serializar o JSON. Se sim, retorna um erro.

return os.WriteFile(filename, data, 0644)        -> Salva o JSON no arquivo especificado pelo parâmetro 'filename'. O parâmetro 'data' é o JSON 
                                                 serializado anteriormente. O parâmetro '0644' é o modo de permissão do arquivo, que significa 
												 que o arquivo será criado com permissão de leitura e escrita para o proprietário e permissão de 
												 leitura para os outros.

OBSERVAÇÃO: 
 - A função 'json.MarshalIndent' é usada para formatar o JSON com indentação, o que torna mais fácil de ler e entender o JSON gerado.
 - A função 'is.WriteFile' é usada para salvar o JSON no arquivo. Ela substitui o conteúdo do arquivo existente, se ele existir.
 - O modo de permissão '0644' é um valor comum para arquivos ded texto, que permite que o proprietário leia e escreva no arquivo, e os outros apenas leiam.
*/

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

/* Função para adicionar uma nova tarefa à lista de tarefas

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Gera um novo ID para a tarefa.
	. Se a lista de tarefas não estiver vazia, gera um novo ID baseado no último ID existente.
. Adiciona a nova tarefa à lista de tarefas.
. Salva a lista de tarefas atualizada.
. Imprime uma mensagem de confirmação.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.

if err != nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

id := 1                           -> Gera um novo ID para a tarefa.

if len(tasks) > 0 {...}           -> Se a lista de tarefas não estiver vazia, gera um novo ID baseado no último ID existente.

t := Tasks{...}                   -> Cria uma nova tarefa com o título e o novo ID.

tasks = append(tasks, t)          -> Adiciona a nova tarefa à lista de tarefas.

if err := SaveTasks(filename, tasks); err != nil {...} -> Salva a lista de tarefas atualizada e verifica se ocorreu um erro. Se sim, retorna um erro.

fmt.Println("Tarefa adicionada:", title)               -> Imprime uma mensagem de confirmação.

return nil                        -> Retorna sem erro.

OBSERVAÇÕES:
 - A função 'LoadTasks' é usada para carregar a lista de tarefas existente.
 - A função 'SaveTasks' é usada para salvar a lista de tarefas atualizada.
 - O novo ID é gerado baseado no último ID existente, se a lista de tarefas não estiver vazia.
*/

func Add(filename, title string) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks) - 1].ID + 1
	}

	t := Task{
		ID:        id,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, t)

	if err := SaveTasks(filename, tasks); err != nil {
		return err
	}
	fmt.Println("Tarefa adicionada:\"", title, "\"")
	return nil
}

/* Função para listar todas as tarefas

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Verifica se a lista de tarefas está vazia.
	. Se a lista de tarefas estiver vazia, imprime uma mensagem informando que não há tarefas.
. Imprime o cabeçalho de lista de tarefas.
. Itera sobre a lista de tarefas e imprime cada tarefa.
	. Formata a data de criação da tarefa.
	. Imprime a tarefa.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.

if err := nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

if len(tasks) == 0 {...}          -> Verifica se a lista de tarefas está vazia. Se sim, imprime uma mensagem informando que não há tarefas e retorna sem erro.

for _, t := range tasks {...}     -> Itera sobre a lista de tarefas e imprime cada tarefa.

fmt.Printf(...)                   -> Imprime a tarefa.

return nil                        -> Retorna sem erro.
*/

func List(filename string) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("Nenhuma tarefa encontrada.")
		return nil
	}

	for _, t := range tasks {
		status := "X - Fazendo"
		if t.Done {
			status = "V - Feito"
		}

		fmt.Printf("[%d] %s %s (criada em %s)\n", t.ID, status, t.Title, t.CreatedAt.Format("02/01/2006"))
	}
	return nil
}

/* Função para marcar uma tarefa como concluída

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro
. Itera sobre a lista de tarefas e verifica se a tarefa com o ID especificado pelo parâmetro id existe.
	. Se a tarefa existir, verifica se ela já foi concluída.
	. Se a tarefa ainda não foi concluída, marca ela como concluída e atualiza a data de conclusão.
. Salva a lista de tarefas atualizada.
	. Se ocorrer um erro ao salvar a lista de tarefas, retorna um erro.
. Imprime uma mensagem de confirmação indicando que a tarefa foi concluída.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.

if err != nil {...} -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

for i := range tasks {...} -> Itera sobre a lista de tarefas para procurar a tarefa com o ID especificado.

if t.ID == id {...} -> Se a tarefa foi encontrada, a marca como concluída.

tasks[i].Done = true -> Marca a tarefa como concluída.

tasks[i].DoneAt = &now -> Atualiza a data de conclusão da tarefa.

if err := SaveTasks(filename, tasks); err != nil {...} -> Salva a lista de tarefas atualizada e verifica se ocorreu algum erro. Se sim, retorna erro.

OBSERVAÇÕES:
 - A função 'LoadTasks()' é usada para crregar a lista  de tarefas existente.
 - A função 'SaveTasks()' é usada para salvar a lista de tarefas atualizada.
 - A tarefa é marcada como concluída alterando o calor do campo 'Done' para true.
 - Se a tarefa não for encontrada, uma mensagem é impressa informando que a tarefa não existe.
*/

func Done(filename string, id int) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Done {
				fmt.Println("Essa tarefa já está concluída.")
				return nil
			}

			now := time.Now()
			tasks[i].Done = true
			tasks[i].DoneAt = &now
			if err := SaveTasks(filename, tasks); err != nil {
				return err
			}

			fmt.Println("Tarefa concluída:", tasks[i].Title)
			return nil
		}
	}

	fmt.Println("Tarefa não encontrada.")
	return nil
}

/* Função para excluir uma tarefa.

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Cria uma nova lista de tarefas a serem excluídas.
	. Se a tarefa foi encontrada, não a inclui na nova lista de tarefas.
	. Se a tarefa não foi encontrada, a inclui na nova lista de tarefas.
. Verifica se a tarefa foi encontrada.
	. Se a tarefa não foi encontrada, imprime uma mensagem informando.
. Salva a nova lista de tarefas.



tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.

if err != nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna erro.

newTasks := []Task{}              -> Cria uma nova lista de tarefas vazia.

for _, t := range tasks {...}     -> Procura a tarefa com o ID especificado.

if t.ID == id {...}               -> Se a tarefa foi encontrada, não a inclui na nova lista e imprime uma mensagem de confirmação.

newTasks = append(newTasks, t)    -> Se a tarefa não foi encontrada, a inclui na nova lista.

if !found {...}                   -> Verifica se a tarefa foi encontrada. Se não foi, imprime uma mensagem informando.

return SaveTasks(filename, newTasks) -> Salva a nova lista de tarefas.

OBSERVAÇÕES:
 - A função 'LoadTasks()' é usada para carregar a lista de tarefas existente.
 - A função 'SaveTasks()' é usada para salvar a nova lista de tarefas.
 - A tarefa é excluída da lista de tarefas criando uma nova lista sem a tarefa a ser excluída.
 - Se a tarefa não for encontrada, uma mensagem é impressa informando que a tarefa não existe.
*/
func Delete(filename string, id int) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID == id {
			found = true
			fmt.Println("Tarefa deletada:", t.Title)
			// não adiciona no newTasks -> significa que foi "deletada"
		} else {
			newTasks = append(newTasks, t)
		}
	}

	if !found {
		fmt.Println("Tarefa não encontrada.")
		return nil
	}

	return SaveTasks(filename, newTasks)
}


/* Função para editar uma tarefa existente.

. Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Itera sobre a lista de tarefas e verifica se a tarefa com o ID especificado pelo parâmetro 'id' existe.
	. Se a tarefa existir, atualiza o título da tarefa com o novo título especificado pelo parâmetro 'newTitle'.
. Salva a lista de tarefas atualizada no arquivo especificado pelo parâmetro 'filename'.
	. Se ocorrer um erro ao salvar a lista de tarefas, retorna um erro.
. Imprime uma mensagem de confirmação indicando que a tarefa foi editada.
. Retorna sem erro.



tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo parâmetro 'filename'.

if err := nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna erro.

for i := range tasks {...}        -> Itera sobre a lista de tarefas para procurar a tarefa com o ID especificado.

if tasks[i].ID == id {...}        -> Se a tarefa foi encontrada, atualiza o título da tarefa com o novo título.

old := tasks[i].Title             -> Armazena o título antigo da tarefa.

tasks[i].Title = newTitle         -> Atualiza o título da tarefa com o novo título.

if err := SaveTasks(filename, tasks); err := nil {...} -> Salva a lista de tarefas atualizada e verifica se ocorreu erro. Se sim, retorna erro.

fmt.Printf(...)                   -> Imprime uma mensagem de confirmação indicando que a tarefa foi editada.

return nil                        -> Retorna sem erro.
*/

func Edit(filename string, id int, newTitle string) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			old := tasks[i].Title
			tasks[i].Title = newTitle
			if err := SaveTasks(filename, tasks); err != nil {
				return err
			}
			fmt.Printf("Tarefa %d editada: \"%s\" -> \"%s\"\n", id, old, newTitle)
			return nil
		}
	}

	fmt.Println("Tarefa não encontrada.")
	return nil
}

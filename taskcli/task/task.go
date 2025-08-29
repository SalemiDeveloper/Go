package task

import (
	"enconding/json"
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

/* Fun��o para carregar uma lista de tarefas de um arquivo JSON.

. Abre o arquivo especificado pelo par�metro filename
	. Se o arquivo n�o existir, retorna uma lista vazia e um erro nil.
	. Se ocorreu um arro ao abrir o arquivo, retorna um erro.
. Cria uma vari�vel para armazenar a lista de tarefas.
. Deserializa o conte�do do arquivo JSON em uma lista de tarefas.
	. Se ocorrer um erro ao deserializar o JSON, retorna um erro.
. Retorna a lista de tarefas carregadas com sucesso.



file, err := os.ReadFile(filename) -> Esta linha linha abre o arquivo especificado pelo par�metro 'filename' 
                                   e armazena o conte�do do arquivo em uma vari�vel 'file'. O par�metro 'err' 
								   � usado para armazenar qualquer erro que ocorra durante a abertura 
								   do arquivo.

if err := nil {...}                -> Esta se��o verifica se ocorreu um erro ao abrir o arquivo. Se o erro for 'nil', 
          			               significa que o arquivo foi aberto com sucesso.

if os.IsNotExist(err) {...}        -> Esta se��o verifica se o erro � um erro de arquivo n�o encontrado. Se for, retorna
							       uma lista vazia e um erro 'nil', indicando que o arquivo n�o existe.

return nil, err                    -> Se ocorreu um erro ao abrir o arquivo e n�o � um erro de arquivo n�o encontrado, retorna um erro.

var tasks []Task                   -> Cria uma vari�vel para armazenar a lista de tarefas.

err = json.Unmarshal(file, &tasks) -> Deserializa o conte�do do arquivo JSON em uma lista de tarefas. O par�metro
                                      '&tasks' � uma refer�ncia � vari�vel 'tasks' que foi criada anteriomente.

2� if err != nil {...}             -> Verifica se ocorreu um erro ao deserializar o JSON. Se for, retorna um erro.

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

/* Fun��o para salvar uma lista de tarefas em um arquivo JSON.

. Serializa a lista de tarefas em um JSON.
	. Se ocorrer um erro ao serializar o JSON, retorna um erro.
. Salva o JSON no arquivo especificado pelo par�metro filename.



data, err := json.MarshalIndent(tasks, "", "  ") -> Esta linha serializa a lista de tarefas em um JSON. O par�metro 'tasks' � a lista de tarefas
                                                 a ser serializada. O par�metro 'json.MarshalIndent' � usado para formatar o JSON com indenta��o.
												 par�metro 'err' � usado para armazenar qualquer erro que ocorrer durante a serializa��o.

ir err := nil {...}	                             -> Verifica se ocorreu um erro ao serializar o JSON. Se sim, retorna um erro.

return os.WriteFile(filename, data, 0644)        -> Salva o JSON no arquivo especificado pelo par�metro 'filename'. O par�metro 'data' � o JSON 
                                                 serializado anteriormente. O par�metro '0644' � o modo de permiss�o do arquivo, que significa 
												 que o arquivo ser� criado com permiss�o de leitura e escrita para o propriet�rio e permiss�o de 
												 leitura para os outros.

OBSERVA��O: 
 - A fun��o 'json.MarshalIndent' � usada para formatar o JSON com indenta��o, o que torna mais f�cil de ler e entender o JSON gerado.
 - A fun��o 'is.WriteFile' � usada para salvar o JSON no arquivo. Ela substitui o conte�do do arquivo existente, se ele existir.
 - O modo de permiss�o '0644' � um valor comum para arquivos ded texto, que permite que o propriet�rio leia e escreva no arquivo, e os outros apenas leiam.
*/

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

/* Fun��o para adicionar uma nova tarefa � lista de tarefas

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Gera um novo ID para a tarefa.
	. Se a lista de tarefas n�o estiver vazia, gera um novo ID baseado no �ltimo ID existente.
. Adiciona a nova tarefa � lista de tarefas.
. Salva a lista de tarefas atualizada.
. Imprime uma mensagem de confirma��o.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo par�metro 'filename'.

if err != nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

id := 1                           -> Gera um novo ID para a tarefa.

if len(tasks) > 0 {...}           -> Se a lista de tarefas n�o estiver vazia, gera um novo ID baseado no �ltimo ID existente.

t := Tasks{...}                   -> Cria uma nova tarefa com o t�tulo e o novo ID.

tasks = append(tasks, t)          -> Adiciona a nova tarefa � lista de tarefas.

if err := SaveTasks(filename, tasks); err != nil {...} -> Salva a lista de tarefas atualizada e verifica se ocorreu um erro. Se sim, retorna um erro.

fmt.Println("Tarefa adicionada:", title)               -> Imprime uma mensagem de confirma��o.

return nil                        -> Retorna sem erro.

OBSERVA��ES:
 - A fun��o 'LoadTasks' � usada para carregar a lista de tarefas existente.
 - A fun��o 'SaveTasks' � usada para salvar a lista de tarefas atualizada.
 - O novo ID � gerado baseado no �ltimo ID existente, se a lista de tarefas n�o estiver vazia.
*/

func Add(filename, title string) error {
	tasks, err := LoadTasks(filename)
	if err := nil {
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
	fmt.Println("Tarefa adicionada:", title)
	return nil
}

/* Fun��o para listar todas as tarefas

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Verifica se a lista de tarefas est� vazia.
	. Se a lista de tarefas estiver vazia, imprime uma mensagem informando que n�o h� tarefas.
. Imprime o cabe�alho de lista de tarefas.
. Itera sobre a lista de tarefas e imprime cada tarefa.
	. Formata a data de cria��o da tarefa.
	. Imprime a tarefa.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo par�metro 'filename'.

if err := nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

if len(tasks) == 0 {...}          -> Verifica se a lista de tarefas est� vazia. Se sim, imprime uma mensagem informando que n�o h� tarefas e retorna sem erro.

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

/* Fun��o para marcar uma tarefa como conclu�da

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro
. Itera sobre a lista de tarefas e verifica se a tarefa com o ID especificado pelo par�metro id existe.
	. Se a tarefa existir, verifica se ela j� foi conclu�da.
	. Se a tarefa ainda n�o foi conclu�da, marca ela como conclu�da e atualiza a data de conclus�o.
. Salva a lista de tarefas atualizada.
	. Se ocorrer um erro ao salvar a lista de tarefas, retorna um erro.
. Imprime uma mensagem de confirma��o indicando que a tarefa foi conclu�da.
. Retorna sem erro.

tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo par�metro 'filename'.

if err != nil {...} -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna um erro.

for i := range tasks {...} -> Itera sobre a lista de tarefas para procurar a tarefa com o ID especificado.

if t.ID == id {...} -> Se a tarefa foi encontrada, a marca como conclu�da.

tasks[i].Done = true -> Marca a tarefa como conclu�da.

tasks[i].DoneAt = &now -> Atualiza a data de conclus�o da tarefa.

if err := SaveTasks(filename, tasks); err != nil {...} -> Salva a lista de tarefas atualizada e verifica se ocorreu algum erro. Se sim, retorna erro.

OBSERVA��ES:
 - A fun��o 'LoadTasks()' � usada para crregar a lista  de tarefas existente.
 - A fun��o 'SaveTasks()' � usada para salvar a lista de tarefas atualizada.
 - A tarefa � marcada como conclu�da alterando o calor do campo 'Done' para true.
 - Se a tarefa n�o for encontrada, uma mensagem � impressa informando que a tarefa n�o existe.
*/

func Done(filename string, id int) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return nil
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if tasks[i].Done {
				fmt.Println("Essa tarefa j� est� conclu�da.")
				return nil
			}

			now := time.Now()
			tasks[i].Done = true
			tasks[i].DoneAt = &now
			if err := SaveTasks(filename, tasks); err != nil {
				return err
			}

			fmt.Println("Tarefa conclu�da:", tasks[i].Title)
			return nil
		}
	}

	fmt.Println("Tarefa n�o encontrada.")
	return nil
}

/* Fun��o para excluir uma tarefa.

. Carrega a lista de tarefas existente.
	. Se ocorrer um erro ao carregar a lista de tarefas, retorna um erro.
. Cria uma nova lista de tarefas a serem exclu�das.
	. Se a tarefa foi encontrada, n�o a inclui na nova lista de tarefas.
	. Se a tarefa n�o foi encontrada, a inclui na nova lista de tarefas.
. Verifica se a tarefa foi encontrada.
	. Se a tarefa n�o foi encontrada, imprime uma mensagem informando.
. Salva a nova lista de tarefas.



tasks, err := LoadTasks(filename) -> Carrega a lista de tarefas existente do arquivo especificado pelo par�metro 'filename'.

if err != nil {...}               -> Verifica se ocorreu um erro ao carregar a lista de tarefas. Se sim, retorna erro.

newTasks := []Task{}              -> Cria uma nova lista de tarefas vazia.

for _, t := range tasks {...}     -> Procura a tarefa com o ID especificado.

if t.ID == id {...}               -> Se a tarefa foi encontrada, n�o a inclui na nova lista e imprime uma mensagem de confirma��o.

newTasks = append(newTasks, t)    -> Se a tarefa n�o foi encontrada, a inclui na nova lista.

if !found {...}                   -> Verifica se a tarefa foi encontrada. Se n�o foi, imprime uma mensagem informando.

return SaveTasks(filename, newTasks) -> Salva a nova lista de tarefas.

OBSERVA��ES:
 - A fun��o 'LoadTasks()' � usada para carregar a lista de tarefas existente.
 - A fun��o 'SaveTasks()' � usada para salvar a nova lista de tarefas.
 - A tarefa � exclu�da da lista de tarefas criando uma nova lista sem a tarefa a ser exclu�da.
 - Se a tarefa n�o for encontrada, uma mensagem � impressa informando que a tarefa n�o existe.
*/
func Delete(filename string, id int) error {
	tasks, err := LoadTasks(filename)
	if err != nil {
		return err
	}

	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID = id {
			newTasks = append(newTasks, t)
		} else {
			found = true
			fmt.Println("Tarefa deletada:", t.Title)
		}
	}

	if !found {
		fmt.Println("Tarefa n�o encontrada.")
		return nil
	}

	return SaveTasks(filename, newTasks)
}

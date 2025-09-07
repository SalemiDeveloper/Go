/*
EXERCÍCIO 01

Usando uma literal composta, crie um array qie suporte 5 valores do tipo int e atribua valores aos seus índices.
Utilize range e demonstre os valores do array. Utilizando format printing, demonstre o tipo do array.
*/

package main
import "fmt"

func main() {

	array := [5]int{10, 20, 30, 40, 50}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)
}

/*
EXERCÍCIO 02

Usando uma literal composta, crie uma slice do tipo int e atribua 10 valores a ela.
Utilize range para demonstrar todos estes valores e utilize format printing para demonstrar seu tipo.
*/

package main
import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)
}


/*
EXERCÍCIO 03

Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
	- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
	- Do quinto ao último item do slice (incluindo o último item!)
	- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
	- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/

package main
import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	
	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])

	fmt.Println(slice[2 : len(slice)-1])
}

/*
EXERCÍCIO 04

Começando com a seguinte slice:
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Anexe a ela o valor 52;
Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
Demonstre a slice;
Anexe a ela a seguinte slice:
	y := []int{56, 57, 58, 59, 60}
Demonstre a slice x.
*/

package main
import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{56, 57, 58, 59, 60}

	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	x = append(x, y...)
	fmt.Println(x)
}


/*
EXERCÍCIO 05

Começando com a seguinte slice:
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Utilizando slicing e append, crie uma slice y que contenha os valores:
[42, 43, 44, 48, 49, 50, 51]
*/

package main
import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := x[:3]
	y = append(y, x[6:]...)
	fmt.Println(y)
}

/*
EXERCÍCIO 06

Crie uma slice usando make que possa conter todos os estados do Brasil.
Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", 
"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", 
"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
Demonstre o len e cap da slice.
Demonstre todos os valores da slice sem utilizar range.
*/

package main
import "fmt"

func main() {
	estados := make([]string, 26, 26)
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",
		"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estados), cap(estados))

	for i := 0; i < len(estados); i++ {
		fmt.Println(estados[i])
	}
}

/*
EXERCÍCIO 07

Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main
import "fmt"

func main() {
	dados := [][]string{
		[]string{
			"Jhon", "Arias", "Futebol",
		},
		[]string{
			"German", "Cano", "Fazer gol",
		},
		[]string{
			"John", "Kennedy", "Ser craque",
		},
	}

	for _, v := range dados {
		fmt.PRintln(v)
	}
}


/*
EXERCÍCIO 08

Crie um map com key tipo string e value tipo []string.
Key deve conter nomes no formato sobrenome_nome
Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.
*/

package main
import "fmt"

func main() {
	dados := map[string][]string{
		"Pedro Sa": []string{"Videogame", "Altinha", "Academia"},
		"Gustavo":  []string{"Comer", "Churrasco", "Doce"},
		"Joao":     []string{"Altinha", "Nada", "Lolo"},
	}

	fmt.Println(dados)

	for i, v := range dados{
		fmt.Println(v)

		for indice, hobby := range v{
			fmt.Println("\t", indice, " - ", hobby)
		}
	}
}

/*
EXERCÍCIO 09

Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/

package main
import "fmt"

func main() {
	dados := map[string][]string{
		"Pedro Sa": []string{"Videogame", "Altinha", "Academia"},
		"Gustavo":  []string{"Comer", "Churrasco", "Doce"},
		"Joao":     []string{"Altinha", "Nada", "Lolo"},
	}

	fmt.Println(dados)

	dados["Grilo"] = []string{"Tchubanze", "Nathalia", "Academia"}

	for i, v := range dados {
		fmt.Println(i)

		for indice, hobby := range v {
			fmt.Println("\t", indice, " - ", hobby)
		}
	}
}


/*
EXERCÍCIO 10

Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range
*/

package main
import "fmt"

func main() {
	dados := map[string][]string{
		"Pedro Sa": []string{"Videogame", "Altinha", "Academia"},
		"Gustavo":  []string{"Comer", "Churrasco", "Doce"},
		"Joao":     []string{"Altinha", "Nada", "Lolo"},
	}

	fmt.Println(dados)

	// Adicioando mais um
	dados["Grilo"] = []string{"Tchubanze", "Nathalia", "Academia"}

	// Deletando um
	delete(dados, "Joao")

	for i, v := range dados {
		fmt.Println(i)

		for indice, hobby := range v {
			fmt.Println("\t", indice, " - ", hobby)
		}
	}
}
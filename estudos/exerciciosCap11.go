/*
EXERCÍCIO 01

Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
	- Nome
	- Sobrenome
	- Sabor favorito de sorvete
Crie dois valores do tipo "pessoa" e demonstre estes valores utilizando range na slice que contem os sabores de sorvete
*/

package main
import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	sabores []string
}

func main() {
	pessoa1 := pessoa{
		nome:      "Alberto",
		sobrenome: "Sa",
		sabores:   []string{"morango", "creme", "chocolate"},
	}

	pessoa2 := pessoa{"Irene", "Lagos", []string{"cupuaçu", "pavê", "flocos"}}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
}


/*
EXERCÍCIO 02

Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
Demonstre esses valores do map utilizando range.
Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior
*/

package main
import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	sabores []string
}

func main() {
	mapa := make(map[string]pessoa)

	pessoa1 := pessoa{
		nome:      "Alberto",
		sobrenome: "Sa",
		sabores:   []string{"morango", "creme", "chocolate"},
	}

	pessoa2 := pessoa{"Irene", "Lagos", []string{"cupuaçu", "pavê", "flocos"}}

	// Usando o sobrenome como índice/chave
	mapa[pessoa1.sobrenome] = pessoa1
	mapa[pessoa2.sobrenome] = pessoa2

	for sobrenome, pessoa := range mapa {
		fmt.Println("Sobrenome:", sobrenome)
		fmt.Println("Nome:", pessoa.nome)
		fmt.Println("Sabores preferidos:")
		for _, sabor := range pessoa.sabores {
			fmt.Println(" -", sabor)
		}
		fmt.Println()
	}
}


/*
EXERCÍCIO 03

Crie um novo tipo: veículo
Deve conter os campos: portas, cor
Crie dois novos tipos: caminhonete e sedan
Os tipos subjacentes devem ser struct e ambos devem conter "veículo" como struct embutido
O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
O tipo sedan deve conter um campo bool chamado "modeloLuxo"
Usando os structs veículo, caminhonete e sedan:
	Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
	Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
Demonstre esses valores
Demonstre um único campo de cada um dos dois
*/

package main
import "fmt"

type veiculo struct{
	portas int
	cor string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	meucarro := sedan{veiculo{4, "cinza"}, true}
	caminhaogalao := caminhonete{veiculo{2, "branca"}, false}

	fmt.Println(meucarro)
	fmt.Println(caminhaogalao)

	fmt.Println(meucarro.cor)
	fmt.Println(caminhaogalao.tracaoNasQuatro)
}


/*
EXERCÍCIO 04

Crie e use um struct anônimo
Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main
import "fmt"

func main() {
	time := struct {
		nome      string
		rival     string
		taças     []string
		jogadores map[int]string
	}{
		nome:  "Fluminense",
		rival: "Flamengo",
		taças: []string{"Libertadores", "Brasileirao", "Copa do Brasil", "Carioca"},
		jogadores: map[int]string{
			14: "German Cano",
			21: "Jhon Árias",
			10: "Paulo Henrique Ganso",
		},
	}

	fmt.Println(time)
}
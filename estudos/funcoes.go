// Sintaxe
// Função básica e função que aceita um argumento -------------------------------------------------
package main 
import "fmt"

func main() {

	//basica()
	argumento("manhã")
	argumento("tarde")
	argumento("jfdhjf")
}

func basica() {
	fmt.Println("Oi, bom dia!")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

// Função com retorno -----------------------------------------------------------------------------
package main 
import "fmt"

func main() {

	valor := soma(10, 10)

	fmt.Println(valor)
}

func soma(x, y int) int {
	return x + y
}

// Função com múltiplos retornos ------------------------------------------------------------------
package main 
import "fmt"

func main() {

	total, quantos, oi := soma("tarde", 10, 10, 1, 2, 3, 5)

	fmt.Println(total, quantos, oi)
}

func soma(s string, x ...int) (int, int, string) {
	oi := ""
	if s == "manhã" {
		oi = "Oi, bom dia!"
	} else if s == "tarde" {
		oi = "Oi, boa tarde!"
	} else {
		oi = "Oi, boa noite!"
	}
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma, len(x), oi
}

// Desenrolando (enumerando) uma slice
package main 
import "fmt"

func main() {

	si := []int{10, 10, 1, 2, 3, 5}
	total := soma(si...)

	fmt.Println(total)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}


// Defer
package main 
import "fmt"

func main() {
	defer fmt.Println("Deixa pra última hora")
	defer fmt.Println("Veio depois mas não vai ser o ultimo")
	fmt.Println("Vem primeiro")
}


// Método
package main
import "fmt"

type pessoa struct{
	nome string
	idade int
}

func (p pessoa) oibomdia() { // func (método) nome(parâmetro) (retorno) {code}
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {
	pedro := pessoa{"Pedro", 24}
	pedro.oibomdia()
}
// func (receiver) identifier(parameters) (returns) {code}
// Structs
package main
import "fmt"

type cliente struct{
	nome string
	sobrenome string
	fumante bool
}

func main() {
	// Duas formar de atribuição
	cliente1 := cliente{
		nome: "João",
		sobrenome: "da Silva",
		fumante: false,
	}

	cliente2 := cliente{"João", "Gomes", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}


// Structs embutidos
package main
import "fmt"

type pessoa struct{
	nome string
	idade int
}

type profissional struct{
	pessoa
	titulo string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Pedro",
		idade: 24,
	}

	pessoa2 := profissional{
		pessoa:  pessoa{"Grilo", 22},
		titulo:  "Técnico de Prótese Dentária",
		salario: 25000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

}


// Acessando apenas uma informação do struct
package main
import "fmt"

type pessoa struct{
	nome string
	idade int
}

type profissional struct{
	pessoa
	titulo string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Pedro",
		idade: 24,
	}

	pessoa2 := profissional{
		pessoa:  pessoa{"Grilo", 22},
		titulo:  "Técnico de Prótese Dentária",
		salario: 25000,
	}

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissional{pessoa{"Artur", 50}, "Major", 50000}

	fmt.Println(pessoa1.idade)
	fmt.Println(pessoa2.pessoa.salario)

}
/*
EXERCÍCIO 01

Demonstre todos os números de 1 a 10000
*/

package main
import "fmt"

func main() {
	for i:=1; i<=10000; i++{
		fmt.Println(i)
	}
}

/*
EXERCÍCIO 02

Demonstre na tela o unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada
*/

package main
import "fmt"

func main() {
	for i:=65; i<=90; i++{
		fmt.Println(i)
		for j:=1; j<=3; j++{
			fmt.Printf("\t%U\n", i)
		}
	}
}

/*
EXERCÍCIO 03

Crie um loop utilizando a sintaxe for condition{}
Utilize o loop para demonstrar os anos desde que você nasceu
*/

package main
import "fmt"

func main() {
	nasci := 2001
	anoQueroContar := 2100

	// Fica tipo um while
	for nasci <= anoQueroContar {
		fmt.Println(nasci)
		nasci++
	}
}

/*
EXERCÍCIO 04

Crie um loop utilizando a sintaxe for{}
Utilize o loop para demonstrar os anos desde que você nasceu
*/

package main
import "fmt"

func main() {
	nasci := 2001
	anoQueroContar := 2100

	// Fica tipo um while também
	for {
		if (nasci > anoQueroContar) {
			break
		}
		fmt.Println(nasci)
		nasci++
	}
}

/*
EXERCÍCIO 05

Demonster o resto da divisão por 4 de todos os números entre 10 e 100
*/

package main
import "fmt"

func main() {
	for i:=10; i<=100; i++ {
		fmt.Println(i%4)
	}
}

/*
EXERCÍCIO 06

Crie um programa que demonstre o funcionamento da declaração if
*/

package main
import "fmt"

func main() {
	cansado := true
	if cansado {
		fmt.Println("Não desista, continue estudando e buscando o melhor.")
	}
}

/*
EXERCÍCIO 07

Utilizando a solução anterior, adicione as opções else if e else
*/

package main
import "fmt"

func main() {
	cansado := "Demais"
	if cansado == "Muito" {
		fmt.Println("Não desista, continue estudando e buscando o melhor.")
	} else if cansado == "Pouco" {
		fmt.Println("Beleza, continue!")
	} else {
		fmt.Println("Força! Mentalize seu sonho e suas conquistas.")
	}
}

/*
EXERCÍCIO 08

Crie um programa utilizando a declaração switch, sem switch statement especificado
*/

package main
import "fmt"

func main() {
	estudando := "De madrugada"

	switch {
	case estudando == "De manhã":
		fmt.Println("Boa, começou bem")
	case estudando == "De noite":
		fmt.Println("Bem, tem que estudar mesmo")
	case estudando == "De madrugada":
		fmt.Println("Antes tarde do que nunca, lembre-se de descansar")
	}
}


/*
EXERCÍCIO 09

Crie um programa utilizando a declaração switch, onde o switch statement seja uma variável do tipo
string e o identificador "esporteFavorito"
*/

package main
import "fmt"

func main() {
	esporteFavorito := "Altinha"

	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Jogador de Futebol")
	case "Altinha":
		fmt.Println("Craque da praia")
	default:
		fmt.Println("Atleta")
	}
}
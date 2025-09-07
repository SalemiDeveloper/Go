/*
EXERCÍCIO 01

Escreva um programa que mostre um número em decimal, binário e hexadecimal
*/

package main

import "fmt"

func main() {
	a := 100
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
}

/*
EXERCÍCIO 02

Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis
==, !=, <=, <, >=, >
Demonstre esses valores
*/

package main

import "fmt"

func main() {
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f)
}

/*
EXERCÍCIO 03

Crie constantes tipadas e não tipadas.
Demonstre esses valores
*/

package main

import "fmt"

const x string = "tipado"
const y = "nao tipado"

func main() {
	fmt.Println(x, y)
}

/*
EXERCÍCIO 04

Crie um programa que
	- Atribua um valor int a uma variável
	- Demonstre esse valor em decimal, binário e hexadecimal
	- Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	- Demonstre essa outra variável em decimal, binário e hexadecimal
*/

package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}

/*
EXERCÍCIO 05

Crie uma variável do tipo string utilizando uma raw string literal e demonstre-a
*/
package main

import "fmt"

func main() {
	a := `Isso
			é uma raw
			string`
	fmt.Println(a)
}

/*
EXERCÍCIO 06

Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos e demonstre.
*/

package main

import "fmt"

const (
	_ = 2001 + iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(b, c, d, e)
}

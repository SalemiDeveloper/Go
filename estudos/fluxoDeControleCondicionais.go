/*
If, else if e else

Apenas anotando sobre IF, ELSE IF e ELSE pois na hora de escrever é um pouco
diferente das linguagens que sei
*/

package main
import "fmt"
func main() {
	if x := 10; x < 100 {
		fmt.Println("X é menor que 100")
	} else if x > 100 {
		fmt.Println("X é maior que 100")
	} else {
		fmt.Println("X não é maior nem menor que 100")
	}
}

/*
Switch

fallthrough faz com que se o case que ele estiver acontecer, o debaixo também irá
*/

package main
import "fmt"

func main() {
	x:=10

	switch {
		case x<5:
			fmt.Println("X é menor que 5")
			fallthrough
		case x==5:
			fmt.Println("X é igual a 5")

		case x > 5:
			fmt.Println("X é maior que 5")

		default:
			fmt.Println("X não é nada")
	}
}

/*
var x interface{}

A interface{} é um tipo especial que pode armazenar qualquer valor, de qualquer tipo. 
É como se fosse o “supertipo” de tudo - similar ao “Object” em outras linguagens como Java ou C#.

Em outras palavras, interface{} aceita qualquer tipo, mas quando for para usar o valor armazenado, 
precisa fazer um type assertion (checar com type switch)
*/


// type assertion
package main
import "fmt"

var x interface{}

func main() {
	x = 99
	n := x.(int) // converte interface{} para int
	fmt.Println(n + 1) // vai imprimir 100
}

// type switch
package main
import "fmt"

var x interface{}

func main() {
	switch v := x.(type) {
		case int:
			fmt.Println("É um int: ", v)
		case string:
			fmt.Println("É uma string: "v)
		default:
			fmt.Println("Tipo desconhecido")
	}
}

// Fluxo de Controle - Loop For
package main
import "fmt"
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

// Treinando Loop com mês e dia
package main

import "fmt"

func main() {
	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês: ", mes)
		for dia := 1; dia <= 31; dia++ {
			fmt.Print("Dia: ", dia, ", ")
		}
		fmt.Println(" ")
	}
}

/*
EXERCÍCIO DE LOOP

Faça um loop do número 33 até 122, e utilize o format print para demonstrá-los como texto/string
Decimal       %d
Hexadecimal   %#x
Unicode       %#U
Tab           \t
Linha nova    \n
*/

package main

import "fmt"

func main() {
	for x:=33; x <= 122; x++ {
		fmt.Printf("%d\t%#x\t%#U\t", x, x, x)
	}
}
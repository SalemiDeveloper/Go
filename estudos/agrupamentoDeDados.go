/*
ARRAY

Treinando Arrays em Go pois são diferentes que em PHP
Em Go são fortemente tipados.
*/ 
package main

import "fmt"

func main() {
	numerosPares := [3]int{}
	numerosImpares := [3]int{1, 3, 5}

	numerosPares[0] = 2
	numerosPares[1] = 4
	numerosPares[2] = 6

	fmt.Printf("%v\n", numerosPares)
	fmt.Printf("%v", numerosImpares)
}

/*
SLICE
*/

package main
import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Não é possível com array
	slice2 := append(slice, 6)
	fmt.Println(slice2)
}


// Range no Slice

package main
import "fmt"

func main() {
	sliceString := []string{"banana", "maçã", "jaca"}

	for indice, valor := range sliceString {
		fmt.Println("No índice:", indice, "temos o valor:", valor)
	}

	sliceInt := []int{1, 2, 3, 4, 5}
	total := 0

	for _, valor := range sliceInt {
		total = total + valor
	}
	fmt.Println("O total é:", total)
}


// Slice de Slice

package main
import "fmt"

func main() {
	sabores := []string{"calabresa", "frango", "queijos"}

	fatia := sabores[2:3] // [indice que começa : até o índice que vai parar porém sem contar com ele]
	fmt.Println(fatia)    // [a:b] "a" é incluse "b" não é incluso
}

// Anexando a uma Slice

package main
import "fmt"

func main() {

	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...)

	fmt.Println(umaslice)
}


// Slice Make

package main
import "fmt"

func main() {
	slice := make([]int, 5, 10)
	// "5" tamamho inicial do slice (0, 0, 0, 0, 0)
	// "10" tamanho da capacidade reservada em memória, permitindo crescer em até 10 elementos sem realocação

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))
}


// Slices multi dimensionais
package main
import "fmt"

func main() {
	tabela := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(tabela[0])     // [1 2 3]
	fmt.Println(tabela[1][2])  // 6
}


// Maps
package main
import "fmt"

func main() {
	amigos := map[string]int{
		"primo": 98908,
		"lg":    99929,
		"grilo": 99321,
	}

	// Adicionando um item novo
	amigos["brinquedo"] = 96880

	fmt.Println(amigos)
	fmt.Println(amigos["grilo"])
	fmt.Println(amigos["lg"])
	
	// comma ok idiom
	if será, ok := amigos["primo"]; !ok {
		fmt.Println("Não tem")
	} else {
		fmt.Println(será)
	}
}


// Maps - Range e Delete
package main
import "fmt"

func main() {
	julgando := map[int]string{
		10: "Muito bom",
		8:  "Bom",
		5:  "Médio",
		3:  "Ruim",
		1:  "Muito ruim",
	}

	fmt.Println(julgando)

	for key, value := range julgando {
		fmt.Println(key, value)
	}

	delete(julgando, 10)

	fmt.Println(julgando)
}
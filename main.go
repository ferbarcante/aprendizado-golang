package main

import (
	"fmt"
	"strconv"
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")

}

func soma(n1, n2 int) int {
	return n1 + n2
}

// dentro de funções é possivel retornar mais de um valor
func convertAndSum(a int, b string) (total int, err error) {

	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i

	return
}

// feita para iniciar o programa
func main() {

	hello("Maria Fernanda")

	fmt.Println("Total:", soma(10, 5))

	total, err := convertAndSum(2, "1")
	fmt.Println("total:", total, err)
}

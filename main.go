// O arquivo main é o ponto de entrada pro programa GO
// É preciso definir algumas coisas: precisa do pacote main, isso ja vai garantir que o arquivo pertence ao package principal do programa
package main

//uma regra de package nao pode ter na mesma pasta dois packages diferentes
import "fmt"

var (
	//variavel a nivel de package, disponiveis a qualquer arq que tiver nesse package
	nome string
	n1   int
	n2   int
	N3   int //quando a var for iniciada com letra maiuscula da pra importar em outros packages
)

//feita para iniciar o programa
func main() {
	nome = "Fernanda"
	fmt.Println("Oi", nome, "!")

	mensagem := "Aula 03 - GO 101" //os := é a mesma coisa q var mensagem
	fmt.Println(mensagem)

	var total int
	fmt.Println("total:", total)
	total++
	fmt.Println("total:", total)

	//multiplas definicoes
	/* var a, b, c int32 */
	//multiplas atribuicoes
	var d, e, f = true, 2, "oi"
	fmt.Println(d, e, f)

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}

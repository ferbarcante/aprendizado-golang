# Introdução
O arquivo main é o ponto de entrada pro programa GO

## Package
É preciso definir algumas coisas: precisa do pacote main, isso ja vai garantir que o arquivo pertence ao package principal do programa
Uma regra de package nao pode ter na mesma pasta dois packages diferentes

## Variáveis
### Nivel de package
Estão disponiveis a qualquer arquivo que tiver nesse package
**exemplo:**
```go
var (
	n1   int
	n2   int
	N3   int 
    //quando a var for iniciada com letra maiuscula da pra importar em outros packages
)
```
### Atribuições
```go 
mensagem := "Aula 03 - GO 101" 
//os := é a mesma coisa q var mensagem 
```
## Funções
É uma maneira de reaproveitar um código, ideia de que sejam menores possiveis para que o sistema seja + reutilizavel e facilitar manutenção

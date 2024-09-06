package main

import "fmt"

func main() {
	//y recebe 5 aqui, as 2 variáveis tem o msm valor
	x := 5
	y := &x
	*y = 10
	fmt.Println(x, y)
	fmt.Println(x, *y)
	printValues(&x, y)
}

func printValues(x *int, y *int) {
	*x = 20
	fmt.Println(x, y)
}

/*
Nesse caso agora, as variaveis são independentes pois ambas tem sua alocação de memória
Elas apontam para lugares diferentees então qd ocorre a reatribuição, ficam diferentes
	x := 5
	y := x
	y = 10

alocação de memória qd envolve funções com parâmetros fica uma crazy shit
nesse caso agora, x e y tem valores diferentes de memória em diferentes lugares, ele fez uma cópia, 2 variáveis ocupam 4 lugares na memória...
	x := 5
	y := x
	y = 10
	fmt.Println(x, y)
	fmt.Println(&x, &y)
	printValues(x, y)
}

func printValues(x int, y int) {
	fmt.Println(x, y)
	fmt.Println(&x, &y)
}


da pra resolver isso com pointers

Quer utilizar a mesma referencia na memória? Usa pointers, usa &
Altera o valor de y? altera x tb
*/

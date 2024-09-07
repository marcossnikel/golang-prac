package basics

import "fmt"

func exercisie() {
	//criar um array de 2 posições e armazenar em uma variável a soma total da lista, imprimir no console

	nArr := [2]int{523, 321}
	sum := 0
	for i := 0; i < len(nArr); i++ {
		sum += nArr[i]
	}
	fmt.Println(sum)

	//Slice com os itens 2,8,3,10,5,4,7,9,1, que vao de 1 a 10, efetuar a some em duas novas variaveis, uma de numeros de 1 a 5 e outra de numeros de 6 a 10, imprimir os dois reuslts

	list := []int{2, 8, 3, 10, 5, 4, 7, 9, 1}
	lessThenFive := 0
	greatherThenFive := 0
	for i := 0; i < len(list); i++ {
		if list[i] <= 5 {
			lessThenFive += list[i]
		} else {
			greatherThenFive += list[i]
		}
	}

	fmt.Println(lessThenFive)
	fmt.Println(greatherThenFive)

}

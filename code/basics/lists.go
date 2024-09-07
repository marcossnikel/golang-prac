package basics

import "fmt"

func lists() {
	list := []int{1, 2, 3, 4, 5, 42, 78, 32, 42, 52}
	listWithNumbersGreatherThanFive := make([]int, 0)
	for i := 0; i < len(list); i++ {
		if list[i] > 10 {
			listWithNumbersGreatherThanFive = append(listWithNumbersGreatherThanFive, list[i])
		}
	}
	sliceList := list[2:6] // first tree items
	fmt.Println(sliceList)
	sliceBeggining := list[:5]

	lastItem := list[len(list)-1:]
	fmt.Println(lastItem)
	fmt.Println(sliceBeggining)
	fmt.Println(listWithNumbersGreatherThanFive)

}

/*
slice x array
slice -> lista podem ser aumentadas...
slice-> no colchete n coloca o tamanho
array -> coloca o tamanho no colchete
append é para slices, array não usa o append

array é MUITO mais performático. qdo performance é essencial é bom utilizar arrays
*/

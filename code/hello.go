package main

import "fmt"

func main() {
	salary := 3000.00
	var salaryPlusBonus float64
	salaryPlusBonus = salary
	var bonus float64 = 1.10
	if salary < 1000 {
		salaryPlusBonus = (salaryPlusBonus * bonus)
	} else if salary >= 1000 && salary < 2000 {
		bonus = 1.30
		salaryPlusBonus = (salaryPlusBonus * bonus)
	} else {
		bonus = 1.50
		salaryPlusBonus = (salaryPlusBonus * bonus)
	}
	fmt.Println(salaryPlusBonus)
}

func ifelse() {
	var isCar bool = false
	vehichlePrice := 1000.00

	if isCar {
		vehichlePrice += 55.50
	} else {
		vehichlePrice *= 25.50
	}

	fmt.Println("Final price: ", vehichlePrice)
}

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	type Employee struct {
// 		Name               string
// 		Salary             float64
// 		DescountPercentage float64
// 	}

// 	employeesArray := []Employee{
// 		{
// 			Name:               "John",
// 			Salary:             3000.00,
// 			DescountPercentage: 0.10,
// 		},
// 		{
// 			Name:               "Mary",
// 			Salary:             5000.00,
// 			DescountPercentage: 0.15,
// 		},
// 		{
// 			Name:               "Paul",
// 			Salary:             7000.00,
// 			DescountPercentage: 0.20,
// 		},
// 		{
// 			Name:               "Marcos Nikel",
// 			Salary:             10980.00,
// 			DescountPercentage: 0.27,
// 		},
// 	}
// 	for _, employee := range employeesArray {
// 		calculateDiscount(employee.Salary, employee.DescountPercentage, employee.Name)
// 	}
// 	nestedLoops(10)
// 	lists()

// }

// func calculateDiscount(salary float64, descountPercentage float64, name string) {
// 	salary = salary - (salary * descountPercentage)
// 	fmt.Println("Final salary for you " + name + " R$" + fmt.Sprintf("%.2f", salary))
// }

// func nestedLoops(expoentNumber int) {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < expoentNumber; j++ {
// 			fmt.Println(strconv.Itoa(i) + " x " + strconv.Itoa(j) + " = " + strconv.Itoa(i*j))
// 		}
// 	}
// }

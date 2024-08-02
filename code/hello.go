package main

import "fmt"
func main(){
    salary := 999.00
    var salaryPlusBonus float64
    salaryPlusBonus = salary
    var bonus float64 = 1.10
    if salary < 1000 {
        salaryPlusBonus = (salaryPlusBonus * bonus)
    }
    fmt.Println(salaryPlusBonus)
}
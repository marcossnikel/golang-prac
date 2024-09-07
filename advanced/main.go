package main

import (
	"advanced_go_study/model"
	"fmt"
	"time"
)

func main() {
	println("Initializing...")

	address := model.Address{
		Street: "Salvador Vaccari",
		Number: 336,
		City:   "Jundiaí",
	}
	me := model.Person{
		Name:      "Marcos",
		Address:   address,
		BirthDate: time.Date(2002, 04, 04, 23, 23, 0, 0, time.Local),
	}
	fmt.Println(me)
	address.Street = "Rua Salvador Vaccari"
	fmt.Println(address)

	me.CalculateAge()
	fmt.Println(me.Age)
}

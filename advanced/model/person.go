package model

import "time"

type Person struct {
	Name      string
	Address   Address
	BirthDate time.Time
	Age       int
}

func (p *Person) CalculateAge() {
	birthYear := p.BirthDate.Year()
	currentYear := time.Now().Year()
	p.Age = currentYear - birthYear
}

// func CalculateAge(p Person) int {
// 	birthYear := p.BirthDate.Year()
// 	currentYear := time.Now().Year()
// 	return currentYear - birthYear
// }

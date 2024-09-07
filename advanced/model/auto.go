package model

type Auto struct {
	Year  int
	Plate string
	Model string
}

type Motocycle struct {
	Auto
	Cilind int
}

type Car struct {
	PortQuantity   int
	Potency        int
	AirConditioner bool
}

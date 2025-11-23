package main

import "fmt"

type Wheel struct{
	Material string
	Radius string
}

type car struct{
	Make string
	Model int
	Weght float32
	Wheel Wheel
}

func main() {

	myCar := car{}
	myCar.Wheel.Radius = "5"
	fmt.Println("Welcome to Textio,", myCar)
}

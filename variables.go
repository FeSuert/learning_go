package main

import "fmt"

func variables() {

	//strings

	var nameOne string = "Luigi"
	var nameTwo = "Mario"
	var nameThree string
	nameFour := "Peach"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	//ints

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory

	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	//floats

	var scoreOne float32 = 12.22
	var scoreTwo float64 = 534534534534.5

	fmt.Println(scoreOne, scoreTwo)
}

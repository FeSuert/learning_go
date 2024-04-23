package main

import "fmt"

func main() {
	//arrays
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [2]string{"Mario", "Luigi"}

	names[1] = "Peach"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := ages[0:1]
	rangeTwo := ages[1:]
	rangeThree := ages[:2]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}

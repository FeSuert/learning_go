package main

import "fmt"

func formatting() {
	age := "35"
	name := "Mario"

	//print
	fmt.Print("Hello")
	fmt.Print("World \n")
	fmt.Print("New line \n")

	//Println

	fmt.Println("Hello Ninjas")
	fmt.Println("Goodbye Ninjas")

	fmt.Println("My age is", age, "and my name is", name)

	//Printf (formatted strings)

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("you scored %0.2f points! \n", 2.55)

	//Sprintf (save formatted strings)

	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)
}

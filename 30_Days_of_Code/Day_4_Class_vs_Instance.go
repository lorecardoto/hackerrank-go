package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	//Add some more code to run some checks on initialAge
	//initialAge is not negative
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		//if a negative argument is passed as initialAge
		initialAge = 0
	}
	p.age = initialAge
	return p
}

func (p person) amIOld() {
	//Do some computation in here and print out the correct statement to the console
	//If age < 13, print
	if p.age < 13 {
		fmt.Println("You are young.")
	} else if p.age >= 13 && p.age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are old.")
	}

}

func (p person) yearPasses() person {
	//Increment the age of the person in here
	//yearPasses() should increase the AGE instance variable by 1.

	p.age += 1
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}

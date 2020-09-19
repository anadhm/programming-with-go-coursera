package main

import (
	"fmt"
)

//Animal is the required struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

//Eat prints what the animal eats
func (a Animal) Eat() {
	fmt.Print(a.food)
}

//Move prints how the animal moves
func (a Animal) Move() {
	fmt.Print(a.locomotion)
}

//Speak prints the sound the animal makes
func (a Animal) Speak() {
	fmt.Print(a.noise)
}
func main() {
	animals := make(map[string]Animal)
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	animals["cow"] = cow
	animals["bird"] = bird
	animals["snake"] = snake
	var name, information string
	for {
		fmt.Print(">")
		fmt.Scan(&name, &information)
		if name != "cow" && name != "bird" && name != "snake" {
			fmt.Println("Wrong animal name, try again.")
			continue
		}
		if information != "eat" && information != "move" && information != "speak" {
			fmt.Println("Wrong animal information, try again.")
			continue
		}
		if information == "eat" {
			animals[name].Eat()
		}
		if information == "move" {
			animals[name].Move()
		}
		if information == "speak" {
			animals[name].Speak()
		}

		fmt.Println()

	}

}

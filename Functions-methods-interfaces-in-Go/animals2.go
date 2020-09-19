package main

import (
	"fmt"
)

//Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	food, move, sound string
}
type Bird struct {
	food, move, sound string
}
type Snake struct {
	food, move, sound string
}

func newCow() Cow {
	return Cow{"grass", "walk", "moo"}
}
func newBird() Bird {
	return Bird{"worms", "fly", "peep"}
}
func newSnake() Snake {
	return Snake{"mice", "slither", "hsss"}
}
func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (b Bird) Eat() {
	fmt.Println(b.food)
}
func (s Snake) Eat() {
	fmt.Println(s.food)
}
func (c Cow) Move() {
	fmt.Println(c.move)
}
func (b Bird) Move() {
	fmt.Println(b.move)
}
func (s Snake) Move() {
	fmt.Println(s.move)
}
func (c Cow) Speak() {
	fmt.Println(c.sound)
}
func (b Bird) Speak() {
	fmt.Println(b.sound)
}
func (s Snake) Speak() {
	fmt.Println(s.sound)
}
func main() {
	var command, name, info string
	for {
		fmt.Print(">")
		fmt.Scan(&command, &name, &info)
		if command == "newanimal" {
			var a Animal
			switch info {
			case "cow":
				a = newCow()
				_ = a
				fmt.Println("Created it!")
				continue
			case "bird":
				a = newBird()
				_ = a
				fmt.Println("Created it!")
				continue
			case "snake":
				a = newSnake()
				_ = a
				fmt.Println("Created it!")
				continue
			default:
				fmt.Println("Wrong type of animal. Retry.")
				continue
			}
		}
		if command == "query" {
			var a Animal
			switch name {
			case "cow":
				a = newCow()
				switch info {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Wrong query. Retry.")
					continue
				}
				continue
			case "bird":
				a = newBird()
				switch info {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Wrong query. Retry.")
					continue
				}
				continue
			case "snake":
				a = newSnake()
				switch info {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				default:
					fmt.Println("Wrong query. Retry.")
					continue
				}
				continue
			default:
				fmt.Println("Wrong type of animal. Retry.")
				continue
			}
		}
	}

}

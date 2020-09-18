package main

import "fmt"

func main() {
	var number float64
	fmt.Println("Please input a floating point number: ")
	fmt.Scan(&number)
	fmt.Println(int(number))
}

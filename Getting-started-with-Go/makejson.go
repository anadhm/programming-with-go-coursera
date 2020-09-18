package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, address string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a name:")
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Println("Please enter an address:")
	if scanner.Scan() {
		address = scanner.Text()
	}
	m := make(map[string]string)
	m[name] = address
	obj, _ := json.Marshal(m)
	fmt.Println(string(obj))
}

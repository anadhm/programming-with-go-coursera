package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var phrase string
	/* Doesn't work for strings with spaces */
	//fmt.Scanln(&phrase)
	/* https://forum.golangbridge.org/t/string-with-spaces/10617/3 */
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		phrase = strings.ToLower(scanner.Text())
		if strings.HasPrefix(phrase, "i") && strings.HasSuffix(phrase, "n") && strings.Contains(phrase, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Error reading input")
	}

}

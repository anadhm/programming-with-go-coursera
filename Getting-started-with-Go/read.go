package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const size = 20

type name struct {
	fname string
	lname string
}

func main() {
	names := make([]name, 0, 10)
	var filename string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Name of the text file:")
	if scanner.Scan() {
		filename = scanner.Text()
	}
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileReader := bufio.NewScanner(f)
	/* assume the text file is correct */
	for fileReader.Scan() {
		text := fileReader.Text()
		sli := strings.Split(text, " ")
		fname := sli[0]
		lname := sli[1]
		n := new(name) //makes a pointer of type name
		if len(fname) > size {
			fname = fname[:size]
		}
		if len(lname) > size {
			lname = lname[:size]
		}
		n.fname = fname
		n.lname = lname
		names = append(names, *n)
	}

	for _, n := range names {
		fmt.Printf("First name: %s, Last name: %s\n", n.fname, n.lname)
	}

	f.Close()

}

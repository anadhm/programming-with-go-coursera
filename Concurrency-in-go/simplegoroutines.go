package main

import (
	"fmt"
	"time"
)

func f(x *int, name string) {
	*x = *x + 1
	fmt.Printf("Name %s: x=%d\n", name, *x)
}

func main() {
	x := 0
	for i := 0; i < 5; i++ {
		go f(&x, "first")
		go f(&x, "second")
		time.Sleep(5 * time.Millisecond)
		x++
	}
	fmt.Printf("Main: x=%d\n", x)
}

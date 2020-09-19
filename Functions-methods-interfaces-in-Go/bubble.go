/*This is a package description. */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*BubbleSort : Performs Bubblesorting on given slice sli */
func BubbleSort(sli []int) {
	var i, j int
	for i = 0; i < len(sli); i++ {
		for j = 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

/*Swap : Swaps sli[index] with sli[index+1] */
func Swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input a sequence of up to 10 integers.")
	var numbers []int
	if scanner.Scan() {
		/* Parse the integers */
		text := scanner.Text()
		strNumbers := strings.Split(text, " ")
		// takes first 10 numbers if more than 10 are provided
		if len(strNumbers) > 10 {
			strNumbers = strNumbers[:10]
		}
		for _, n := range strNumbers {
			intNumber, err := strconv.Atoi(n)
			if err != nil {
				fmt.Print(err)
				break
			} else {
				numbers = append(numbers, intNumber)
			}

		}
		BubbleSort(numbers)
		fmt.Println(numbers)
	}
}

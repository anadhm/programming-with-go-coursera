package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func subr(sli []int, c chan []int) {
	fmt.Println("This subroutine will sort the subarray:", sli)
	sort.Ints(sli)
	c <- sli
}
func merge(s1 []int, s2 []int) []int { //This is the merge subroutine of MergeSort
	n1 := len(s1)
	n2 := len(s2)
	var newSlice []int
	if n1 > n2 {
		min := n2
		i := 0
		j := 0
		for i < min && j < min {
			if s1[i] < s2[j] {
				newSlice = append(newSlice, s1[i])
				i++
			} else {
				newSlice = append(newSlice, s2[j])
				j++
			}
		}
		if i >= min {
			newSlice = append(newSlice, s2[j:]...)
		} else {
			newSlice = append(newSlice, s1[i:]...)
		}

	} else {
		min := n1
		i := 0
		j := 0
		for i < min && j < min {
			if s1[i] < s2[j] {
				newSlice = append(newSlice, s1[i])
				i++
			} else {
				newSlice = append(newSlice, s2[j])
				j++
			}
		}
		if i >= min {
			newSlice = append(newSlice, s2[j:]...)
		} else {
			newSlice = append(newSlice, s1[i:]...)
		}
	}
	return newSlice
}
func main() {
	var numbers []int
	c := make(chan []int, 4)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input a series of integers:")
	// Read the slice of integers
	if scanner.Scan() {
		t := scanner.Text()
		for _, f := range strings.Fields(t) {
			i, err := strconv.Atoi(f)
			if err == nil {
				numbers = append(numbers, i)
			}
		}
	}
	n := len(numbers)
	go subr(numbers[:n/4], c)
	go subr(numbers[n/4:n/2], c)
	go subr(numbers[n/2:3*n/4], c)
	go subr(numbers[3*n/4:], c)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c
	arr5 := merge(arr1, arr2)
	arr6 := merge(arr3, arr4)
	//fmt.Println("arr5 is: ", arr5)
	//fmt.Println("arr6 is: ", arr6)
	arr7 := merge(arr5, arr6)
	fmt.Println("Sorted array is: ", arr7)
}

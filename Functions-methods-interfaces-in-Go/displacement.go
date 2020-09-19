package main

import (
	"fmt"
)

/*GenDisplaceFn generates a function that computes displacement after given time, with accel,vel and initialDis as initial conditions */
func GenDisplaceFn(accel, vel, initialDis float64) func(float64) float64 {
	return func(time float64) float64 {
		return (accel/2)*time*time + vel*time + initialDis
	}
}

func main() {
	var a, v, s0, t float64
	fmt.Println("Please input value for acceleration, inital velocity and initial displecement: ")
	fmt.Scan(&a, &v, &s0)
	computeDisplace := GenDisplaceFn(a, v, s0)
	fmt.Println("Please enter a value for time: ")
	fmt.Scan(&t)
	fmt.Printf("Displacement after %f time is %f.\n", t, computeDisplace(t))
}

package main

import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	x1, x2 := 0.0, 0.0
	descriminant := b*b - 4*a*c
	if descriminant < 0 {
		fmt.Println("This equation has no real solution")
	} else if descriminant == 0 {
		x1 = (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
		x2 = x1
		fmt.Println("This equation has one solutions: ", x)
	} else {
		x1 = (-b + math.Sqrt((b*b)-(4*(a*c)))) / (2 * a)
		x2 = (-b - math.Sqrt((b*b)-(4*(a*c)))) / (2 * a)
		fmt.Println("This equation has two solutions: ", x1, " or", x2)
	}
	return x1, x2
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}

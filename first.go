package main

import "fmt"

func main() {
	fmt.Println("First main package Call")

	squareVal := func(x float64) float64 {
		return x * x
	}

	fmt.Println("Value returned", squareVal(3))
	nextVSl := getSequence()
	fmt.Println(nextVSl())
	fmt.Println(nextVSl())
	fmt.Println(nextVSl())
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

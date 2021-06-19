package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	// we can do something before make decision
	// v's scope is only in the scope of if
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// we can't use v here
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Print("Go runs on ")

	// we do not use break in Colang
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	num := 10
	switch num {
	case 1, 2, 3:
		fmt.Println("num is 1, 2 or 3")
	case 10, 11, 12:
		fmt.Println("num is 10, 11 or 12")
	}
}

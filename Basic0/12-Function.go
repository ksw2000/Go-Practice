package main

import "fmt"

// basic type
func add(x, y int) int {
	return x + y
}

// named return values
func model(x, y int) (result int) {
	return x % y
}

func mul(x, y int) (result int) {
	result = x * y
	return
}

// variadic function
func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func main() {
	// basic calling
	fmt.Println(add(1, 2))
	fmt.Println(model(9, 4))
	fmt.Println(mul(3, 4))
	fmt.Println()

	// funcion can see as a variable
	f := add
	fmt.Printf("%T\n", f) // func(int, int) int
	fmt.Printf("f(10,10) = %d\n", f(10, 10))

	var g func(int, int) int
	g = mul
	fmt.Printf("g(10,10) = %d\n", g(10, 10))

	fmt.Println()
	fmt.Printf("sum(8,7) = %d\n", sum(8, 7))
	fmt.Printf("sum(1,2,3,4) = %d\n", sum(1, 2, 3, 4))

	// we can use slice as parameters when calling variadic function
	slice := []int{8, 9, 10, 11}
	fmt.Printf("sum(slice...) = %d\n", sum(slice...))
}

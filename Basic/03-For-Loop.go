package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {
		for j := 2; j < 10; j++ {
			fmt.Printf("%dx%d=%-2d ", i, j, i*j)
		}
		fmt.Printf("\n")
	}

	i := 0
	for i < 10 { // while(i < 10) in C
		fmt.Printf("%d ", i)
		i++
	}

outer:
	for i := 0; i <= 10; i++ {
		for j := 0; j <= 10; j++ {
			if i+j == 10 {
				break outer
			}
			fmt.Printf("(%d, %d)\n", i, j)
		}
	}
}

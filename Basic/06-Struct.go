package main

import "fmt"

type vertex struct {
	x int
	y int
}

type Circle struct {
	vertex // inherit
	radius float64
}

func main() {
	v := vertex{1, 2}
	v.x = 4
	fmt.Println(v.x, v.y)

	var c Circle
	c.x = 0
	c.y = 0
	c.radius = 10
	fmt.Printf("%v\n", c)  // default format
	fmt.Printf("%#v\n", c) // Go-syntax representation of the value
}

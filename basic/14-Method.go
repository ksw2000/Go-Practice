package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Inner(u Vertex) float64 {
	return v.X*u.X + v.Y*u.Y
}

// 指標接收器
func (v *Vertex) Add(u Vertex) Vertex {
	v.X += u.X
	v.Y += u.Y
	return *v
}

// 建立一個 method
// https://tour.golang.org/methods/3
type list []int

func (arr list) bubbleSort() {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	w := Vertex{6, 7}
	fmt.Println(w, "dot", "{1, -2} =", w.Inner(Vertex{1, -2}))

	z := &Vertex{1, 2}
	fmt.Println(z.Add(Vertex{3, 8}))

	arr := list{12, 32, 43, 11, 54, 25, 37}
	arr.bubbleSort()
	fmt.Println(arr)
}

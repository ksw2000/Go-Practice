package main

import (
	"fmt"
	"sort"
)

/*
sort package
type Interface interface {
	Len is the number of elements in the collection.
	Len() int

	Less reports whether the element with index i
	must sort before the element with index j.

	If both Less(i, j) and Less(j, i) are false,
	then the elements at index i and j are considered equal.
	Sort may place equal elements in any order in the final result,
	while Stable preserves the original input order of equal elements.

	Less must describe a transitive ordering:
	 - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	 - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.

	Note that floating-point comparison (the < operator on float32 or float64 values)
	is not a transitive ordering when not-a-number (NaN) values are involved.
	See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

type Pair struct {
	key int
	val string
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].key-p[j].key < 0
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p1 := Pair{
		key: 2009,
		val: "Go",
	}
	p2 := Pair{
		key: 2010,
		val: "Rust",
	}
	p3 := Pair{
		key: 2011,
		val: "Kotlin",
	}
	p4 := Pair{
		key: 2011,
		val: "Dart",
	}
	p5 := Pair{
		key: 2012,
		val: "typescript",
	}
	list := Pairs{p1, p3, p2, p4, p5}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)

	// Interfaces in Go provide a way to specify the behavior of an object:
	// if something can do this, then it can be used here.
	// Pairs can do `Len()`, `Less()` and `Swap()`
	// so we can consider `Pairs` as `sort.Interface()`
}

package main

import "fmt"

func main() {
	// 類型標記法
	var array1 [3]string
	array1[0] = "C"
	array1[1] = "C++"
	array1[2] = "GO"
	fmt.Println(array1)

	// 值標記法
	// := 宣告 + 賦值 限函數內
	array2 := [3]string{"C", "C++", "GO"}
	fmt.Println(array2)

	array3 := [...]string{"C", "C++", "GO"}
	fmt.Println(array3)

	array4 := [6]string{0: "C", 1: "C++", 2: "GO", 3: "PHP"}
	fmt.Println(array4)
}

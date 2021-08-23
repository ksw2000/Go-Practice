package main

import "fmt"

func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		temp := back1
		back1, back2 = back2, (back1 + back2)
		return temp
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	//demo fibonacci()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	// 0 1 1 2 3 5 8 13 21 34

	//demo intSeq
	fmt.Println()

	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println()

	newNextInt := intSeq()
	fmt.Println("update...")
	fmt.Println(newNextInt()) // 1
	fmt.Println(newNextInt()) // 2
	fmt.Println(newNextInt()) // 3
}

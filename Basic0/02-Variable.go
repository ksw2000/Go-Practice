package main

import "fmt"

const boilingF = 212.0

func main() {
	/*
	   bool            true, false
	   byte
	   rune            Unicode (32bits)
	                   rune is an alias for int32.
	   int/uint
	   int8/unit8      8bits
	   int16/uint16    16bits
	   int32/uint32    32bits (long int in C)
	   int64/uint64    64bits (long long int in C)
	   float32         32bits (float in C)
	   float64         64bits (double in C)
	   complex64       float32(real part)+float32(imaginary part)
	   complex128      float64(real part)+float64(imaginary part)
	   string          Once string is declared, it will be unchangeable
	*/

	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)

	var (
		year  = 2005
		month = "Jan"
		day   = 30
		name  = "maikichi"
	)
	fmt.Printf("%s was born on %d %s, %d\n", name, day, month, year)

	const (
		e  = 2.71828182846 //常數不能使用 := 語法定義。
		pi = 3.14159
	)
	fmt.Println("e =", e)
	fmt.Println("pi =", pi)

	var a complex64
	var b rune
	a = 1 + 5i
	b = '\u0022'
	fmt.Println(a)
	fmt.Printf("%c", b)
}

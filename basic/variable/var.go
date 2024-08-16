package main

import "fmt"

func main() {
	// VARIABLE
	var a = 1

	s := "hello"

	fmt.Println(a, s)

	var (
		b = 2
		y = "hai"
	)
	fmt.Println(b, y)

	// CONSTANT
	const c = 3
	fmt.Println(c)

	const (
		d = 4
		z = "bye"
	)
	fmt.Println(d, z)

	// CONVERSION
	var f = string(a)
	fmt.Println(f)

	var g = string(s[0])
	fmt.Println(g)

}

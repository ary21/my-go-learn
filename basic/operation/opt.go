package main

import "fmt"

func main() {
	var a = 1
	var b = 2
	a++
	b += 3
	fmt.Println("hai", a+b)

	fmt.Println(a == b, a != b)

	var x bool = a > b && a > 1
	fmt.Println(x)
}

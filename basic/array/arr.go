package main

import "fmt"

func main() {
	var nums [3]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3

	fmt.Println(nums)

	var ar = [...]int{1, 2, 3, 4}
	var arr = [3]int{1, 2, 3}
	fmt.Println(ar, arr)

	var arr2 = [100]int{}
	fmt.Println(arr2)
	fmt.Println(len(arr)) // length
}

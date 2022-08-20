package main

import "fmt"

func main() {
	var x int

	fmt.Println("enter no.")
	fmt.Scanln(&x)
	fmt.Println(cube(x))

}
func cube(x int) int {
	var result int
	result = x * x * x
	return result
}

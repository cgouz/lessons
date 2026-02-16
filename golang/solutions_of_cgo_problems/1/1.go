package main

import "fmt"

func main () {

	var a, b int

	fmt.Print("Enter side A of the right quadrilateral: ")
	fmt.Scan(&a)

	fmt.Print("Enter side B of the right quadrilateral: ")
	fmt.Scan(&b)

	s := a * b
	p := 2 * (a + b)

	fmt.Println("surface area of ​​a right triangle:", s)
	fmt.Println("perimeter of a right triangle:", p)
}
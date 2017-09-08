package main

import (
	"fmt"
	"math"
)

const PI  = 3.14

func add(x int, y int) float32 {
	return  float32(x) + float32(y)
}

/**
declare type at the end
 */
func subtract(x, y int)  int {
	return x - y
}

func swapString(string1, string2 string) (string, string) {
	return string2, string1
}

func split(sum int) (x, y int)  {
	x = sum * 4 / 9
	y = sum - x

	return
}

var c, python, java bool

func main()  {
	fmt.Println("Let check some square roots: ", math.Pi)
	fmt.Println("Sum of 3 + 4 =",add(3,4))
	fmt.Println("Substract 4 - 3 =",subtract(4,3))

	p, q := swapString("Sohel", "Rana")
	fmt.Println("Swaping string \"Sohel\" and \"Rana\"", p, q)

	a, b := split(15)
	fmt.Println(a,b)

	fmt.Printf("Let know about some programming languages java %v, c %v, python %v", c, python, java)

	fmt.Println("The valu of PI : ", PI)

	//try to change PI
	//PI = 3; it will generate an error
}
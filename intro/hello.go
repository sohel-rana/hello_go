package main

import (
	"fmt"
	"math"
	"runtime"
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
	defer fmt.Println("Splited the supplied number. Returning from the function")
	x = sum * 4 / 9
	y = sum - x

	return
}

func sqrt(x float64) string  {
	if x<0{
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

var c, python, java bool

func main()  {

	defer fmt.Println("This will be printed when main function returns")

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

	//calculate sum using for loop
	sum := 0
	for i:=0; i<=10; i++{
		sum = sum+i
	}
	fmt.Println(sum)

	//for loop without init
	sum2 := 1
	for ; sum2<1000; {
		sum2 +=sum2
	}
	fmt.Println("Sum without init: ", sum2)

	//for as while
	sum3 := 1
	for sum3<1000 {
		sum3 += sum3
	}

	fmt.Println("Sum as while loop: ", sum3)

	fmt.Println(sqrt(45))
	fmt.Println(sqrt(-5))


	switch os:= runtime.GOOS;  os{
		case "linux":
			fmt.Printf("Running on Linux: %s", os)
		case "darwin":
			fmt.Printf("%s.", os)
			fmt.Printf("Running on Mac %s ", os)
		default:
			fmt.Println("OS could not detect")

	}
	fmt.Println("The sqrt of 10 is : ", mySqrt(10))

}
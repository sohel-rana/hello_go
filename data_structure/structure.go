package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func printSlices(s[] int)  {
	fmt.Printf("length=%d, capacity=%d, data=%v\n", len(s), cap(s), s)
}

func main()  {
	//pointers example
	fmt.Println("------- Pointer Example ------")
	i, j := 42, 81
	p := &i
	*p = 10

	fmt.Println("Value of i is: ", i)
	fmt.Println("Value of p is: ", *p)

	p = &j
	*p = *p/10

	fmt.Println("Value of j is:", j)
	fmt.Println("Value of p is:", j)
	fmt.Println("---- Pointer Example Ends ----")
	fmt.Println("")

	//structure examples
	fmt.Println("----- Structer Example -----")
	v := Vertex{1,3}
	fmt.Println(v)
	v.X = 5
	fmt.Println("changing the value of V.X to 5 :", v)
	fmt.Println("----Structer Example Ends----")
	fmt.Println()


	//array examples
	var a[1] int
	a[0] = 1

	var words[2] string
	words[0] = "Hello"
	words[1] = "World"
	fmt.Println("----Array/slice Example----")

	fmt.Println("Whole array: ", a, words)

	primes := [8]int {1,3,5,7,11,13, 17, 19}
	fmt.Println("Primes array", primes)

	//let make slices :D
	var s[] int = primes[1:4]
	fmt.Println("Print slice of primes[1:4]: ", s)

	//slice literal
	q := []int {1, 1, 2, 3, 5, 8, 13}
	fmt.Println("Print fmt slice literal []int{} : ", q)


	sliceWithStruct := []struct{
		x int
		status bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println("1st element of slice with struct: ", sliceWithStruct[0])
	fmt.Println("slice with struct: ", sliceWithStruct)

	fmt.Println("Printing slices with size and capacity")
	printSlices(q)

	makeVar := make([]int, 5)
	fmt.Println("Created arry with make makeVar: ", makeVar)

	makeVar2 := make([]int, 2, 5);
	fmt.Println("initialzed with 2, created with make(): ", makeVar2)

}

package main

import (
	"fmt"
)

func SumInt(x, y int) int {
	return x + y
}

func SumInt32(x, y int32) int32 {
	return x + y
}

func SumInt64(x, y int64) int64 {
	return x + y
}

func SumFloat32(x, y float32) float32 {
	return x + y
}

func SumFloat64(x, y float64) float64 {
	return x + y
}

func SumNumber[T int | int8 | int16 | int32 | int64 | float32 | float64 | string](x T, y T) T {
	return x + y
}

// there is a better approach to the above
// using a type constraint
type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}

func SumImproved[T Numeric](x, y T) T {
	return x + y
}

func main() {

	fmt.Println("Hello generics")

	// Non-generic functions
	fmt.Printf("Calling Int functions with %v and %v: %v, %v and %v\n",
		3, 5, SumInt(3, 5), SumInt32(3, 5), SumInt64(3, 5))

	fmt.Printf("Calling Float functions with %v and %v: %v and %v\n",
		3.5, 5.5, SumFloat32(3.5, 5.5), SumFloat64(3.5, 5.5))

	// Generic functions
	fmt.Printf("Calling Generic function with %v and %v: %v and %v\n",
		3.5, 5.5, SumNumber(3.5, 5.5), SumNumber(3.5, 5.5))

	fmt.Printf("Calling Generic function with %v and %v: %v and %v\n",
		3, 5, SumNumber(3, 5), SumNumber(3, 5))

	fmt.Printf("Calling Generic function with %v and %v: %v \n",
		"John ", "Doe", SumNumber("John ", "Doe"))

	// Generic functions Improved
	fmt.Printf("Calling Improved Generic function with %v and %v: %v and %v\n",
		3.5, 5.5, SumImproved(3.5, 5.5), SumImproved(3.5, 5.5))

	fmt.Printf("Calling Improved Generic function with %v and %v: %v and %v\n",
		3, 5, SumImproved(3, 5), SumImproved(3, 5))

	fmt.Printf("Calling Improved Generic function with %v and %v: %v \n",
		"John ", "Doe", SumImproved("John ", "Doe"))

}

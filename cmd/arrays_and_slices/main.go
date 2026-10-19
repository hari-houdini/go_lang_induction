package main

import "fmt"

func main() {
	/*
		Arrays
	*/
	var int32Array [5]int32
	fmt.Println("Array of int32:", int32Array[0])
	// Prints the first two elements of the array using slicing
	fmt.Println("Array of int32:", int32Array[1:3])

	// Fixed length
	// Same type
	// Indexable
	// Contiguous memory allocation
	fmt.Println("Address of int32Array:", &int32Array[0], &int32Array[1], &int32Array[2], &int32Array[3], &int32Array[4])

	// Initializing an array with values
	var int32Array2 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Println("Array of int32:", int32Array2)

	/*
		Slices
	*/
	// Slices are just a wrapper around arrays.
	// In other words, slices are fancier arrays with extra functions.
	var int32Slice []int32 = []int32{1, 2, 3, 4, 5}
	fmt.Println("Slice of int32:", int32Slice)
	fmt.Printf("The length of the slice: %v with capacity of %v\n", len(int32Slice), cap(int32Slice))

	// Slices are dynamic in size, so you can append elements to them.
	int32Slice = append(int32Slice, 6)
	fmt.Println("Slice of int32 after appending:", int32Slice)
	// Length of the array would be 6, but capacity would be 12
	// since after append, it everything would be copied to a new array with double the capacity.
	// Although you cannot access the other indices outside the length but within the capacity
	fmt.Printf("The length of the slice: %v with capacity of %v\n", len(int32Slice), cap(int32Slice))

	int32Slice2 := []int32{5, 10}
	// Appending using spread operator
	int32Slice = append(int32Slice2, int32Slice...)
	fmt.Println("Slice of int32 after appending:", int32Slice)
	fmt.Printf("The length of the slice: %v with capacity of %v\n", len(int32Slice2), cap(int32Slice2))

	// Another way to declare a slice is by using the make function.
	// len is 5 and cap is 10
	int32Slice3 := make([]int32, 5, 10)
	fmt.Println("Slice of int32:", int32Slice3)
	fmt.Printf("The length of the slice: %v with capacity of %v\n", len(int32Slice3), cap(int32Slice3))
}

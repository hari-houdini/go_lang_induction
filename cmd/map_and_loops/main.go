package main

import "fmt"

func main() {
	/*
		Maps
	*/
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{
		"Ada":   25,
		"Sarah": 30,
	}
	fmt.Println(myMap2["Ada"])
	fmt.Println(myMap2["Sarah"])
	// If I acces something that doesn't exist, it will return the default value of the type, which is 0 for uint8
	fmt.Println(myMap2["John"])

	// Flipside you could do it like this
	var age, ok = myMap2["John"]
	if ok {
		fmt.Println("John's age is", age)
	} else {
		fmt.Println("John's age is not found")
	}

	// to delete a key from a map, you can use the built-in delete function
	delete(myMap2, "Ada")
	fmt.Println(myMap2)

	/*
		Loops
	*/
	for name := range myMap2 {
		fmt.Println("Name:", name, "Age:", myMap2[name])
	}
	for name, age := range myMap2 {
		fmt.Println("Name:", name, "Age:", age)
	}

	var int32Array [5]int32 = [5]int32{1, 2, 3, 4, 5}
	for i, v := range int32Array {
		fmt.Println("Index:", i, "Value:", v)
	}

	// there is no while loop in Go, but you can use for loop to achieve the same thing
	var i int32 = 0
	for i < 5 {
		fmt.Println("Index1:", i)
		i++
	}

	// or this is the loop equivalent of rust
	for {
		if i >= 10 {
			break
		}
		fmt.Println("Index2:", i)
		i++
	}

	// regular for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Index3:", i)
	}

}

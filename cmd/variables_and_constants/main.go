// A Collection of go files are whats called a package.
// A package is a way to group related Go source files together.
// Each Go file in a package must start with the same package declaration, which defines the package name.
// The main package is a special package that serves as the entry point for a Go program.
// When you run a Go program, the Go runtime looks for the main function in the main package to start execution.

// Similarly a collection of packages can be organized into a module.
// A module is a collection of related Go packages that are versioned together as a single unit.
// Modules are defined by a go.mod file, which specifies the module's path and its dependencies.

// So we define the package a Go file belongs to using the package keyword followed by the package name.
// Has to be the same for all files in the same folder.
// The main package is a special package that serves as the entry point for a Go program.
package main

// Go doesnt let you import packages that you dont use, so if you import a package and dont use it,
// the compiler will throw an error.
import (
	"fmt"
	"unicode/utf8"
) // string length function

func main() {
	fmt.Println("Hello World!")

	/*
		Variables
	*/
	// Similarly, any variable declared in a Go file must be used, otherwise the compiler will throw an error.
	// int, int8, int16, int32, int64
	// just int is an alias for int32 or int64 depending on the architecture of the machine.
	var intNum int16 = 32767 // 32767 is the maximum value for int16

	// intNum = 32768 will throw an error at compile time because it exceeds the maximum value for int16.
	// but will not throw an error at runtime because Go is a statically typed language and the type of intNum is already defined as int16.
	// intNum = intNum + 1
	fmt.Println("intNum:", intNum)

	// uint, uint8, uint16, uint32, uint64
	var uIntNum uint16 = 65535 // 65535 is the maximum value for uint16

	fmt.Println("uIntNum:", uIntNum)

	// float32, float64 (no float type, just float32 and float64)
	var floatNum float32 = 12345678.9  // float32 is a 32-bit floating point number
	fmt.Println("floatNum:", floatNum) // this will print 12345679.00000, so decide which one to use based on the precision you need.

	var int1, int2 int = 10, 20 // multiple variable declaration
	fmt.Println("int1:", int1, "int2:", int2)
	/*
		Arithmetic Operations
	*/
	// Go supports the following arithmetic operations: +, -, *, /, %
	// But cannot perform arithmetic operations on different types, so you cannot add an int and a float, for example.
	var result float32 = floatNum + float32(intNum) // need to convert intNum to float32 before adding
	fmt.Println("result:", result)

	/*
		Strings
	*/
	var str string = "Hello, Go!"
	fmt.Println("str:", str)

	// Length of a string using len() function will return the number of bytes in the string,
	// not the number of characters.
	// ASCII characters are 1 byte each
	fmt.Println("Length of str:", len("A"))                    // prints 1
	fmt.Println("Length of str:", len("😊"))                    // prints 4 because the emoji is 4 bytes
	fmt.Println("Length of str:", utf8.RuneCountInString("😊")) // prints 1 because utf8.RuneCountInString counts the number of Unicode code points

	var myRune rune = 'A'          // rune is an alias for char, and represents a Unicode code point
	fmt.Println("myRune:", myRune) // prints 65, which is the Unicode code point for 'A'

	var str2 = "Hello, 世界" // String type is inferred
	fmt.Println("Length of str2:", len(str2))

	str3 := "Hello, 世界" // Short variable declaration, type is inferred
	fmt.Println("Length of str3:", len(str3))

	/*
		Defaults
	*/
	// Go has default values for variables that are not initialized.
	// The default value for int, uint, rune is 0, for float is 0.0, for string is "", and for bool is false.
	var defaultInt int
	fmt.Println("defaultInt:", defaultInt) // prints 0

	/*
		Constants
	*/
	// Constants are immutable values that are known at compile time and do not change for the life of the program.
	// Constants should always be initialized at the time of declaration, and cannot be declared using the := syntax.
	const pi float64 = 3.14159 // constant declaration
	fmt.Println("pi:", pi)
}

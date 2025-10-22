/* multiline comment */
// Single line comment

package main // Package declaration

import "fmt" // Importing the fmt package for formatted I/O

// Function main - entry point of the program
func main() {

	//01_statements and expressions
	fmt.Println("Hello Developers") // Print statement

	//02_variables and data types (Explict method)
	var age int = 30
	var name string = "alice"
	var isStudent bool = true
	var height float64 = 5.7
	var decimalNumber float32 = 3.14
	fmt.Println("\nName:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Height:", height)
	fmt.Println("Decimal Number:", decimalNumber)

	//03_constants (Explict method)
	const pi float64 = 3.14159
	const appName string = "GoApp"
	fmt.Println("\nApp Name:", appName)
	fmt.Println("Value of Pi:", pi)

	//04_short variable declaration
	city := string("New York")
	temperature := float32(72.5)
	isRaining := bool(false)
	fmt.Println("\nCity:", city)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Is Raining:", isRaining)

	//05_Arrays

	// Fixed array size
	var farr = [5]int{}
	farr[0] = 10
	farr[1] = 20
	farr[2] = 30
	farr[3] = 40
	farr[4] = 50
	fmt.Println("\nArray elements:", farr)
	//array with initial values
	var iarr = [3]string{"Go", "Python", "Java"}
	fmt.Println("Programming Languages:", iarr)
	//array with inferred size
	var carr = [...]int{1, 2, 3, 4, 5}
	fmt.Println("inferred array:", carr)

	//Accessing array elements
	fmt.Println("\nFirst element of farr:", farr[0])
	// Access last element using len(farr)-1
	fmt.Println("Last element:", farr[len(farr)-1]) // No negative indexing in Go
	// Modifying array elements
	fmt.Println("\nOriginal farr:", farr)
	farr[2] = 35
	fmt.Println("Modified farr:", farr)

	//Iterating Over an array
	for i, v := range farr {
		fmt.Println("\nElement at index", i, "is", v)
	}

	//Array methods & properties
	fmt.Println("\nLength of iarr:", len(iarr))

	// Comparing arrays
	fmt.Println("Are farr and carr equal?", farr == carr) // false

	//Multi-dimensional arrays (type 1)
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	fmt.Println("\n3D Array (Matrix) Type 1:", matrix)

	//type 2
	var mat = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("3D Array (Matrix) Type 2:", mat)

	//length of array
	fmt.Println("\nLength of matrix:", len(matrix))

	//06 Creating a slice
	//using a literal
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//using Make
	slc := make([]int, 5)
	slc[0] = 10
	slc[1] = 20
	slc[2] = 30
	slc[3] = 40
	slc[4] = 50

	fmt.Println("\nslice:", slice)
	fmt.Println("slc:", slc)

	//slicing
	subset := slc[0:2]
	fmt.Println("Sliced Elements:", subset)

	//Length & Capacity of a slice
	fmt.Println("Length: ",len(slc))
	fmt.Println("Capacity: ",cap(slc))

}

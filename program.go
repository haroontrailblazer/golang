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

}

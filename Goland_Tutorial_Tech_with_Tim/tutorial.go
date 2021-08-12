package main

import "fmt"

// Creating structures
type student struct {
    name   string
    branch string
}

type teacher struct {
    language string
    marks    int
}

// Same name methods, but with
// different type of receivers
func (s student) show() {

    fmt.Println("Name of the Student:", s.name)
    fmt.Println("Branch: ", s.branch)
}

func (t teacher) show() {

    fmt.Println("Language:", t.language)
    fmt.Println("Student Marks: ", t.marks)
}

func methodOverloading() {

    // Initializing values
    // of the structures
    val1 := student{"Rohit", "EEE"}

    val2 := teacher{"Java", 50}

    // Calling the methods
    val1.show()
    val2.show()
}

type value_1 string
type value_2 int

// Creating same name function with
// different types of non-struct receivers
func (a value_1) display() value_1 {

    return a + "forGeeks"
}

func (p value_2) display() value_2 {

    return p + 298
}

func functionOverloading() {

    // Initializing the values
    res1 := value_1("Geeks")
    res2 := value_2(234)

    // Display the results
    fmt.Println("Result 1: ", res1.display())
    fmt.Println("Result 2: ", res2.display())
}

func main() {
	variableDeclarations()
	functionOverloading()
	methodOverloading()
}

func variableDeclarations() {
	// explicit declaration
	var timName string = "Tim"
	var number uint8 = 255
	// implicit declaration
	myBool := true // walrus operator
	// if  (no ternary operator)
	if myBool {
		fmt.Printf("Hello, %s! The number is: %d\n", timName, number)
	}
	// increment operator
	number++
	fmt.Println("The number is now:", number)
	// viewing types
	fmt.Printf("myBool's type is: %T\n", myBool)
	fmt.Printf("timName's type is: %T\n", timName)
	fmt.Printf("number's type is: %T\n", number)

	// Lesson #4: https://www.youtube.com/watch?v=GQ880MlHBBE

	// for i := range  {
	// 	fmt.Printf("%d, ", i)
	// }

}

// Declaration
package main

import "fmt"

// package level declaration
const boilingF =212.0
func main(){
	// local level declaration
	var f = boilingF
	var c = (f-32)*5/9
	fmt.Printf("boiling point = %g F or %g C\n", f,c)

	// var name type = expression
	// if no initializer then int = 0, string ="",boolean = false
	var myName string = "Clement"
	fmt.Println(myName)

	// This is a short declaration
	i := 100
	fmt.Println(i)

	// Pointers: A pointer value is the address of a variable ,&x is (address of x)
	x := 1
	p := &x //p , of type int , points to x
	fmt.Println(p , &x) // p = 1
	*p = 2
	fmt.Println(x) // 2

	// The new function : new(T) create unnamed variable of type T, init to 0 and return address which is the type of T

	n := new(int) // n, of type int , points to an unnamed int variable
	fmt.Println(*n)
	*n = 2 //set unnamed int to 2
	fmt.Println(*n)


	

}


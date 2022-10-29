// Package declearation always first, it's a convention to always use package main
package main

import "log"

// Every go package must consist of at least on function i.e main function
// Functions are declared with "func" keyword
// Unitialized int is asigned to 0
// _ means ignore that value

func main(){
	// Variable declaration
	
	var whatToSay string
	whatToSay,_= saySomething("Hello world")

	log.Println(whatToSay)

}


func saySomething(s string) (string,string){
	return s, "Goodbye"
}
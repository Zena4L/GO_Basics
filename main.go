// Package declearation always first, it's a convention to always use package main
package main

import "log"

// Every go package must consist of at least on function i.e main function
func main(){
	// Variable declaration
	var whatToSay string= saySomething("Hello world")

	log.Println(whatToSay)

}


func saySomething(s string) string{
	return s
}
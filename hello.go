package main

import (
	"fmt"

	// "learn.com/greetings"
	// "learn.com/loops"
	"learn.com/ErrorHandling"
	"learn.com/Array"
)

func main(){
//    message, err := greetings.Hello("Aditya")
//    if err != nil {
// 	fmt.Println("An error occurred:", err)
//    } else {
// 	fmt.Println(message)
//    }

//** for loop

    // number := loops.For()
	// fmt.Println(number)

// ** returning error

	number, err := ErrorHandling.Divide(64.2, 2.0)
	   if err != nil {
		fmt.Println("An error occurred:", err)
	   } else {
		fmt.Println(number)
}

// ** code will panic

// fmt.Println("Starting program")
// ErrorHandling.CausePainc()
// fmt.Println("Program ended normally")

//** recovery from panic

fmt.Println("Starting program")
    ErrorHandling.PanicRecovery()
    fmt.Println("Program ended normally")
 
	Array.Array()
}
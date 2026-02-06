package main

import (
    "fmt"
    "log"
    "example.com/greet"
)

func main() {
	//loging errors to cmd line
	//error format
	log.SetPrefix("Error: ")
	//disable timstamp and extra details
	log.SetFlags(0)
    // Get a greeting message and print it.
    names :=[]string {"Rashmith","Thara","Ravi"}

    message,err := greet.MutiName(names)

    if err!=nil{
	    log.Fatal(err)
    }
    fmt.Println(message)
}

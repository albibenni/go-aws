package main

import "fmt"

func main () {

    var myName string = "Benni"
    fmt.Printf("hello %s\n", myName)
    inferName:= "benni inferred"
    myInt:= 10
    myFloat:=10.0
    
    fmt.Printf("Hello %s my int is %d my float is %f\n", inferName, myInt, myFloat)
}


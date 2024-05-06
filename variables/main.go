package main

import "fmt"

func main () {

    var myName string = "Benni"
    fmt.Printf("hello %s\n", myName)
    inferName:= "benni inferred"
    myInt:= 10
    myFloat:=10.0
    
    fmt.Printf("Hello %s my int is %d my float is %f\n", inferName, myInt, myFloat)

    var anotherName string
    var aBool bool
    var otherInt int

    fmt.Printf("they are not declared: zero value then: %s %t %d", anotherName, aBool, otherInt)
}


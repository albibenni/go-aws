package main

import (
	"fmt"
	"slices"
)

func main() {
	animals := [2]string{}

	animals[0] = "dog"
	animals[1] = "cat"
	// invalid    animals[2] = "cat2"
    
    fmt.Println(animals)

    //slice flex dynamic array
    animalsSlice:= []string{
        "dog",
        "cat",
    }
    animalsSlice = append(animalsSlice, "moose")
    
    fmt.Println(animalsSlice)

    animalsSlice = slices.Delete(animalsSlice, 0, 1) //goes for range deletion
    fmt.Println(animalsSlice)

    for i:=0; i<len(animalsSlice); i++ {
        fmt.Printf("animal %s\n", animalsSlice[i])
    }
    for i:= 0; i < 10; i++ {
       fmt.Println(i) 
    }

    for index, value := range animalsSlice {
        fmt.Printf("index %d and animal is %s\n", index, value)
    }
    for value := range 12{
        fmt.Printf("%d\n", value)
    }
}

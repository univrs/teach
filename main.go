package main

import "fmt"

func main() {

        // foo := &myStruct{"bar"}
        foo := newMyStruct()

	    // foo := myStruct{"bar"}
        foo.myField = "foo"
        
        foo.myMap["bar"] = "baz"

        fmt.Println(foo.myField)
        fmt.Println(foo.myMap)
}

type myStruct struct {
     myMap map[string]string
     myField string
}

func newMyStruct() *myStruct {
    
    result := myStruct{}
    result.myMap = map[string]string{}
    
    return &result
}
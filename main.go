package main

import "fmt"

func main() {

        // foo := &myStruct{"bar"}
        foo := newMyStruct()

	    // foo := myStruct{"bar"}
        foo.myField = "foo"
        
        foo.myMap["bar"] = "baz"

        fmt.Println(foo.myField)  //  prints foo
        fmt.Println(foo.myMap["bar"])  // print baz
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
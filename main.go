package main

import "fmt"

func main() {

        // foo := &myStruct{"bar"}
        foo := newMyStruct()

	    // foo := myStruct{"bar"}
        foo.myField = "foo"
        
        foo.myMap["bar"] = "baz"

        fmt.Println(foo.myField)  //  prints foo
        fmt.Println(foo.myMap["bar"])  // prints baz

        mp := messagePrinter{"printed by printMessage method"}
        
        emp := composeMessagePrinter{ messagePrinter{"printed by printMessage composed method"}} 
        emp.printMessage() 
        mp.printMessage()  // method use, printed by...
    
}


type messagePrinter struct {
    message string
}

func (mp *messagePrinter) printMessage() {
    
    fmt.Println(mp.message)
}

type composeMessagePrinter struct {
    messagePrinter 
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
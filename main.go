package main

import "fmt"

func main() {
	foo := myStruct{"bar"}
        // foo.myField = "bar"

     fmt.Println(foo.myField)
}

type myStruct struct {
     myField string
}


package main 

import (
    "fmt"
)
var above string = "above the main func"
func main() {
    var a = "initial"
    fmt.Println(a)

    var b, c int = 1,2
    fmt.Println(b + c)

    var d = 1
    fmt.Println(d)

    f := "apple"
    fmt.Println(f)

    f = "sravya"
    fmt.Println(f)

    fmt.Println(above)

    // f = 12 throws an error
    // fmt.Println(f)
}
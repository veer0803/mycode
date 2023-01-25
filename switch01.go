/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
//    "runtime"
)

func main() {

           // init gover                
    switch gover := "go1.16" ; gover {
    case "go1.19":  // if matched "go1.19", do the code below and then STOP
        fmt.Println("Released in 2022")
    case "go1.18":                 // if matches "go1.18", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case "go1.17":
        fmt.Println("This version of Go is fine")
    case "go1.16", "go1.16.5", "go1.15":       // if matches "go1.16", "go1.16.5", OR "go1.15" 
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}


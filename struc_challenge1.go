

/* Challenge 1 */

package main

import "fmt"


type Astro1 struct {
    name string
    age  int
    mission string
    isneeded bool

 }

 func main () {

  astr1 := Astro1{"Veer	 " ,35,"cir", true}
  astr2 := Astro1{"bha ",23, "cir" ,false}
  astr3 := Astro1{"sud" ,32,"space" ,true }


  fmt.Println(astr1)
  fmt.Println(astr2)
  fmt.Println(astr3)

 }

package main

import (
    "fmt"
    "reflect"
)

func main() {
   a := "abc"
   b := a[1:]
   fmt.Printf("%+v\n", b)
   fmt.Println(reflect.TypeOf(b))

   c := a[1:2]
   fmt.Println(c)
}

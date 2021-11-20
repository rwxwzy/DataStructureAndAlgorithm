package main

import (
    "fmt"
    "reflect"
)

func main() {
   a := "abc"
   fmt.Printf("%c\n", a[1])
   b := a[1:]
   fmt.Printf("%+v\n", b)
   fmt.Println(reflect.TypeOf(b))

   c := a[1:2]
   fmt.Println(c)

   d :=[]byte(a)
   fmt.Println(reflect.TypeOf(d))

   fmt.Println(a[1])
   fmt.Println("a 0: 0")
   fmt.Println(a[:0])
}

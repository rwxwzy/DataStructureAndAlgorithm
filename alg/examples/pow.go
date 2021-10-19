package main

import (
    "fmt"
)

func pow(value float64, n int) float64 {

    // 1. normal case

    return 0
}

func powRecursion(value float64, n int) float64 {

    // 2.
    if n == 0 {
        return 1
    }
    if n == 1 {
        return value
    }

    // 1. normal case
    /*
        1.1 n is odd
           value^n = (value ^ 2) ^ (n/2) * value
        1.2 n is even
           value ^ n = (value ^ 2) ^ (n/2)
    */
    if n % 2 != 0 { // odd
        return powRecursion(value * value, n / 2) * value // 每一层调用的时候将value等压栈，然后调用返回的时候将value等弹出
    } else { // even
        return powRecursion(value * value, n / 2)
    }
}

func powLoop(value float64, n int) float64 {// wrong
    if n == 0 {
        return 1
    }
    if n == 1 {
        return value
    }

    var product float64 = value
    for i := 1; i < n ; i = i * 2 {
        product = product * product
    }
    
    if n % 2 != 0 {
        return product
    } else {
        return product * value
    }
}

func main() {
    fmt.Println(powRecursion(3, 4))
    fmt.Println(powLoop(3, 5))
}
package main

import (
    "fmt"
)

func BinarySearch(a []int, value int) int{
    left := 0
    right := len(a) - 1
    for ;left <= right; {
        middle := left + (right - left) /2
        if a[middle] == value {
            return middle
        } else if a[middle] > value {
            right = middle - 1
        } else if a[middle] < value {
            left = middle + 1
        }
    }

    return -1
}

func main() {
    a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    b := BinarySearch(a, 3)
    fmt.Println(b)
    b = BinarySearch(a, 11)
    fmt.Println(b)

}

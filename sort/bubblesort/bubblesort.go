package main

import (
    "fmt"
)

func bubbleSort(data []int) {
    len := len(data)
    for i := 0; i < len; i ++ {
        for j := i; j < len - 1; j ++ {
            if data[j] < data[j + 1] {
                tmp := data[j]
                data[j] = data[j + 1]
                data[j + 1] = tmp
            }
        }
    }
}

func main() {
    data := []int{9, 8, 7, 6, 6, 4, 5, 3, 3, 1, 2, 0}
    bubbleSort(data)
    for _, v := range data {
        fmt.Println(v)
    }
    
    data = []int{}
    bubbleSort(data)
    fmt.Println("Print empty")
    for _, v := range data {
        fmt.Println(v)
    }
    fmt.Println("end print empty")
}
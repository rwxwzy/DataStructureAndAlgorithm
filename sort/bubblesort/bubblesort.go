package main

import (
    "fmt"
)

func bubbleSort(data []int) {
    length := len(data)
    for i := 0; i < length; i ++ {
    // 每次循环把最小或最大的移动到最右边，所以只需要排列前面的就可以
    // 外循环i表示排列好了几个，所以j只需要排列到length - i - 1
    // 因为下面交换的时候是j + 1可以达到最后一个，所以需要length -i -1
        for j := 0; j < length - i - 1 ; j ++ {
            if data[j] < data[j + 1] {
                tmp := data[j]
                data[j] = data[j + 1]
                data[j + 1] = tmp
            }
        }
    }
}

func main() {
    data := []int{9, 8, 7, 6, 6, 4, 5, 3, 3, 1, 0, 2}
    bubbleSort(data)
    for _, v := range data {
        fmt.Println(v)
    }
}
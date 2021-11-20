package main
import (
    "fmt"
)

/*
（5）给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返
回该子数组的长度。
示例 1:
输入: nums = [0,1]
输出: 2
说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组」

*/


func GetLongestSubArray(array []int)int {

    counterMap := make(map[int]int) // counter: prefix
    
    //为了获取前两个为1, 0(1, -1)或者是0, 1(-1, 1)的情况
    counterMap[0] = -1

    counter := 0
    maxLength := 0
    for i := 0; i < len(array); i ++ {
        if array[i] == 1 {
            counter ++
        } else {
            counter --
        }

        if prefix, ok := counterMap[counter]; ok { //相同的counter，说明中间的0, 1数相等
            length := i - prefix
            if length > maxLength {
                maxLength = length

            }
        } else {
            counterMap[counter] = i
        }
    }

    return maxLength
}

func main() {
    array := []int{0, 1}
    length := GetLongestSubArray(array)
    fmt.Println(length)
}






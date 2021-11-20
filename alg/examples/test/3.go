package main

import (
    "fmt"
)

/*
（3）给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标
值 target 的那 两个 整数，并返回它们的数组下标。
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

遍历: 第一个和后续的和
            第二个和后续的和
*/

func GetSumIndex(array []int, sum int) (int, int) {
    for i := 0; i < len(array); i ++ {
        for j := i + 1; j < len(array); j ++ {
            if array[i] + array[j] == sum {
                return i, j
            }
        }
    }
    return -1, -1
}


func GetSumIndexMap(array []int, sum int) (int, int) {
    var sets = make(map[int]int) // map[value]index
    for i := 0; i < len(array); i ++ {
        value := sum - array[i]
        if index, ok := sets[value]; ok {
            return i, index
        }
        // 不需要单独的一遍存储，因为从头开始的时候，头 + 尾 
        // 虽然没有取到值，但是当查找到尾部的时候，头(前面那个值已经存储了)
        //单独的循环用来存储到map, 会导致两遍
        sets[array[i]] = i
    }
    
    
    return -1, -1
}


func main() {
    nums :=[]int{2, 0,1, 7,11,15}
    target := 9
    i, j := GetSumIndex(nums, target)
    fmt.Printf("%d, %d\n", i, j)
    

    i, j = GetSumIndexMap(nums, target)
    fmt.Printf("%d, %d\n", i, j)
}
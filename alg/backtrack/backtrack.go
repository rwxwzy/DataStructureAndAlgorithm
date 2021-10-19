package main
import (
    "fmt"
)

/*
回溯算法的这段核心框架
for 选择 in 选择列表:
    # 做选择
    将该选择从选择列表移除
    路径.add(选择)
    backtrack(路径, 选择列表)
    # 撤销选择
    路径.remove(选择)
    将该选择再加入选择列表
    
只要在递归之前做出选择，在递归之后撤销刚才的选择


*/

// full permutation 全排列 n!
func fullPermutationWrapper() {
    var SliceList [][]int = make([][]int, 0)
    var track []int
    fullPermutation([]int{1, 2, 3}, SliceList, track) 
    
    for _, list := range SliceList {
        for _, v := range list {
            fmt.Println(v)
        }
    }
    
}
func fullPermutation(nums []int, lists [][]int, track []int) {

    if len(track) == len(nums) {
        lists = append(lists, track)
        return
    }

    if len(track) == 0 {
        track = make([]int, 0)
    }
    
    for _, numsValue := range nums {
        if len(track) == 0 {
            track = append(track, numsValue)
            fullPermutation(nums, lists, track)
        }
        for _, listValue := range track {
            if numsValue == listValue {
                continue
            }
            fmt.Printf("value: %d\n", numsValue)
            track = append(track, numsValue)
            fullPermutation(nums, lists, track)
            
        }
        track = track[ : len(track) - 1 ]
    }
}

func main() {
    fullPermutationWrapper()
}
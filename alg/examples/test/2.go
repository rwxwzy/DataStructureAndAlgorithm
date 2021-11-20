package main
import (
    "fmt"
)


/* 前缀就是从头开始的
其实就是所有字符串都有的子串，并且这个子串的位置还比较特殊，它在字符串的开始位置。
编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字
符串 ""。
示例 1：
输入：strs = ["flower","flow","flight"] 输出："fl"
*/

func longestPrefix(array []string) string{
    minLength := len(array[0])
    var minStr string
    for _, v := range array {
        tmp := len(v)
        if minLength > tmp {
            minLength = tmp
            minStr = v
        }
    }

    fmt.Printf("len: %d, minStr: %s\n", minLength, minStr)

    pos := 0
    isEqual := true
    for i := 0; i < minLength; i ++ {
        for _, v := range array {
            if minStr[i] != v[i] {
                isEqual = false
                break
            }
        }
        if isEqual == false {
            break
        }
        fmt.Printf("pos : %d\n", pos)
        pos ++
    }

    if pos == 0 {
        return ""
    }

    return minStr[:pos]
}


// 水平扫描法
func levelLongestPrefix(array []string) string{
    // 0, 1
    var result string = array[0] ////very important!!!!!!!!!!, 提前取一个
    // result 后续做减法
    for i := 1; i < len(array); i ++ {
        for j := 0; j < len(result) && j < len(array[i]); j ++ {
            if result[j] != array[i][j] {
                result = result[:j]
                break
            }
        }
    }
    
    return result
}









func main() {
    strs := []string{"fleower","fleow","fleight"}
    str := longestPrefix(strs)
    fmt.Println(str)
    
    
    fmt.Println("level")
    str = levelLongestPrefix(strs)
    fmt.Println(str)
    
}
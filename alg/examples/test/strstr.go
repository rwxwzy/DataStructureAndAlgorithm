package main
import (
    "fmt"
)

/*
（ 1）实现 strStr() 函数。
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle
字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
示例 1：By_dzz

输入：haystack = "hello", needle = "ll"
输出：2
*/

func strStr(haystack, needle string) int {
    for i := 0; i <= len(haystack) - len(needle); i ++ {
        if haystack[i: i + len(needle)] == needle {
            return i
        }
    }

    return -1
}

func strStr1(haystack, needle string) int {
    for i := 0; i <= len(haystack) - len(needle); i ++ {

        // 新函数中，比较haystack[i: i + len(needle)] 和needle的值
        for j := 0; j < len(needle); j ++ {
            if haystack[i + j] != needle[j] {
                break
            } else {
                if j == len(needle) - 1 {
                    return i
                }
            }
        }
        
        
    }

    return -1
}

func main() {
    var pos int

    pos = strStr("hello", "ll")
    fmt.Println(pos)

    pos = strStr("hello", "hel")
    fmt.Println(pos)

    pos = strStr("hello", "hl")
    fmt.Println(pos)
    
    // 
    fmt.Print("----------------------\n")
    pos = strStr1("hello", "ll")
    fmt.Println(pos)

    pos = strStr1("hello", "hel")
    fmt.Println(pos)

    pos = strStr1("hello", "hl")
    fmt.Println(pos)
}
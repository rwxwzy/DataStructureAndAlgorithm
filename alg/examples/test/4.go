package main
import (
    "fmt"
    "strings"
)

/*

（4）给你一个字符串 s ，逐个翻转字符串中的所有单词 。单词是由非空格字符组成的字
符串。s 中使用至少一个空格将字符串中的单词分隔开。请你返回一个翻转 s 中单词顺序
并用单个空格相连的字符串。
示例 1：
输入：s = "the  sky  is blue "
输出："blue is sky the"
*/
func RevertStrLib(input string) string {
    inputSlice := strings.Fields(input) // 不能用split, 多个空格，会保留空格
    var ret string
    for i := len(inputSlice) - 1; i >= 0; i -- {
        ret = ret + " " + strings.Trim(inputSlice[i], " ")
    }
    ret = ret[1:]
    return ret
}

func RevertStr(input string) string {
    commonInputStr := " " + input + " "

    var ret string
    start, end := -1, -1
    for i := len(commonInputStr) - 1; i >= 0; i -- {
        if commonInputStr[i] != ' ' {
            if end == -1 {
                end = i + 1
            }
        } else {// byte == ' '
            if end != -1 {
                start = i
                ret = ret + " " + string(commonInputStr[start + 1: end])
                start = -1
                end = -1
            }
        }
    }
    
    ret = ret[1:]
    return ret
}

func main() {
    var input = "the  sky  is blue"
    ret := RevertStrLib(input)
    fmt.Println(ret)
    
    ret = RevertStr(input)
    fmt.Println(ret) 
}
















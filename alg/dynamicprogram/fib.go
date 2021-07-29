package main

import (
    "fmt"
)

func fibRecursive(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    
    return fibRecursive(n - 2) + fibRecursive(n - 1)
}

func fibMemo(n int) int {
    recurMemo := make([]int, n + 1)
    ret := fibRecurMemo(n, recurMemo)
    return ret
}
func fibRecurMemo(n int, recurMemo []int) int {    
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    
    if recurMemo[n] != 0 {
        return recurMemo[n]
    }
    recurMemo[n] = fibRecurMemo(n - 2, recurMemo) + fibRecurMemo(n - 1, recurMemo)
    return recurMemo[n]
}

func fibDP(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    var cur int
    prev := 1
    pprev := 0
    
    for i := 2; i <= n; i ++ {
        cur = prev + pprev
        pprev = prev
        prev = cur
    }
    
    return cur
}

func main() {
    ret := fibRecursive(3)
    fmt.Println(ret)
    
    ret = fibRecursive(10)
    fmt.Println(ret)
    
    ret = fibMemo(10)
    fmt.Println(ret)
    
    ret = fibDP(10)
    fmt.Println(ret)
}
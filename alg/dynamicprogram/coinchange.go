package main
import (
    "fmt"
    "math"
    "time"
)
// 输入目标金额amount, 返回凑出目标金额的最少硬币数
func coinchange(amount int, coins []int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
    
    var count int = math.MaxInt32
    for _, v := range coins {
        ret := coinchange(amount - v, coins)
        if ret < 0 {
            continue
        }
        if ret < count {
            count = ret
        }
    }
    
    return count + 1 //最底层都是0开始，+ 1
}

func dpMemo(amount int, coins []int, states map[int]int) int {
    if amount == 0 {
        return 0
    }
    if amount < 0 {
        return -1
    }
      
    if v, ok := states[amount]; ok {
        return v
    }
      
    var count int = math.MaxInt32
    for _, v := range coins {
        ret := dpMemo(amount - v, coins, states)
        if ret < 0 {
            continue
        }
        if ret < count {
            count = ret
        }
    }
      
    states[amount] = count + 1
    return count + 1
}
func coinchangeMemo(amount int, coins []int) int {
    states := make(map[int]int)

    return dpMemo(amount, coins, states)
}


//dp array
func coinChangeBottomUp(amount int, coins []int) int {
    
    return 0
}

func main() {
    coins := []int{1, 3, 5}
    c1 := time.Now()
    count := coinchange(30, coins) // 121.403431ms
    c2 := time.Now()
    
    fmt.Println(count)
    fmt.Println(c2.Sub(c1)) 
    
    c3 := time.Now()
    count = coinchangeMemo(30, coins) //41.892µs
    c4 := time.Now()
    fmt.Println(c4.Sub(c3))
    fmt.Println(count)

}
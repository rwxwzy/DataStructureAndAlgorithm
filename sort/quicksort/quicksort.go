package main
import (
    "fmt"
)

/*
复杂度 NlogN (n次比较 * logN)
1.
从数列中挑出一个元素，称为 “基准”（pivot）
2.
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
3.
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
*/

func Partition(a []int, left, right int) int {
    //1. 选择最左边一个作为pivot
    // lo 指向小于pivot的的所有值中最右边的值
    current, lo := left + 1, left
    
    
    //2. 循环中, 只需要和lo + 1对调， 不需要和pivot对调
    // 只需要和小于pivot的所有值的最右边值的下一个交换即可
    for current <= right { //= 是因为也要和最右边的值比较
        if a[current] < a[left] {
            lo ++
            a[current], a[lo] = a[lo], a[current]
        }
        current ++
    }
    
    //3. 交换pivot和小于pivot的最右边的值
    a[left], a[lo] = a[lo], a[left]
    return lo
}

func QuickSort(a []int, left, right int) {
    if left >= right {
        return
    }
    pos := Partition(a, left, right)
    QuickSort(a, left, pos - 1)
    QuickSort(a, pos + 1, right)
}

func Print(a []int) {
    for i := 0; i < len(a); i ++ {
        fmt.Printf("%d ", a[i])
    }
    fmt.Printf("\n")
}
func main() {
    a := []int{10, 10, 9, 7, 8, 6, 5, 4, 3, 2, 0, 1, 0}
    Print(a)
    
    QuickSort(a, 0, len(a) - 1)
    fmt.Println("After Sort:")
    Print(a)
}
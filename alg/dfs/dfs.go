package main

import (
    "fmt"
)


type Node struct {
    Value interface{}
    Children []*Node
}

func DepthFirstSearch(root *Node, ret *[]*Node){
    *ret = append(*ret, root) // 因为append可能会改变ret的值，必须传入指针
    for _, val := range root.Children {
        DepthFirstSearch(val, ret)
    }
}

func main() {
/*
                1
     2          3         11
   5 6 7      8 9 10    4 12 13
*/
    n13 := &Node{13, nil}
    n12 := &Node{12, nil}
    n11 := &Node{4, nil}

    n10 := &Node{10, nil}
    n9 := &Node{9, nil}
    n8 := &Node{8, nil}

    n7 := &Node{7, nil}
    n6 := &Node{6, nil}
    n5 := &Node{5, nil}

    n4 := &Node{11, []*Node{n11, n12, n13}}
    n3 := &Node{3, []*Node{n8, n9, n10}}
    n2 := &Node{2, []*Node{n5, n6, n7}}

    n1 := &Node{1, []*Node{n2, n3, n4}}
    
    ret := make([]*Node, 0, 16)
    DepthFirstSearch(n1, &ret)
    
    fmt.Println("-----------------------")
    for _, val := range ret {
        fmt.Println(val.Value)
    }
}
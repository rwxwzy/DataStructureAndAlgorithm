package main

import (
    "fmt"
)

type Node struct{
    data int
    left *Node
    right *Node
}

type BinaryTree struct{
    root *Node
}

func (node *Node)AddNode(new *Node) {
    if node.data > new.data {
        if node.left != nil {
            node.left.AddNode(new)
        } else {
            node.left = new
        }
    } else {// right
        if node.right != nil {
            node.right.AddNode(new)
        } else {
            node.right = new
        }
    }
}

func (bt *BinaryTree)AddNode(data int) {
    n := &Node{data: data}
    if bt.root == nil {
        bt.root = n
    } else {
        bt.root.AddNode(n)
    }
}
func (bt *BinaryTree)PreOrder(node *Node){
    if node == nil {
        return
    }
    bt.PreOrder(node.left)
    fmt.Println(node.data)
    bt.PreOrder(node.right)
}

func(bt *BinaryTree)PostOrder(node *Node) {
    if node == nil {
        return
    }
    bt.PostOrder(node.right)
    fmt.Println(node.data)
    bt.PostOrder(node.left)
}

func(bt *BinaryTree)InOrder(node *Node) {
    if node == nil {
        return
    }
    fmt.Println(node.data)
    bt.InOrder(node.right)
    
    bt.InOrder(node.left)
}

// TODO: remove


func main() {
    b := &BinaryTree{}
    b.AddNode(10)
    b.AddNode(9)
    b.AddNode(12)
    b.AddNode(13)
    b.AddNode(9)
    b.AddNode(11)
    fmt.Println("PreOrder---------------")
    b.PreOrder(b.root)
    fmt.Println("PostOrder---------------")
    b.PostOrder(b.root)
    fmt.Println("InOrder---------------")
    b.InOrder(b.root)
}




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
            node.left.Add(new)
        } else {
            node.left = new
        }
    } else {// right
        if node.right != nil {
            node.right.Add(new)
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
func (bt *BinaryTree)PrintBefore(){
    
}

func main() {
    b := &BinaryTree{}
    b.AddNode(10)
    b.AddNode(9)
    b.AddNode(11)
}




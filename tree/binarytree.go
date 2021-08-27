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

/*
void traverse(TreeNode root) {
    // 前序遍历代码位置
    traverse(root.left)
    // 中序遍历代码位置
    traverse(root.right)
    // 后序遍历代码位置
}
*/


func (bt *BinaryTree)PreOrderTraverse(node *Node){
    if node == nil {
        return
    }
    fmt.Println(node.data)
    bt.PreOrderTraverse(node.left)
    bt.PreOrderTraverse(node.right)
}

func(bt *BinaryTree)InOrderTraverse(node *Node) {
    if node == nil {
        return
    }
    bt.InOrderTraverse(node.left)
    fmt.Println(node.data)
    bt.InOrderTraverse(node.right)
}

func(bt *BinaryTree)PostOrderTraverse(node *Node) {
    if node == nil {
        return
    }
    bt.PostOrderTraverse(node.left)
    bt.PostOrderTraverse(node.right)
    fmt.Println(node.data)
}

// TODO: remove


func main() {
    b := &BinaryTree{}
    b.AddNode(10)
    b.AddNode(9)
    b.AddNode(12)
    b.AddNode(13)
    b.AddNode(7)
    b.AddNode(11)
    b.AddNode(8)

    fmt.Println("PreOrderTraverse---------------")
    b.PreOrderTraverse(b.root)

    fmt.Println("InOrderTraverse---------------")
    b.InOrderTraverse(b.root)


    fmt.Println("PostOrderTraverse---------------")
    b.PostOrderTraverse(b.root)


}




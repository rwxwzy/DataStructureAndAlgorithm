package main

import (
    "fmt"
)

type Node struct {
   data int
   next *Node
}

type SList struct {
    length int
    head *Node
    tail *Node
}

func NewSList() *SList {
    s := &SList{length: 0, head: nil, tail: nil}
    return s
}

func (s *SList)AddNode(data int) error {
    node := &Node{data: data, next: nil}

    if s.head == nil {
        s.head = node
    } else {
        s.tail.next = node
    }
    s.tail = node
    s.length ++
    return nil
}
func (s *SList)DelNode(data int) error {
    if s.head == nil { // 0
        return nil
    }

    var prev *Node
    for i := s.head; i != nil; i = i.next {
        if i.data == data {
            s.length --
            if s.head == s.tail { // 1
                s.head = nil
                s.tail = nil
            } else { // > 1
                if s.head == i { // head
                    s.head = i.next
                } else if s.tail == i { // tail
                    s.tail = prev
                    s.tail.next = nil
                } else { //non head or tail
                    prev.next = i.next
                }
            }
        }

        prev = i
    }

    return nil
}

func (s *SList)Traverse(){
    fmt.Printf("slist length: %d\n", s.length)
    for i := s.head; i != nil; i = i.next {
        fmt.Println(i.data)
    }
}

func (s *SList) Reverse() {
    prev := s.head
    current := s.head
    next := s.head
    for current != nil{
        next = current.next
        current.next = prev
        prev = current
        current = next
    }

    s.tail = s.head
    s.tail.next = nil
    s.head = prev
}

func (s *SList) ReverseByRecursionWrapper() {
    s.tail = s.head
    last := s.ReverseByRecursion(s.head)
    s.head = last
}

func (s *SList)ReverseByRecursion(head *Node) *Node {
    if head.next == nil {
        return head
    }

    last := s.ReverseByRecursion(head.next)
    head.next.next = head
    head.next = nil
    fmt.Printf("last: %d\n", last.data)
    return last
}


func main() {
    slist := NewSList()
    slist.AddNode(9)
    slist.AddNode(1)
    slist.AddNode(2)
    slist.AddNode(3)
    slist.AddNode(4)
    slist.AddNode(5)
    slist.AddNode(6)
    slist.AddNode(7)
    slist.AddNode(8)
    slist.AddNode(9)
    slist.AddNode(0)
    slist.AddNode(9)
    slist.Traverse()

    slist.DelNode(9)
    slist.Traverse()

    /*
    slist.Reverse()
    slist.Traverse()
    */

    fmt.Println("recursion")
    slist.ReverseByRecursionWrapper()
    slist.Traverse()
    fmt.Println("finish")


}

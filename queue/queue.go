package main

import (
    "fmt"
)

type ItemType interface{}

type Queue struct {
    Items []ItemType
}

func New() *Queue {
    return &Queue{Items: make([]ItemType, 0)}
}

func (q *Queue) Push(a ItemType) {
    q.Items = append(q.Items, a)
}

func (q *Queue) Pop() *ItemType {
    if len(q.Items) == 0 {
        return nil
    }
    
    ret := q.Items[0]
    
    q.Items = q.Items[1: len(q.Items)]
    
    return &ret
}

func (q *Queue) PrintAll() {
    for _, v := range q.Items {
        fmt.Println(v)
    }
}

func main() {
    q := New()
    q.Push(0)
    q.Push(1)
    q.Push(2)
    q.Push(3)
    q.Push(4)
    q.Push(5)
    q.Push(6)
    q.Push(7)
    
    fmt.Printf("%v, %v, %v\n", *(q.Pop()), *(q.Pop()), *(q.Pop()))
    q.PrintAll()
}
package main

import (
    "fmt"
)

type ItemType interface{}

type Stack struct {
    Items []ItemType
}

func New() *Stack {
    return &Stack{Items: make([]ItemType, 0)}
}

func (s *Stack)Push(a ItemType) {
    s.Items = append(s.Items, a)
}
func (s *Stack)Pop() *ItemType {
    if len(s.Items) == 0 {
        return nil
    }
    
    ret := s.Items[len(s.Items) - 1]
    s.Items = s.Items[0: len(s.Items) - 1]
    return &ret
}

func(s *Stack)PrintAll() {
    for _, v := range s.Items {
        fmt.Println(v)
    }
}

func (s *Stack)GetAll()[]ItemType {
    return s.Items
}

func main() {
    s := New()
    s.Push(1)
    s.Push(2)
    s.Push(3)
    s.Push(4)
    s.Push(5)
    s.Push(6)
    s.Push(7)
    s.Pop()
    
    s.PrintAll()
}

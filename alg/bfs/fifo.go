package main

import (
    "fmt"
    "container/list"
)

type FIFO struct {
    queue []interface{}  //copy waste time
}

func New() *FIFO {
    return &FIFO{queue: make([]interface{}, 0)}
}

func (f *FIFO) Push(v interface{}) {
    f.queue = append(f.queue, v)
}

func (f *FIFO)PopFront() interface{}{
    if len(f.queue) == 0 {
        return nil
    }
    
    first := f.queue[0]
    f.queue[0] = nil // avoid memory leak
    f.queue = f.queue[1:]
    return first
}


func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
    
    /*
    f := New()
    for _, v := range a {
        f.Push(v)
    }
    
    for {
        b := f.PopFront()
        if b == nil {
            break
        }
        fmt.Println(b)
    }
    */

    f := list.New()
    
    for _, val := range a {
        f.PushBack(val)
    }
    
    for val := f.Front(); val != nil; val = val.Next() {
        fmt.Println(val.Value)
    }
}

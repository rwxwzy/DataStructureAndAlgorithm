package main
import (
    "fmt"
)

// https://github.com/golang/go/blob/eba0e866fafe3f8223d654a29fb953e02c07364a/src/container/list/list.go#L103



type Element struct {
    value interface{}
    prev *Element
    next *Element
}

type DoublyLinkedList struct {
    root Element  //sentinel
    len int
}

func New() *DoublyLinkedList{
    d := &DoublyLinkedList{}
    d.Init()
    return d
}

func (d *DoublyLinkedList) Init() {
    d.root.prev = &d.root //!!! sentinel root, pointer (root.prev) is not nil which is key to not check nil in other place
    d.root.next = &d.root //!!!
    d.len = 0
}
// insert after at
func (d *DoublyLinkedList) Insert(e *Element, at *Element) {
    e.prev = at
    e.next = at.next
    
    at.next.prev = e
    at.next = e
    
    d.len ++
}
func (d *DoublyLinkedList) InsertValue(v interface{}, at *Element) {
    e := &Element{value: v}
    d.Insert(e, at)
}

func (d *DoublyLinkedList) PushFront(v interface{}) {
    d.InsertValue(v, &d.root)
}

func (d *DoublyLinkedList) PushBack(v interface{}) {
    d.InsertValue(v, d.root.prev)
}

func (d *DoublyLinkedList) Remove(e *Element) {
    e.prev.next = e.next
    e.next.prev = e.prev
    d.len --
    e.next = nil  // avoid memory leak
    e.prev = nil  // avoid memory leak
}

// move next to at
func (d *DoublyLinkedList)Move(e, at *Element) {
    if e == at {
        return
    }
    
    e.prev.next = e.next
    e.next.prev = e.prev
    
    e.next = at.next
    e.prev = at
    at.next.prev = e
    at.next = e
}

func (d *DoublyLinkedList)Traverse() {
    e := d.root.next
    for i := 0; i < d.len; i ++ {
        fmt.Println(e.value)
        e = e.next
    }
}



func main() {
    dlist1 := New()
    for i := 0; i < 10; i ++ {
        dlist1.PushBack(i)
    }
    fmt.Println("PushBack-------------------------")
    dlist1.Traverse()
    
    dlist2 := New()
    for i := 0; i < 10; i ++ {
        dlist2.PushFront(i)
    }
    fmt.Println("PushFront------------------------")
    dlist2.Traverse()
}

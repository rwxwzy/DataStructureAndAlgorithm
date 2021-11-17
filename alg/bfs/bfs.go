package main

import (
    "fmt"
)

/*
    问题的本质就是让你在一幅「图」中找到从起点 start 到终点 target 的最近距离

    Breadth First Search (BFS) will explore a graph widely as opposed to deeply.
    While DFS tried to get to the endpoints of the graph as quickly as possible, BFS will explore a graph by layers

    In conclusion, the primary difference between DFS and BFS is that DFS immediately traverses the next child, while DFS will explore the graph by levels via a queue.

    https://levelup.gitconnected.com/depth-breadth-first-search-in-go-8a6ddcdc73d9


    https://cybernetist.com/2019/03/09/breadth-first-search-using-go-standard-library/



    Firstly, when traversing a graph, BFS splits the graph nodes into two groups:
    (1) 分为两类: visited, not yet visited
    (2) use FIFO queue for storing the graph

visited
not [yet] visited
By doing this we skip the nodes which have already been visited, thus avoiding cycles leading to infinite loops.

Secondly, the algorithm uses FIFO queue for storing the graph nodes that have yet to be visited. Both of these points will hopefully become more clear as we walk through the actual implementation.

(1) Initialize the queue for storing the nodes that have yet to be visited
(2) Pick any graph node and add it to the queue (NOTE: since this is a FIFO queue the node is pushed at the end of the queue)
(3) Take the first node from the front of the queue and add it to the list of visited nodes
(4) Add all adjacent nodes of the just visited node to the queue only if they haven’t already been visited
(5) Keep repeating 3 and 4 until the queue is empty




*/

type Node struct {
    Value interface{}
    Children []*Node
}

/*
从queue中拿出第一个处理，然后将第一个的children放到queue中
这样就可以将同一层的连续的放在一起
*/

func BreadthFirstSearch(root *Node) (ret []*Node){
    visited := make(map[*Node]struct{})

    queue := make([]*Node, 0)  //replace with "container/list"
    queue = append(queue, root)

    for len(queue) > 0 {// len会随着queue的变化而变化，而range slice是不变化的
        current := queue[0]

        visited[current] = struct{}{}

        ret = append(ret, current)


        queue = queue[1:]

        for _, val := range current.Children { //一层连续放在一起
            if _, ok := visited[val]; ok {
                continue
            }

            queue = append(queue, val)
        }
    }

    return ret
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


    ret := BreadthFirstSearch(n1)
    for _, v := range ret {
        fmt.Println(v.Value)
    }
}








package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    defer close(ch)
    var walk func(t *tree.Tree)
    walk = func(t *tree.Tree) {
        if t == nil {return}
        walk(t.Left)
        ch <- t.Value
        walk(t.Right)
    }
    walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)

    go Walk(t1, c1)
    go Walk(t2, c2)
    for {
        v1, ok1 := <-c1
        v2, ok2 := <-c2
        if (v1 != v2) || (ok1 != ok2) {
            return false
        }
        if !ok1 {
            return true
        }
    }
}

func main() {
    ch := make(chan int, 10)
    t := tree.New(1)
    go Walk(t, ch)
    for i := range ch {
        fmt.Println(i)
    }
    fmt.Println("t and tree.New(1) same?", Same(t, tree.New(1)))
    fmt.Println("t and tree.New(2) same?", Same(t, tree.New(2)))
}


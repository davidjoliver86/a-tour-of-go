package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    x, y := 0, 1
    return func() int {
        current := x
        next := x + y
        x = y
        y = next
        return current
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 20; i++ {
        fmt.Println(f())
    }
}


package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1

    fmt.Println("initial i >>> ", i)
    zeroval(i)
    
    fmt.Println("after zeroval i >>> ", i)

    zeroptr(&i)
    fmt.Println("after zeroptr i >>> ", i)

    fmt.Println("after zeroptr i 2 >>> ", &i)
}

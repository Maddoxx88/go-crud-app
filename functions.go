package main

import "fmt"

func multi(a int, b int) int {

    return a * b
}

func multiMulti(a, b, c int) int {
    return a * b * c
}

func main() {

    res := multi(1, 2)
    fmt.Println("1*2 =", res)

    res = multiMulti(1, 2, 3)
    fmt.Println("1*2*3 =", res)
}
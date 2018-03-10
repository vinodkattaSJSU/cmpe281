
package main

import "fmt"

func main() {

    b := [5] int {1, 2, 3, 4, 5}
    fmt.Println("Array:", b)
    fmt.Println("Array:", b[1:2])

    t := []string{"g", "h", "i"}
    fmt.Println("Slice:", t)
    fmt.Println("Slice:", t[1:2])

    primes := [6]int{2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)
    fmt.Println(cap(s))
    fmt.Println(len(s))
}

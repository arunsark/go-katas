package main
import "fmt"
import "time"
func main() {
    fmt.Println("hello world")
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's a weekend")
    default:
        fmt.Println("it's a week day")
    }
    t := time.Now()
    fmt.Println(t)
    switch {
    case t.Hour() < 12:
        fmt.Println("before noon")
    default:
        fmt.Println("after noon")
    }


   twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
    delete(n, "bar")
    fmt.Println("map:", n)
    j, prs := n["bar"]
    fmt.Println(j)
    fmt.Println(prs)
    nextSeq := intSeq()
    fmt.Println(nextSeq())
    fmt.Println(nextSeq())
    fmt.Println(nextSeq())
    nextSeq = intSeq()
    fmt.Println(nextSeq())
}

func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

package main

import "fmt"
import "github.com/arunsark/string"
import "github.com/arunsark/fizz"

func main() {
    fmt.Println(string.Reverse("Hello gopher"))
    fmt.Println(string.Reverse("Hello éàñ荷"))
    fmt.Println(fizz.FizzBuzz(1))
    fmt.Println(fizz.FizzBuzz(3))
    fmt.Println(fizz.FizzBuzz(5))
    fmt.Println(fizz.FizzBuzz(15))
}

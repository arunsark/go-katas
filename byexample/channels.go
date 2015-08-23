package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)

    messages = make(chan string, 2)
    messages <- "hello"
    messages <- "world"
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

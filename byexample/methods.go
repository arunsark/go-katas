package main

import "fmt"

type geometry interface {
    area() int
    perim() int
}

type rect struct {
    width, height int
}

func (r rect) area() int {
    return r.width * r.height;
}

func (r rect) perim() int {
    return r.width * 2 + r.height * 2;
}

func measure(g geometry) {
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main(){
    r := rect{width: 10, height: 20}
    measure(r)
}

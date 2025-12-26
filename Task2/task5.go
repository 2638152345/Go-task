package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func main() {
    var w, h float64
    fmt.Scan(&w, &h)
    rect := Rectangle{Width: w, Height: h}

    var r float64
    fmt.Scan(&r)
    circle := Circle{Radius: r}

    var s Shape

    s = rect
    fmt.Println(s.Area())
    fmt.Println(s.Perimeter())

    s = circle
    fmt.Println(s.Area())
    fmt.Println(s.Perimeter())
}

package main

import (
    "fmt"
    "io"
)

func sum(nums ...int) int {
    total := 0
    for _, n := range nums { total += n }
    return total
}

type Counter struct{ n int }
func (c *Counter) Inc() { c.n++ }
func (c Counter) Val() int { return c.n }

type Reader interface{ Read([]byte) (int, error) }
type MyBuf struct{}
func (MyBuf) Read(p []byte) (int, error) { return 0, io.EOF }

func main() {
    fmt.Println("sum:", sum(1,2,3))
    var c Counter; c.Inc(); fmt.Println("counter:", c.Val())
    var r Reader = MyBuf{}
    _, err := r.Read(make([]byte, 8))
    fmt.Println("read err:", err)
}


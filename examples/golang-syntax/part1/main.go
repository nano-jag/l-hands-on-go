package main

import (
    "fmt"
    "time"
)

func main() {
    var i int
    name := "Go"
    const Pi = 3.14159
    fmt.Printf("i=%d name=%s Pi=%.2f\n", i, name, Pi)

    nums := []int{1, 2, 3}
    nums = append(nums, 4)
    sub := nums[1:3]
    fmt.Println("nums:", nums, "sub:", sub)

    m := map[string]int{"apple": 2}
    if v, ok := m["apple"]; ok { fmt.Println("apple:", v) }

    fmt.Println("now:", time.Now().Format(time.RFC3339))
}


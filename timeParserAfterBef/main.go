package main

import (
	"fmt"
	"time"
)

var st1 = "2024-01-23T20:25:28Z"
var st2 = "2024-01-23T20:50:00Z"

func main() {
	St1, _ := time.Parse(time.RFC3339, st1)
	St2, _ := time.Parse(time.RFC3339, st2)
	fmt.Println(St1, St2, St2.After(St1))
}

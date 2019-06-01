package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Format(time.RFC3339)
	fmt.Println(t)
}

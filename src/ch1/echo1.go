package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now().UnixNano()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(time.Now().UnixNano() - now)
}

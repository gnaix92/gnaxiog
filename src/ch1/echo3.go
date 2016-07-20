package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now().UnixNano()

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(time.Now().UnixNano() - now)
}

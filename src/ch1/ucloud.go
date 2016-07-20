package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//go run ucloud.go 1.txt
func main() {
	count := 0
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ucloud: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			count += strings.Count(line, "UCanUup")
		}
		fmt.Println("UCloud:", count)
	}
}

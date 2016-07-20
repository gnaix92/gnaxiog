package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	keyfiles := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, keyfiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts, keyfiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, keyfiles[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, keyfiles map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		keyfiles[input.Text()] += f.Name() + " "
		counts[input.Text()]++
	}
}

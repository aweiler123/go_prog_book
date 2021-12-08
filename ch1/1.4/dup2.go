//Exercise 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

type stat struct {
	count     int
	filenames []string
}

func main() {
	counts := make(map[string]stat)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n", n.count, line)
			for _, filename := range n.filenames {
				fmt.Printf("  %s\n", filename)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]stat) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		filename := f.Name()
		text := input.Text()
		var temp stat
		temp.count = counts[text].count + 1
		temp.filenames = counts[text].filenames
		if !contains(temp.filenames, filename) {
			temp.filenames = append(temp.filenames, filename)
		}
		counts[text] = temp
	}
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if n == x {
			return true
		}
	}
	return false
}

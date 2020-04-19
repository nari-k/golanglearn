package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filemap := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filemap)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stdout, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filemap)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			filenames, sep := "", ""
			for filename := range filemap[line] {
				filenames += sep + filename
				sep = ","
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filemap map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		str := input.Text()
		counts[str]++
		if filemap[str] == nil {
			filemap[str] = make(map[string]int)
		}
		filemap[str][f.Name()]++
	}
}

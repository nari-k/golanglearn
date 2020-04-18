package main

import (
	"fmt"
	"io"
	"os"
)

func Echo(out io.Writer, args []string) {
	for _, arg := range args {
		fmt.Fprintln(out,arg)
	}
}

func main() {
	Echo(os.Stdout, os.Args[0:])
}

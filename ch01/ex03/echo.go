package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func EchoEfficient(out io.Writer, args []string) {
	fmt.Fprintln(out, strings.Join(args, " "))
}

func EchoInefficient(out io.Writer, args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

func main() {
	EchoEfficient(os.Stdout, os.Args[0:])
	EchoInefficient(os.Stdout, os.Args[0:])
}
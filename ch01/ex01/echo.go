package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Echo(out io.Writer, args []string) {
	//os.Args[0:]からって、テストじゃないなぁ。。。
	fmt.Fprintln(out, strings.Join(args, " "))
}

func main() {
	Echo(os.Stdout, os.Args[0:])
}

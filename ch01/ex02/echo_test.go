package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T){
	out := new(bytes.Buffer)
	Echo(out, []string{"./echo", "aaaa", "oooo"})

	actual := out.String()
	expected := "./echo\naaaa\noooo\n"
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}


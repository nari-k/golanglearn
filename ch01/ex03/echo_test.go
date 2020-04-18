package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestEcho(t *testing.T){
	oute := new(bytes.Buffer)
	EchoEfficient(oute, []string{"./echo", "aaaa", "oooo"})

	actual := oute.String()

	outi := new(bytes.Buffer)
	EchoInefficient(outi, []string{"./echo", "aaaa", "oooo"})
	expected := outi.String()
	//同じになる確認
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func BenchmarkEchoEfficient(b *testing.B) {
	args := []string{"./echo", "aaaa", "oooo"}
	for i := 0; i < b.N; i++ {
		EchoEfficient(ioutil.Discard, args)
	}
}
func BenchmarkEchoInefficient(b *testing.B) {
	args := []string{"./echo", "aaaa", "oooo"}
	for i := 0; i < b.N; i++ {
		EchoInefficient(ioutil.Discard, args)
	}
}

package tempconv

import (
	"testing"
)

func TestCToK(t *testing.T) {
	k := CtoK(Celsius(AbsoluteZeroC))
	if k != Kelvin(0) {
		t.Errorf("TestCToK actual: %f, expected: %f", k, Kelvin(0))
	}
}
func TestKtoC(t *testing.T) {
	c := KtoC(Kelvin(0))
	if c != AbsoluteZeroC {
		t.Errorf("TestKtoC actual: %f, expected: %f", c, AbsoluteZeroC)
	}
}
func TestFtoK(t *testing.T) {
	k := FtoK(Fahrenheit(-268.15))
	if k != Kelvin(-15) {
		t.Errorf("TestFtoK actual: %f, expected: %f", k, Kelvin(-15))
	}
}
func TestKtoF(t *testing.T) {
	f := KtoF(Kelvin(41))
	if f != Fahrenheit(-385.87) {
		t.Errorf("TestKtoF actual: %f, expected: %f", f, Fahrenheit(-385.87))
	}
}

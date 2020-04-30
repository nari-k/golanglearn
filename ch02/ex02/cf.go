package main

import (
	"./lengthconv"
	"./tempconv"
	"./weightconv"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func print(t float64) {

	fa := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", fa, tempconv.FtoC(fa), c, tempconv.CtoF(c))
	fe := lengthconv.Feet(t)
	m := lengthconv.Meters(t)
	fmt.Printf("%s = %s, %s = %s\n", fe, lengthconv.FtoM(fe), m, lengthconv.MtoF(m))
	k := weightconv.Kilogram(t)
	p := weightconv.Pound(t)
	fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KtoP(k), p, weightconv.PToK(p))
}

func main() {

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cf %v\n", err)
			}
			print(t)
		}
	} else {
		input := bufio.NewScanner(os.Stdout)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "cf %v\n", err)
			}
			print(t)
		}
	}
}

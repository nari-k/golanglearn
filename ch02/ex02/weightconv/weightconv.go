package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func KtoP(k Kilogram) Pound { return Pound(k / 453.59237) }
func PToK(p Pound) Kilogram { return Kilogram(p * 453.59237) }

func (c Kilogram) String() string { return fmt.Sprintf("%gkg", c) }
func (f Pound) String() string    { return fmt.Sprintf("%glb", f) }

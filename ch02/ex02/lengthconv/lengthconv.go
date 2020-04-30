package lengthconv

import "fmt"

type Feet float64
type Meters float64

func MtoF(m Meters) Meters { return Meters(m * 0.3048) }
func FtoM(f Feet) Feet     { return Feet(f / 0.3048) }

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string   { return fmt.Sprintf("%gft", f) }

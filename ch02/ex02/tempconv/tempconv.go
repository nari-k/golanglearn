package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingF     Celsius = 0
	BoilingF      Celsius = 0
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }

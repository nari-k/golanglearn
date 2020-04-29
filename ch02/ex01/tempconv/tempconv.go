package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingF     Celsius = 0
	BoilingF      Celsius = 0
)

func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}
func KtoC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func FtoK(f Fahrenheit) Kelvin {
	return Kelvin(FtoC(f + 273.15))
}
func KtoF(k Kelvin) Fahrenheit {
	return CtoF(Celsius(k - 273.15))
}

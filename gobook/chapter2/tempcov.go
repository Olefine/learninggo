package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func main() {
	useTypes()
}

func useTypes() {
	var c Celsius = 100.0

	fmt.Println(c.String())
	//No need to call String() explicitly, analog of toString()
	fmt.Println(c)
}

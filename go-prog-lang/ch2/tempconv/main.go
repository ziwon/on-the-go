package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroc Celsius = -273.15
	FreesingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func main() {
	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) compile error: type mismatch
	fmt.Println(c == Celsius(f))

	c1 := FToC(212.0)
	fmt.Println(c1.String())
	fmt.Printf("%v\n", c1)
	fmt.Printf("%s\n", c1)
	fmt.Println(c1)
	fmt.Printf("%g\n", c1)
	fmt.Println(float64(c1))
}

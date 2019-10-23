/*
* Добавьте в пакет tempconv типы, константы и функции для работы
* с температурой по шкале Кельвина, в которой нуль градусов соответствует
* температуре -273.15°С, а разница температур в 1К имеет ту же величину,
* что и 1°С.
 */

package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func main() {
	t := BoilingK
	f := Fahrenheit(t)
	c := Celsius(t)
	k := Kelvin(t)
	fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, FToC(f), c, CToF(c), k, KToC(k))
}

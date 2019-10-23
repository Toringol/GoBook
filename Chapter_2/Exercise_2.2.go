/*
* Напишите программу общего назначения для преобразования единиц,
* аналогичную cf, которая считывает числа из аргументов командной строки
* (или из стандартного ввода, если аргументы командной строки отсутствуют) и
* преобразует каждое число в другие единицы, как температуру - в градусы
* Цельсия и Фаренгейта, длину - в футы и метры, вес в фунты и килограммы и т.д.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

type Foot float64
type Meter float64

type Pound float64
type Kilogram float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func FToM(f Foot) Meter { return Meter(f / 3.281) }
func MToF(m Meter) Foot { return Foot(m * 3.281) }

func PToK(p Pound) Kilogram { return Kilogram(p / 2.205) }
func KToP(k Kilogram) Pound { return Pound(k * 2.205) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func TemperatureConv(numeric float64) {
	far := Fahrenheit(numeric)
	cel := Celsius(numeric)
	fmt.Printf("Temperature: %s = %s, %s = %s\n", far, FToC(far), cel, CToF(cel))
}

func LengthConv(numeric float64) {
	foot := Foot(numeric)
	meter := Meter(numeric)
	fmt.Printf("Length: %s = %s, %s = %s\n", foot, FToM(foot), meter, MToF(meter))
}

func WeigthConv(numeric float64) {
	pound := Pound(numeric)
	kilogram := Kilogram(numeric)
	fmt.Printf("Weight: %s = %s, %s = %s\n", pound, PToK(pound), kilogram, KToP(kilogram))
}

func ConvertValues(numeric float64) {
	TemperatureConv(numeric)
	LengthConv(numeric)
	WeigthConv(numeric)
}

func main() {

	if len(os.Args) == 1 {
		in := bufio.NewScanner(os.Stdin)

		for in.Scan() {
			numeric, err := strconv.ParseFloat(in.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			ConvertValues(numeric)
		}
	} else {
		for _, arg := range os.Args[1:] {
			numeric, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
			ConvertValues(numeric)
		}
	}

}

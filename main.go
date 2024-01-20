package main

import "fmt"

func main() {
	/*
		Convert Kelvin to Celsius

		Boiling temperature in Celsius = 100
	*/

	boilingTempCelsius := 100.0
	boilingTempKelvim := boilingTempCelsius + 273.0

	fmt.Printf("A temperatura de ebulição em Celsius é de %g ºC que é equivalente a %v K.", boilingTempCelsius, boilingTempKelvim)
}

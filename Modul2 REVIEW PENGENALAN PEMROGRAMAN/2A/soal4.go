package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	reamur := (4.0 / 5.0) * celsius
	fahrenheit := (9.0 / 5.0) * celsius + 32
	kelvin := celsius + 273

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
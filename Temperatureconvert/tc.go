// Problem statement 
// Implement function toFahrenheit(t Celsius)to convert from Celsius to Fahrenheit. Types of temperatures are float32 which are aliased to Celsius and Fahrenheit for you.

// Temperature in ˚F = (Temp in °C * 9/5) + 32

// Input 
// A variable of type Celsius

// Output 
// A variable of type Fahrenheit
package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func main(){
	fmt.Println(toFahrenheit(100))
}

func toFahrenheit(t Celsius) Fahrenheit {
	k :=  Fahrenheit((t * 9/5) + 32)
	fmt.Println("This is %g",k)
	return k
}
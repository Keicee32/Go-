package main

import (
	"fmt"
	"math"
)
	
type Converter struct {
}
type Centimeter float64
type Feet float64

type Seconds float64
type Minutes float64

type Fahrenheit float64
type Celsius float64

type Kilogram float64
type Pound float64

type Radian float64
type Degree float64

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet{
	var what = Feet(c / 30.48)
	return what
}

func (cvr Converter) FeetToCentimeter(f Feet) Centimeter{
	var what = Centimeter(f * 30.48)
	return what
}

func (cvr Converter) MinutesToSeconds(m Minutes) Seconds{
	var min = Seconds(m*60)
	return min 
}

func (cvr Converter) SecondsToMinutes(s Seconds) Minutes{
	var sec = Minutes(s / 60)
	return sec
}

func (cvr Converter) FahrenheitToCelsius(F Fahrenheit) Celsius{
	var fah = Celsius(F - 32) * 0.555
	return fah
}

func (cvr Converter) CelsiusToFahrenheit(C Celsius) Fahrenheit{
	var cel = Fahrenheit(C * 1.8) + 32
	return cel
}

func (cvr Converter) KilogramToPound(K Kilogram) Pound{
	var kilo = Pound(K * 2.205) 
	return kilo
}

func (cvr Converter) PoundToKilogram(P Pound) Kilogram{
	var pound = Kilogram(P / 2.205) 
	return pound
}

func (cvr Converter) RadianToDegree(R Radian) Degree{
	var rad = Degree (R * (180/math.Pi)) 
	return rad
}

func (cvr Converter) DegreeToRadian(R Radian) Radian{
	var deg = Radian (R * (math.Pi/180)) 
	return deg
}

func main(){
	cvr := Converter{}
	fmt.Println(cvr.FeetToCentimeter(50))
	fmt.Println(cvr.SecondsToMinutes(60))
	fmt.Println(cvr.CelsiusToFahrenheit(50))
	fmt.Println(cvr.DegreeToRadian(50))
}
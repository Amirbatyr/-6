package main

import "fmt"

func main() {

	// Задание 1
	age := 16
	fmt.Println(age)

	age = age + 1
	fmt.Println(age)

	// Задание 2
	height := 170
	fmt.Println(height)

	height_in_meters := 1.7
	fmt.Println(height_in_meters)

	// Задание 3
	isStudent := true
	fmt.Println(isStudent)

	// Задание 4
	temperature := 25
	fmt.Println(temperature)

	if temperature >= 20 {
		fmt.Println("Погода теплая")
	} else {
		fmt.Println("Погода холодная")
	}

	// Задание 5
	favoriteQuote := "Никогда не сдавайся"
	fmt.Println(favoriteQuote)

	// Задание 6
	PI := 3.14
	fmt.Println(PI)

	// PI = "3.1415"
	// Это не сработает, потому что PI уже имеет тип float64,
	// а строку нельзя присвоить переменной типа float64.
}

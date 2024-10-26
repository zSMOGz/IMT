package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")
	userHeight, userWeight := getUserInput()
	IMT := claculateIMT(userHeight, userWeight)

	if IMT < 16 {
		fmt.Println("Сильный дефицит массы тела")
	} else if IMT < 18.5 {
		fmt.Println("Дефицит массы тела")
	} else if IMT < 25 {
		fmt.Println("Нормальнный вес")
	} else if IMT < 30 {
		fmt.Println("Избыточный вес")
	} else {
		fmt.Println("Ожирение")
	}

	outputResult(IMT)
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)
}

func claculateIMT(userHeight float64, userWeight float64) float64 {
	const IMTPower = 2

	return userWeight / math.Pow(userHeight/100, IMTPower)
}

func getUserInput() (float64, float64) {
	var userHeight float64
	var userWeight float64

	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}

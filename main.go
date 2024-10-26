package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("__ Калькулятор индекса массы тела __")

	isContinueCalculation := true
	for isContinueCalculation {
		userHeight, userWeight := getUserInput()
		IMT, err := claculateIMT(userHeight, userWeight)

		if err != nil {
			fmt.Println(err)
			isContinueCalculation = true
			continue
		}

		outputResult(IMT)

		isContinueCalculation = checkRepeatCalculation()
	}
}

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	fmt.Println(result)

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
}

func claculateIMT(userHeight float64, userWeight float64) (float64, error) {
	const IMTPower = 2

	if userHeight <= 0 || userWeight <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}

	return userWeight / math.Pow(userHeight/100, IMTPower), nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Хотите повторить расчёт (y/n):")
	fmt.Scan(&userChoise)

	if userChoise == "y" || userChoise == "Y" {
		return true
	} else if userChoise == "n" || userChoise == "N" {
		return false
	} else {
		return checkRepeatCalculation()
	}
}

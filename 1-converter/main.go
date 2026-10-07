package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	const USD_EUR float64 = 0.8884
	const USD_RUB float64 = 84.75
	const RUB_EUR float64 = 0.0105
	dictValue := map[string]float64{
		"USD_EUR": USD_EUR,
		"USD_RUB": USD_RUB,
		"EUR_USD": 1 / USD_EUR,
		"EUR_RUB": 1 / RUB_EUR,
		"RUB_EUR": RUB_EUR,
		"RUB_USD": 1 / USD_RUB,
		"RUB_RUB": 1.0,
		"USD_USD": 1.0,
		"EUR_EUR": 1.0,
	}
	var first string = getFirstCurrency()
	var amount = getAmount()
	var second string = getSecondCurrency()
	result := calculate(amount, dictValue[fmt.Sprintf("%s_%s", first, second)])
	fmt.Printf("Результат - %f", result)
}

func getAmount() float64 {
	fmt.Println("Введите число:")
	var amount float64
	for {
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка при считывании значения, введите заново:")
		} else {
			break
		}
	}
	return amount
}

func getFirstCurrency() string {
	fmt.Println("Введите исходную валюту (USD, RUB, EUR):")
	var firstCurrency string
	for {
		fmt.Scan(&firstCurrency)
		if checkCurrency(firstCurrency) {
			break
		}
		fmt.Println("Данная валюта не поддерживатся программой, введите заново:")
	}
	return strings.ToUpper(firstCurrency)
}

func getSecondCurrency() string {
	fmt.Println("Введите целевую валюту (USD, RUB, EUR):")
	var secondCurrency string
	for {
		fmt.Scan(&secondCurrency)
		if checkCurrency(secondCurrency) {
			break
		}
		fmt.Println("Данная валюта не поддерживатся программой, введите заново:")
	}
	return strings.ToUpper(secondCurrency)
}

func checkCurrency(name string) bool {
	values := []string{"USD", "RUB", "EUR"}
	upperStr := strings.ToUpper(name)
	if slices.Contains(values, upperStr) {
		return true
	} else {
		return false
	}
}

func calculate(amount, cur float64) float64 {
	result := amount * cur
	return result
}

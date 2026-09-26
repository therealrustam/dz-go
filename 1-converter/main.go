package main

import "fmt"

func main() {
	const USD_EUR float64 = 0.86
	const USD_RUB float64 = 86.7
	var EUR_RUB float64
	var amount = getAmount()
	EUR_RUB = USD_RUB/USD_EUR
	calculate(amount, "USD", "RUB")
	fmt.Println(amount, EUR_RUB)
}

func getAmount() int {
	var amount int
	fmt.Scan(&amount)
	return amount
}

func calculate (amount int, currentCurrency string, newCurrency string) {
	
}
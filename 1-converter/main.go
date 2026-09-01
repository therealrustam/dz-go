package main

import "fmt"

func main() {
	const USD_EUR float64 = 0.86
	const USD_RUB float64 = 86.7
	var EUR_RUB float64
	EUR_RUB = USD_RUB/USD_EUR
	var amount int = 5
	fmt.Println(amount*int(EUR_RUB))
}
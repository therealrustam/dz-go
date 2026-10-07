package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Создаем один сканер для всей программы
	scanner := bufio.NewScanner(os.Stdin)

	operation := getOperation(scanner)
	data := getData(scanner)

	if len(data) == 0 {
		fmt.Println("Нет данных для вычислений")
		return
	}

	var result float64
	switch operation {
	case "AVG":
		result = avg(data)
	case "SUM":
		result = sum(data)
	case "MED":
		result = med(data)
	}

	fmt.Printf("Результат - %f\n", result)
}

func getOperation(scanner *bufio.Scanner) string {
	operations := []string{"AVG", "SUM", "MED"}
	fmt.Println("Введите тип операции ('AVG' - среднее, 'SUM' - сумма, 'MED' - медиана):")

	for scanner.Scan() {
		operation := strings.ToUpper(strings.TrimSpace(scanner.Text()))
		if slices.Contains(operations, operation) {
			return operation
		}
		fmt.Println("Некорректный тип операции, введите снова:")
	}
	return ""
}

func getData(scanner *bufio.Scanner) []int {
	fmt.Println("Введите числа через запятую (например: 1, 2, 3):")

	if !scanner.Scan() {
		return nil
	}

	line := scanner.Text()
	parts := strings.Split(line, ",")
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		num, err := strconv.Atoi(trimmed)
		if err != nil {
			fmt.Printf("Ошибка при конвертации '%s': %v\n", trimmed, err)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func sum(data []int) float64 {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return float64(sum)
}

func avg(data []int) float64 {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return float64(sum) / float64(len(data))
}

func med(data []int) float64 {
	if len(data) == 0 {
		return 0
	}
	sorted := make([]int, len(data))
	copy(sorted, data)
	sort.Ints(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	mid1 := sorted[n/2-1]
	mid2 := sorted[n/2]
	return float64(mid1+mid2) / 2.0
}

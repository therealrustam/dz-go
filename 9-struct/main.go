package main

import (
	"fmt"
	. "main/bins"
)

func main() {
	binList := NewBinList()
	bins := binList.Bins
	for {
		fmt.Println("Введите команду(1 - создать, 2 - выход):")
		var command uint
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Не удалось считать команду, введите заново:")
			continue
		}
		if command == 1 {
			bin := CreateBin()
			bins = append(bins, bin)
		} else {
			break
		}
	}
}

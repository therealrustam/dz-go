package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func newBinList() *BinList {
	return &BinList{bins: []Bin{}}
}

func main() {
	binList := newBinList()
	bins := binList.bins
	for {
		fmt.Println("Введите команду(1 - создать, 2 - выход):")
		var command uint
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Не удалось считать команду, введите заново:")
			continue
		}
		if command == 1 {
			bin := createBin()
			bins = append(bins, bin)
		} else {
			break
		}
	}
}

func createBin() Bin {
	var id string
	var private bool
	var name string
	for {
		fmt.Println("Введите id:")
		_, err := fmt.Scan(&id)
		if err == nil {
			break
		} else {
			fmt.Println("Ошибка при чтении id, введите заново:")
		}
	}
	for {
		fmt.Println("Введите private:")
		_, err := fmt.Scan(&private)
		if err == nil {
			break
		} else {
			fmt.Println("Ошибка при чтении private, введите заново:")
		}
	}
	for {
		fmt.Println("Введите название:")
		_, err := fmt.Scan(&name)
		if err == nil {
			break
		} else {
			fmt.Println("Ошибка при чтении name, введите заново:")
		}
	}
	bin := Bin{id, private, time.Now(), name}
	return bin
}

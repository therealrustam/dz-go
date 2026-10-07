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

func main() {
	binList := []*Bin{}
	for {
		fmt.Println("Введите команду(1 - создать, 2 - выход):")
		var command uint
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Не удалось считать команду, введите заново:")
		}
		if command == 1 {
			bin := createBin()
			binList = append(binList, bin)
		} else {
			break
		}
	}
}

func createBin() *Bin {
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
	return &bin
}

package bins

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBinList() *BinList {
	return &BinList{Bins: []Bin{}}
}

func CreateBin() Bin {
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

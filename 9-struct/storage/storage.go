package storage

import (
	"os"
	"fmt"
	"main/bins"
	"encoding/json"
)

func saveBins(binList *bins.BinList) {
	data, err := json.Marshal(binList.Bins)
	if err != nil {
		fmt.Println("Ошибка перевода данных в json")
	}
	err = os.WriteFile("bins.json", data, 0644)
		if err != nil {
			fmt.Println("Ошибка записи файла:", err)
			return
		}
	fmt.Println("Данные успешно сохранены в bins.json")
}

func readBins() bins.BinList {
	var bins bins.BinList
	data, err := os.ReadFile("bins.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return bins
	}
	err = json.Unmarshal(data, &bins)
	if err != nil {
		fmt.Println("Ошибка перевода данных в json")
		return bins
	}
	return bins
}

func storage() {

}

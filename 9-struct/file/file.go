package file

import (
	"os"
	"fmt"
	"path/filepath"
)

func ReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil
	}
	return data
}

func IsJSON(name string) bool {
	if filepath.Ext(name) == ".json" {
		return true
	}
	return false
}

func file() {

}

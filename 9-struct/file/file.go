package file

import (
	"os"
	"fmt"
	"encoding/json"
)

func ReadFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil
	}
	var js json.RawMessage
	if err := json.Unmarshal(data, &js); err != nil {
		fmt.Println("Расширение файла не соответствует json")
		return nil
	}
	return data
}

func file() {

}

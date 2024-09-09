package tgz4b

import (
	"encoding/json"
	"io"
	"os"
)

type AppTG struct {
	API_KEY_TG string `json:"API_KEY_TG"`
}

// Загрузить данные из файла
func LoadConfig(filename string) (config AppTG, ErrorFIle error) {
	// Открыть файл
	jsonFile, ErrorFIle := os.Open(filename)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFIle := io.ReadAll(jsonFile)
	if ErrorFIle != nil {
		return config, ErrorFIle
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &config); ErrorFIle != nil {
		return config, ErrorFIle
	}
	return config, ErrorFIle
}

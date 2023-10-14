package zachestnyibiznes

import (
	"encoding/json"
	"io"
	"os"
)

type Z4B struct {
	API_KEY string `json:"API_KEY"`
	URL     string `json:"URL"`
}

// Загрузить данные из файла
func LoadConfig(filename string) (config Z4B, ErrorFIle error) {
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

package couterparser

import (
	"encoding/json"
	"io"
	"os"

	api2captcha "github.com/2captcha/2captcha-go"
)

// Загрузить данные 2captcha из файла
func LoadConfig(filename string) (config *ParseStruct, ErrorFIle error) {
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

	// Создаём объект
	config.CaptchaClient = api2captcha.NewClient(config.Token)
	return config, ErrorFIle
}

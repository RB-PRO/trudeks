package tgcouter

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Token  string `json:"Token"`
	ChatID int64  `json:"ChatId"`
}

func LoadConfig(FileName string) (cf Config, Err error) {
	// Прочитать файл
	fileBytes, Err := os.ReadFile(FileName)
	if Err != nil {
		return cf, fmt.Errorf("os.ReadFile: %v", Err)
	}

	// Распарсить
	Err = json.Unmarshal(fileBytes, &cf)
	if Err != nil {
		return cf, fmt.Errorf("json.Unmarshal: %v", Err)
	}

	return cf, nil
}

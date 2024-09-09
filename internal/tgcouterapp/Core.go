// Пакет телеграм бота для работы с судами
//
// бот отпавить судебные дела админу
package tgcouterapp

import (
	"fmt"

	tgcouter "trudeks/pkg/tgcouter"
)

// Структура для работы приложения для защиты файла
type Secret struct {
	tg          *tgcouter.Telegram
	messageChan chan string
}

// Создаём приложение сервиса
func New() (*Secret, error) {
	var Err error
	s := &Secret{}

	// Канал общения TG и парсера
	s.messageChan = make(chan string, 1)

	// TELEGRAM
	cf, Err := tgcouter.LoadConfig("tg.json")
	if Err != nil {
		return nil, fmt.Errorf("tg.LoadConfig: %v", Err)
	}
	fmt.Println("Загрузил конфиг для телеграма")

	s.tg = &tgcouter.Telegram{}
	s.tg, Err = tgcouter.NewTelegram(cf)
	if Err != nil {
		return nil, fmt.Errorf("tg.NewTelegram: %v", Err)
	}
	fmt.Println("Создано приложение телеграм")

	return s, nil
}

func (s *Secret) Run() error {
	return s.tg.Watch()
}

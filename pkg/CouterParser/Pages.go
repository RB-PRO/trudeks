package couterparser

import (
	"fmt"
	"time"
)

func Pages(CouterURL string) (meets []Meeting, Err error) {
	// Перменная, которая содержит о существовании следующей страницы
	var NextPageIsExit bool = true

	// От такого числа (От текущей даты отнимаем месяц)
	DateFrom := time.Now().AddDate(-1, 0, 0).Format("02.01.2006")
	// До такого числа (Текущая дата)
	DateTo := time.Now().Format("02.01.2006")

	for page := 1; NextPageIsExit; page++ {

		// Формируем ссылку на страницу
		url := fmt.Sprintf(PrefixURL, CouterURL, DateFrom, DateTo, page) + PostFix

		fmt.Println(page)

		// Парсим текущую страницу
		var meetsLocal []Meeting
		meetsLocal, NextPageIsExit, Err = Page(url)
		if Err != nil {
			return nil, fmt.Errorf("Pages: %w", Err)
		}

		// Сохраняем данные в общем слайсе судебных дел.
		meets = append(meets, meetsLocal...)
	}

	return meets, nil
}

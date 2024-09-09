package couterparser

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Спасить судебные дела на одноим листе и проверить существование следущего листа
func Page(link string) (meets []Meeting, NextIsExit bool, Err error) {

	c := colly.NewCollector()

	// Определение существования следущего листа
	c.OnHTML("td[valign=bottom]", func(element *colly.HTMLElement) {
		if element.DOM.Find("a:last-child ").Text() == ">>" {
			NextIsExit = true
		}
	})

	// Парсинг контента
	c.OnHTML("table[id=tablcont]>tbody>tr", func(e *colly.HTMLElement) {
		var meet Meeting

		// Колонка "№ дела"
		NumberDOM := e.DOM.Find("td:nth-child(1)>a")
		NumberStr := NumberDOM.Text() // Номер и код дела
		NumberStrs := strings.Split(NumberStr, "~")
		if len(NumberStrs) == 2 {
			meet.Number = strings.TrimSpace(NumberStrs[0])
			meet.Code = strings.TrimSpace(NumberStrs[1])
		} else {
			meet.Number = strings.TrimSpace(NumberStr)
		}
		meet.Link, _ = NumberDOM.Attr("href")

		// Колонка "Дата поступления дела"
		DateReceiptStr := e.DOM.Find("td:nth-child(2)").Text()
		DateReceiptStr = strings.TrimSpace(DateReceiptStr)
		if DateReceiptStr != "" {
			DateReceipt, ErrParseDateReceipt := time.Parse("02.01.2006", DateReceiptStr)
			if ErrParseDateReceipt == nil {
				meet.DateReceipt = DateReceipt
			}
		}

		// Колонка "Категория / Стороны"
		CategoryStr, _ := e.DOM.Find("td:nth-child(3)").Html()
		CategoryStrs := strings.Split(CategoryStr, "<br/>") // Разделяем, чтобы кластеризовать от сторон дела
		if len(CategoryStrs) > 0 {                          // Если есть хоть какое-нибудь содержимое(оно будет в любом случае)
			CategoryStr = strings.TrimSpace(CategoryStrs[0]) // Чистим строку
			if strings.Contains(CategoryStr, "КАТЕГОРИЯ:") { // Оставляем голую информацию
				CategoryStr = strings.ReplaceAll(CategoryStr, "КАТЕГОРИЯ:", "")
				CategoryStrs = strings.Split(CategoryStr, "→") // Разделяем строку на какие-то стрелочки между категориями
				if len(CategoryStrs) > 0 {
					for i := range CategoryStrs {
						CategoryStrs[i] = strings.TrimSpace(CategoryStrs[i])
					}
					meet.Category = CategoryStrs
				}
			}
		}

		// Колонка "Судья"
		Judge := e.DOM.Find("td:nth-child(4)").Text()
		meet.Judge = strings.TrimSpace(Judge)

		// Колонка "Дата решения"
		DateDoneStr := e.DOM.Find("td:nth-child(5)").Text()
		DateDoneStr = strings.TrimSpace(DateDoneStr)
		if DateDoneStr != "" {
			DateDone, ErrParseDateDone := time.Parse("02.01.2006", DateDoneStr)
			if ErrParseDateDone == nil {
				meet.DateDone = DateDone
			}
		}

		// Колонка "Решшение"
		DoneReport := e.DOM.Find("td:nth-child(6)").Text()
		meet.DoneReport = strings.TrimSpace(DoneReport)

		// Колонка "Дата вступления в законную силу"
		DateEffectiveStr := e.DOM.Find("td:nth-child(7)").Text()
		DateEffectiveStr = strings.TrimSpace(DateEffectiveStr)
		if DateEffectiveStr != "" {
			// В этой ячейке может содержаться как дата,
			// так и информация о том, обжалуется дело или нет.
			DateEffective, ErrParseDateReceipt := time.Parse("02.01.2006", DateEffectiveStr)
			if ErrParseDateReceipt == nil {
				meet.DateEffective = DateEffective
			} else {
				if strings.Contains(DateEffectiveStr, "ОБЖАЛУЕТСЯ") {
					meet.Appealed = true // Дело обжалуется
				}
			}
		}

		// Колонка "Судебные акты"
		meet.CourtActURL, _ = e.DOM.Find("td:nth-child(8)>a").Attr("href")

		// Сохранение результата парсинга
		if meet.Link != "" {
			meets = append(meets, meet)
		}
	})

	// Поситить сайт с целью скрепинга ;)
	Err = c.Visit(link)
	if Err != nil {
		return nil, false, fmt.Errorf("Page: Visit: %w", Err)
	}

	return meets, NextIsExit, nil
}

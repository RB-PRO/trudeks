package region

import (
	"fmt"
	"strings"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	"github.com/gocolly/colly"
)

// Константная ссылка, пример
//
//	fmt.Sprintf(URLpage, couterLink, date.Format("02.01.2006"))
//
// [пример]: https://himki--mo.sudrf.ru/modules.php?name=sud_delo&srv_num=1&H_date=11.12.2023
const URLpage string = "%s/modules.php?name=sud_delo&srv_num=1&H_date=%s"

// Собрать список судебных дел в суде couterLink за определённую дату date
func Page(couterLink string, date time.Time) (meets []gocouter.Meeting, Err error) {
	var TypeStr string
	c := colly.NewCollector()

	// Парсинг контента
	c.OnHTML(`table[id=tablcont]>tbody>tr`, func(e *colly.HTMLElement) {

		// Пропускаем header таблицы
		if e.Attr("align") == "center" {
			return
		}

		// Запоминаем тип дела:
		//	- Гражданские дела - апелляция
		//	- Административные дела (КАC РФ) - первая инстанция
		//	- Дела об административных правонарушениях - первая инстанция
		//	- Дела об административных правонарушениях - жалобы на постановления
		//	- Гражданские дела - первая инстанция
		//	- Уголовные дела - первая инстанция
		//	- Производство по материалам
		if e.Attr("bgcolor") == "#DEDEDE" {
			TypeStr = e.Text
			return
		}

		var meet gocouter.Meeting
		meet.Type = TypeStr

		// Колонка "№ дела"
		NumberDOM := e.DOM.Find("td:nth-child(2)>a")
		NumberStr := NumberDOM.Text() // Номер и код дела
		NumberStrs := strings.Split(NumberStr, "~")
		if len(NumberStrs) == 2 {
			meet.Number = strings.TrimSpace(NumberStrs[0])
			meet.Code = strings.TrimSpace(NumberStrs[1])
		} else {
			meet.Number = strings.TrimSpace(NumberStr)
		}
		// Ссылка на дело
		Link, _ := NumberDOM.Attr("href")
		meet.Link = couterLink + Link
		meet.DateCouterProcess = date

		// Колонка "Категория / Стороны"
		CategoryStr, _ := e.DOM.Find("td:nth-child(5)").Html()
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
		Judge := e.DOM.Find("td:nth-child(6)").Text()
		meet.Judge = strings.TrimSpace(Judge)

		// Колонка "Решшение"
		DoneReport := e.DOM.Find("td:nth-child(7)").Text()
		meet.DoneReport = strings.TrimSpace(DoneReport)

		// Колонка "Судебные акты"
		CourtActURL, _ := e.DOM.Find("td:nth-child(8)>a").Attr("href")
		if CourtActURL != "" {
			meet.CourtActURL = couterLink + CourtActURL
		}

		// Сохранение результата парсинга
		if meet.Link != "" {
			meets = append(meets, meet)
		}
	})

	// Поситить сайт с целью скрепинга ;)
	url := fmt.Sprintf(URLpage, couterLink, date.Format("02.01.2006"))
	Err = c.Visit(url)
	if Err != nil {
		return nil, fmt.Errorf("Page: %s Visit: %w", url, Err)
	}

	return meets, nil
}

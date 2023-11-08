package couterparser

import (
	"fmt"

	"github.com/gocolly/colly"
)

func (cr *Couter) ParseCase(link string) (cs Case, Err error) {

	c := colly.NewCollector()

	// Уникальный идентификатор дела
	c.OnHTML("div[id=cont1]>table>tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			if h.DOM.Find("td:first-of-type b").Text() ==
				"Уникальный идентификатор дела" {
				SelectorUID := h.DOM.Find("td:last-of-type a")
				cs.Idntifier = SelectorUID.Text()
				cs.IdntifierLink, _ = SelectorUID.Attr("href")
			}
		})
	})

	// Стороны дела
	c.OnHTML("div[id=cont3]>table>tbody", func(e *colly.HTMLElement) {
		cs.Attack = make([]Side, 0, 1)  // Выделяем память в Истца
		cs.Defense = make([]Side, 0, 1) // Выделяем память в ответчик
		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			// Истец
			if h.DOM.Find("td:nth-of-type(1)").Text() ==
				"ИСТЕЦ" {
				cs.Attack = append(cs.Attack, Side{
					Name: h.DOM.Find("td:nth-of-type(2)").Text(),
					INN:  h.DOM.Find("td:nth-of-type(3)").Text(),
					KPP:  h.DOM.Find("td:nth-of-type(4)").Text(),
					OGRN: h.DOM.Find("td:nth-of-type(5)").Text(),
				})
			}

			// Ответчик
			if h.DOM.Find("td:nth-of-type(1)").Text() ==
				"ОТВЕТЧИК" {
				cs.Defense = append(cs.Attack, Side{
					Name: h.DOM.Find("td:nth-of-type(2)").Text(),
					INN:  h.DOM.Find("td:nth-of-type(3)").Text(),
					KPP:  h.DOM.Find("td:nth-of-type(4)").Text(),
					OGRN: h.DOM.Find("td:nth-of-type(5)").Text(),
				})
			}
		})
	})

	// Поситить сайт с целью скрепинга ;)
	Err = c.Visit(link)
	if Err != nil {
		return Case{}, fmt.Errorf("ParseCase: Visit: %w", Err)
	}
	return cs, nil
}

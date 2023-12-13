package region

import (
	"fmt"
	"strings"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	"github.com/gocolly/colly"
)

func ParseCase(link string) (cs gocouter.Case, Err error) {

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
		cs.Attack = make([]gocouter.Side, 0, 1)  // Выделяем память в Истца
		cs.Defense = make([]gocouter.Side, 0, 1) // Выделяем память в ответчик
		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			RowName := h.DOM.Find("td:nth-of-type(1)").Text()
			RowName = strings.TrimSpace(RowName)

			// Истец
			if RowName == "ИСТЕЦ" {
				cs.Attack = append(cs.Attack, gocouter.Side{
					Name: h.DOM.Find("td:nth-of-type(2)").Text(),
					INN:  h.DOM.Find("td:nth-of-type(3)").Text(),
					KPP:  h.DOM.Find("td:nth-of-type(4)").Text(),
					OGRN: h.DOM.Find("td:nth-of-type(5)").Text(),
				})
			}

			// Ответчик
			if RowName == "ОТВЕТЧИК" {
				cs.Defense = append(cs.Defense, gocouter.Side{
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
		return gocouter.Case{}, fmt.Errorf("ParseCase: Visit: %w", Err)
	}
	return cs, nil
}

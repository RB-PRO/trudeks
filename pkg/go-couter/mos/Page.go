package mos

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	gocouter "trudeks/pkg/go-couter"
)

const PageLink string = "https://mos-gorsud.ru/search?caseDateFrom=%s&caseDateTo=%s&category=%s&processType=2&formType=fullForm&page=%d"

// Спарсить каждую страницу по запросу
func Pages(DateFrom, DateTo time.Time, category string) (meets []gocouter.Meeting, Err error) {
	DateFromStr := DateFrom.Format("02.01.2006")
	DateToStr := DateTo.Format("02.01.2006")

	Next := true
	for page := 1; Next; page++ {
		url := fmt.Sprintf(PageLink, DateFromStr, DateToStr, category, page)
		// fmt.Println(url)
		var linksMeets []gocouter.Meeting
		linksMeets, Next, Err = Page(url)
		if Err != nil {
			return nil, Err
		}
		meets = append(meets, linksMeets...)
		time.Sleep(time.Millisecond * 300)
	}

	return meets, Err
}

func Page(url string) (Meets []gocouter.Meeting, Next bool, Err error) {
	// Create a collector
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.660 YaBrowser/23.9.5.660 Yowser/2.5 Safari/537.36"

	c.OnHTML("div[class=searchResultContainer] table[class=custom_table]>tbody>tr", func(e *colly.HTMLElement) {
		var Meet gocouter.Meeting

		// if link, isexit := e.DOM.Attr("href"); isexit {
		// 	urlMeets = append(urlMeets, URL+link)
		// }

		// Номер дела ~ материала
		NumberSel := e.DOM.Find("td:nth-child(1)>nobr:nth-child(1)>a")
		Meet.Number = strings.TrimSpace(NumberSel.Text())
		Meet.Link, _ = NumberSel.Attr("href")
		Meet.Link = URL + Meet.Link
		CodeSel := e.DOM.Find("td:nth-child(1)>nobr:nth-child(2)")
		Meet.Code = strings.TrimSpace(CodeSel.Text())

		// Стороны
		td2 := e.DOM.Find("td:nth-child(2)>div>div[class=right]")
		strsSides := strings.Split(td2.Text(), "Ответчик:")
		if len(strsSides) == 2 {
			strsSides[0] = strings.ReplaceAll(strsSides[0], "Истец: ", "")
			strsSides[0] = strings.TrimSpace(strsSides[0])
			Atacks := strings.Split(strsSides[0], ",")
			for _, atack := range Atacks {
				Meet.Case.Attack = append(Meet.Case.Attack, gocouter.Side{Name: atack})
			}

			strsSides[1] = strings.TrimSpace(strsSides[1])
			Defenses := strings.Split(strsSides[1], ",")
			for _, defense := range Defenses {
				Meet.Case.Defense = append(Meet.Case.Defense, gocouter.Side{Name: defense})
			}
		}

		// Текущее состояние
		td3 := e.DOM.Find("td:nth-child(3)")
		Meet.Status = strings.TrimSpace(td3.Text())

		// Судья
		td4 := e.DOM.Find("td:nth-child(4)")
		Meet.Judge = strings.TrimSpace(td4.Text())

		// Статья

		// Категория дела
		td6 := e.DOM.Find("td:nth-child(6)")
		Meet.Category = append(Meet.Category, strings.TrimSpace(td6.Text()))

		Meets = append(Meets, Meet)
	})

	// Проверка существование следующего листа
	c.OnHTML("a[class=intheend]", func(e *colly.HTMLElement) {
		if _, isexit := e.DOM.Attr("disabled"); !isexit {
			Next = true
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("failed with response: %v Error: %v", r, err)
	})
	Err = c.Visit(url)
	return Meets, Next, Err
}

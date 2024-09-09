package mos

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
	gocouter "trudeks/pkg/go-couter"
)

func ParseMeet(url string) (Meet gocouter.Meeting, Err error) {
	// Create a collector
	c := colly.NewCollector()

	// Основные сведения по делу
	c.OnHTML("div[class=cardsud_wrapper]", func(e *colly.HTMLElement) {
		e.ForEach("div[class=row_card]", func(i int, h *colly.HTMLElement) {
			LeftStr := strings.TrimSpace(h.DOM.Find("div[class=left]").Text())
			RightContent := h.DOM.Find("div[class=right]")
			RightStr := strings.TrimSpace(RightContent.Text())

			switch LeftStr {
			case "Уникальный идентификатор дела":
				Meet.Case.Idntifier = RightStr
			case "Номер дела":
				Meet.Number = RightStr
			case "Суд первой инстанции, судья":
				strs := strings.Split(RightStr, "(")
				if len(strs) == 2 {
					Meet.CouterName = strings.TrimSpace(strs[0])
					strs[1] = strings.ReplaceAll(strs[1], ")", "")
					Meet.Judge = strings.TrimSpace(strs[1])
				}
			case "Дата поступления дела в апелляционную инстанцию":
				Meet.DateReceipt = ParsingDate(RightStr)
			case "Номер дела в суде нижестоящей инстанции":
				Meet.Code = strings.TrimSpace(RightStr)
				Meet.Link, _ = RightContent.Attr("href")
			case "Категория дела":
				Meet.Category = append(Meet.Category, strings.TrimSpace(RightStr))
			// case "Стороны":
			// 	// fmt.Println(RightStr)
			// 	RightContent.Find("").Each(func(i int, s *goquery.Selection) {
			// 		fmt.Println(i, s.Text())
			// 	})

			default:
				// fmt.Println(LeftStr, "-", RightStr)
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		Err = fmt.Errorf("failed with response: %v Error: %v", r, err)
	})

	Err = c.Visit(url)

	return Meet, Err
}

func ParsingDate(dateStr string) time.Time {
	dateStr = strings.TrimSpace(dateStr)
	date, _ := time.Parse("02.01.2006", dateStr)
	return date
}

package couterparser

import (
	"fmt"
	"log"
	"time"

	api2captcha "github.com/2captcha/2captcha-go"
)

// Структура для решения капчи
type Captcha struct {
	CaptchaClient *api2captcha.Client `json:"-"`
	Token         string              `json:"ParseStruct"`
}

// Структура объект для решения капчи
func NewParseStruct(token string) *Captcha {
	return &Captcha{CaptchaClient: api2captcha.NewClient(token)}
}

func (cr *Couter) Pages(CouterURL string) (meets []Meeting, Err error) {

	// От такого числа (От текущей даты отнимаем месяц)
	DateFrom := time.Now().AddDate(0, -1, 0).Format("02.01.2006")
	// До такого числа (Текущая дата)
	DateTo := time.Now().Format("02.01.2006")

	// Обработка капчи. Логика такая:
	// Делаем запрос на определение существования капчи.
	// В случае существования капчи, в цикле с определённым к-вом иттераций ищем решение
	CapthaCode := ""
	for i := 0; CapthaCode == ""; i++ {
		if i == 5 { // 5 попыток на парсинг
			return nil, fmt.Errorf("CaptchaParse: Сделано 5 запросов, но ответа не получилось %w", Err)
		}

		// Получаем капчу
		base64captcha, IDcaptcha, Err := CaptchaParse(CouterURL)
		if Err != nil {
			return nil, fmt.Errorf("CaptchaParse: %w", Err)
		}

		// Если капча не предусмотрена для данного запроса
		if IDcaptcha == "" {
			break
		}

		// Стртуктура запросника
		cap := api2captcha.Normal{
			Base64:   base64captcha, // base64 картинки
			Numberic: 1, MinLen: 5, MaxLen: 5, Lang: "en",
			HintText: "5 digits, lines on top of the picture",
		}
		code, err := cr.Ch.CaptchaClient.Solve(cap.ToRequest())
		if err != nil {
			if err == api2captcha.ErrTimeout {
				log.Println("ps.CaptchaClient.Solve: Timeout")
			} else if err == api2captcha.ErrApi {
				log.Println("ps.CaptchaClient.Solve: API error")
			} else if err == api2captcha.ErrNetwork {
				log.Println("ps.CaptchaClient.Solve: Network error")
			} else {
				log.Println(err)
			}
		}
		CapthaCode = fmt.Sprintf("&captcha=%s&captchaid=%s", code, IDcaptcha)
	}
	fmt.Println("CapthaCode", CapthaCode)

	// Перменная, которая содержит о существовании следующей страницы
	var NextPageIsExit bool = true
	for page := 1; NextPageIsExit; page++ {

		// Формируем ссылку на страницу
		url := fmt.Sprintf(PrefixURL, CouterURL, DateFrom, DateTo, page) + cr.PostFix + CapthaCode
		// fmt.Println(url)
		// fmt.Println()

		// Парсим текущую страницу
		var meetsLocal []Meeting
		meetsLocal, NextPageIsExit, Err = Page(url)
		if Err != nil {
			return nil, fmt.Errorf("Page: %w", Err)
		}

		// Сохраняем данные в общем слайсе судебных дел.
		meets = append(meets, meetsLocal...)
	}

	return meets, nil
}

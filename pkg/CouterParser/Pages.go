package couterparser

import (
	"fmt"
	"log"
	"time"

	api2captcha "github.com/2captcha/2captcha-go"
)

// Структура для решения капчи
type ParseStruct struct {
	CaptchaClient *api2captcha.Client `json:"-"`
	Token         string              `json:"ParseStruct"`
}

// Структура объект для решения капчи
func NewParseStruct(token string) *ParseStruct {
	return &ParseStruct{CaptchaClient: api2captcha.NewClient(token)}
}

func (ps *ParseStruct) Pages(CouterURL string) (meets []Meeting, Err error) {
	// Перменная, которая содержит о существовании следующей страницы
	var NextPageIsExit bool = true

	// От такого числа (От текущей даты отнимаем месяц)
	DateFrom := time.Now().AddDate(-1, 0, 0).Format("02.01.2006")
	// До такого числа (Текущая дата)
	DateTo := time.Now().Format("02.01.2006")

	// Обработка капчи. Логика такая:
	// Делаем запрос на определение существования капчи.
	// В случае существования капчи, в цикле с определённым к-вом иттераций ищем решение
	CapthaCode := ""
	for i := 0; ; {
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
			Numberic: 1,             // Только цифры
			MinLen:   5,             // Минимальное к-во символов
			MaxLen:   5,             // Максимальное к-во символов
			Lang:     "en",
			HintText: "5 digits, lines on top of the picture",
		}
		code, err := ps.CaptchaClient.Solve(cap.ToRequest())
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

	for page := 1; NextPageIsExit; page++ {

		// Формируем ссылку на страницу
		url := fmt.Sprintf(PrefixURL, CouterURL, DateFrom, DateTo, page) + PostFix + CapthaCode
		fmt.Println(url)
		fmt.Println()

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

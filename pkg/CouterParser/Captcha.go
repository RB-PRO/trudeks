package couterparser

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Получить  base64 и ID каптчи
func CaptchaParse(URL string) (base64captcha, IDcaptcha string, Err error) {

	c := colly.NewCollector()

	// Уникальный идентификатор дела
	c.OnHTML("form[id=calform] table table>tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, h *colly.HTMLElement) {
			ColName := h.DOM.Find("td:first-of-type").Text()
			if strings.Contains(ColName, "Проверочный код") {
				SelectorCaptchaID := h.DOM.Find("td:last-of-type")
				base64captcha, _ = SelectorCaptchaID.Find("img").Attr("src")
				IDcaptcha, _ = SelectorCaptchaID.Find("input[name=captchaid]").Attr("value")
				return
			}
		})
	})

	// Поситить сайт
	URL += "/modules.php?name=sud_delo&srv_num=1&name_op=sf&delo_id=1540005&case_type=0"
	Err = c.Visit(URL)
	if Err != nil {
		return "", "", fmt.Errorf("CaptchaParse: Visit: %w", Err)
	}
	return base64captcha, IDcaptcha, nil
}

package couterparser

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// Получить  base64 и ID каптчи
func CaptchaParse(URL string) (base64captcha, IDcaptcha string, Err error) {

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.5845.931 YaBrowser/23.9.3.931 Yowser/2.5 Safari/537.36"

	// капча
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
	c.OnHTML(`img[style="border: 1px solid #4a739d;"`, func(e *colly.HTMLElement) {
		if IDcaptcha == "" {
			base64captcha, _ = e.DOM.Parent().Find("img").Attr("src")
			IDcaptcha, _ = e.DOM.Parent().Find("input[name=captchaid]").Attr("value")
		}
	})

	// Поситить сайт
	URL += "/modules.php?name=sud_delo&srv_num=1&name_op=sf&delo_id=1540005&case_type=0"
	fmt.Println("URL", URL)
	Err = c.Visit(URL)
	if Err != nil {
		return "", "", fmt.Errorf("CaptchaParse: Visit: %w", Err)
	}
	return base64captcha, IDcaptcha, nil
}

package zachestnyibiznes

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Загрузить список товаров
func (z4b Z4B) Contacts(value string) (contacts Сontacts, ErrLine error) {

	// Сформировать ссылку
	url, ErrParse := url.Parse(z4b.URL)
	if ErrParse != nil {
		return Сontacts{}, ErrParse
	}
	url = url.JoinPath("contacts")
	q := url.Query()
	q.Set("api_key", z4b.API_KEY)
	q.Set("id", value)
	q.Set("_format", "json")
	url.RawQuery = q.Encode()

	client := &http.Client{}
	req, ErrNewRequest := http.NewRequest(http.MethodGet, url.String(), nil)
	if ErrNewRequest != nil {
		return Сontacts{}, ErrNewRequest
	}

	// Выполнить запорос
	Response, ErrDo := client.Do(req)
	if ErrDo != nil {
		return Сontacts{}, ErrDo
	}
	defer Response.Body.Close()

	// Получить массив []byte из ответа
	BodyPage, ErrorReadAll := io.ReadAll(Response.Body)
	if ErrorReadAll != nil {
		return Сontacts{}, ErrorReadAll
	}

	// os.WriteFile("LinePost.txt", BodyPage, 0666)

	// Распарсить полученный json в структуру
	ErrorUnmarshal := json.Unmarshal(BodyPage, &contacts)
	if ErrorUnmarshal != nil {
		return Сontacts{}, ErrorUnmarshal
	}

	return contacts, ErrNewRequest
}

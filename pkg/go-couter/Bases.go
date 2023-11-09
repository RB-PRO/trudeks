package couterparser

import (
	"os"
)

// Основная рабочая ссылка
const PrefixURL string = `%s/modules.php?g1_case__ENTRY_DATE1D=%s&g1_case__ENTRY_DATE2D=%s&page=%d`

// Структура по работе с судами
type Couter struct {
	PostFix string //
	Ch      *Captcha
}

func NewCouter(FileNamePostFix string) (*Couter, error) {

	// Загружаем данные по Postfix
	b, ErrReadFile := os.ReadFile(FileNamePostFix) // just pass the file name
	if ErrReadFile != nil {
		return nil, ErrReadFile
	}

	return &Couter{
		PostFix: string(b),
	}, nil
}

package tgz4b

import (
	"fmt"

	couter "github.com/RB-PRO/trudeks/pkg/go-couter"
)

// Спарсить суды по списку
func ParsingCounterGoRoutines() (FileName string, Err error) {
	// загрузка данных
	cr, ErrNewCouter := couter.NewCouter("PostFix")
	if ErrNewCouter != nil {
		return "", fmt.Errorf("couter.NewCouter: %w", ErrNewCouter)
	}
	pg, ErrLoadPG := couter.LoadConfig("2captcha.json")
	if ErrLoadPG != nil {
		return "", fmt.Errorf("couter.LoadConfig: %w", ErrLoadPG)
	}
	cr.Ch = pg

	MapCouter, ErrInput := couter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		return "", fmt.Errorf("couter.InputFileXlsxCouter: %w", ErrInput)
	}
	if len(MapCouter) == 0 {
		return "", fmt.Errorf("couter.InputFileXlsxCouter: Lenght of MapCouter = 0")
	}

	cn := make(chan<- []couter.Meeting)
	var MEETS []couter.Meeting
	for Region, mCouter := range MapCouter { // Регионы
		fmt.Println("Исследуем регион:", Region)

		wg.Add(1)

		cr.ParseRegion(Region, mCouter, cn)

	}

	return "", nil
}

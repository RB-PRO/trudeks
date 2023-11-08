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

	for Region, mCouter := range MapCouter { // Регионы
		fmt.Println("Исследуем регион:", Region)
		for nROI := 0; nROI < len(mCouter); nROI++ { // Суды региона

			fmt.Println("Исследуем суд:", mCouter[nROI].Name, mCouter[nROI].URL)

			//  Получаем список всех судебных дел
			meets, ErrPages := cr.Pages(mCouter[nROI].URL)
			if ErrPages != nil {
				return "", fmt.Errorf("couter.Pages: %w for url couter %s", ErrPages, mCouter[nROI].URL)
			}
			fmt.Println("Всего судебных дел в суде", mCouter[nROI].Name, ":", len(meets))

			// Подробности по каждому судебному деву
			for imeet := range meets {
				fmt.Println(imeet, "/", len(meets))
				meets[imeet].Link = mCouter[nROI].URL + meets[imeet].Link
				cs, ErrCase := cr.ParseCase(mCouter[nROI].URL + meets[imeet].Link)
				if ErrCase != nil {
					return "", fmt.Errorf("couter.ParseCase: %w for url couter %s", ErrCase, mCouter[nROI].URL+meets[imeet].Link)
				}
				meets[imeet].Case = cs
			}
		}

		break
	}

	return "", nil
}

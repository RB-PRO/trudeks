package tgcouter

import (
	"fmt"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	region "github.com/RB-PRO/trudeks/pkg/go-couter/region"
)

// Спасить Всё, кроме Московской области
func ParseRussia() (FileName string, Err error) {

	MapCouter, ErrInput := gocouter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		panic(ErrInput)
	}
	if len(MapCouter) == 0 {
		panic("Lenght of MapCouter = 0")
	}

	todayDate := time.Now() // Текущая дата
	ROI_days := 60          // Выборка по дням, за которую нужно парсить сайты судов

	FileName = fmt.Sprintf("russia_%s.xlsx", todayDate.Format("15-04_02-01-2006"))
	xx, errxlsx := gocouter.NewSavingFile(FileName)
	if errxlsx != nil {
		return "", fmt.Errorf("gocouter.NewSavingFile: %v", FileName)
	}
	xx.SetHead()

	for Region, Couters := range MapCouter {
		if Region == "Московская область" {
			continue
		}

		for _, Couter := range Couters {

			URL_couter := Couter.URL
			for day := 0; day < ROI_days; day++ {

				// Добавляем к дате один день
				TecalDay := todayDate.AddDate(0, 0, day)
				// fmt.Println(day, TecalDay.Format("02.01.2006"))

				// Собираем данные со страницы по определённому суду и дате
				MeetsCouterDay, ErrCouterDay := region.Page(URL_couter, TecalDay)
				if ErrCouterDay != nil {
					fmt.Printf("Парсинг страницы суда %s судебных дел за %s: %s - %v\n",
						URL_couter, TecalDay.Format("02.01.2006"), fmt.Sprintf(region.URLpage, URL_couter, TecalDay.Format("02.01.2006")), ErrCouterDay)
				}
				time.Sleep(300 * time.Millisecond)

				// region.TrudFilter()
				xx.InputRows(Region, Couter.Name, region.TrudFilter(MeetsCouterDay))
			}
		}
	}

	xx.Save()
	xx.Close()

	return FileName, nil
}

// Спасить Всё, кроме Московской области
func ParseMO() (FileName string, Err error) {

	MapCouter, ErrInput := gocouter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		panic(ErrInput)
	}
	if len(MapCouter) == 0 {
		panic("Lenght of MapCouter = 0")
	}

	todayDate := time.Now() // Текущая дата
	ROI_days := 60          // Выборка по дням, за которую нужно парсить сайты судов

	FileName = fmt.Sprintf("mo_%s.xlsx", todayDate.Format("15-04_02-01-2006"))
	xx, errxlsx := gocouter.NewSavingFile(FileName)
	if errxlsx != nil {
		return "", fmt.Errorf("gocouter.NewSavingFile: %v", FileName)
	}
	xx.SetHead()

	Region := "Московская область"
	Couters := MapCouter[Region]
	for _, Couter := range Couters {

		URL_couter := Couter.URL
		for day := 0; day < ROI_days; day++ {

			// Добавляем к дате один день
			TecalDay := todayDate.AddDate(0, 0, day)
			// fmt.Println(day, TecalDay.Format("02.01.2006"))

			// Собираем данные со страницы по определённому суду и дате
			MeetsCouterDay, ErrCouterDay := region.Page(URL_couter, TecalDay)
			if ErrCouterDay != nil {
				fmt.Printf("Парсинг страницы суда %s судебных дел за %s: %s - %v\n",
					URL_couter, TecalDay.Format("02.01.2006"), fmt.Sprintf(region.URLpage, URL_couter, TecalDay.Format("02.01.2006")), ErrCouterDay)
			}
			time.Sleep(300 * time.Millisecond)

			// region.TrudFilter()
			xx.InputRows(Region, Couter.Name, region.TrudFilter(MeetsCouterDay))
		}
	}

	xx.Save()
	xx.Close()

	return FileName, nil
}

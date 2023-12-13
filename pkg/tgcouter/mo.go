package tgcouter

import (
	"fmt"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	region "github.com/RB-PRO/trudeks/pkg/go-couter/region"
)

// Спасить Всё, кроме Московской области
func ParseRussia(updmsg *UpdateMassage) (FileName string, Err error) {

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

	var iRegion int
	for Region, Couters := range MapCouter {
		if Region == "Московская область" {
			continue
		}

		for iCouter, Couter := range Couters {
			URL_couter := Couter.URL
			for day := 0; day < ROI_days; day++ {

				if day%10 == 0 {
					updmsg.Update(fmt.Sprintf("Парсинг РФ\n> Регион[%d/%d]: %s\n> Суд[%d/%d]: %s\n> %d-й день из %d",
						iRegion+1, len(MapCouter), Region, iCouter+1, len(Couters), Couter.Name, day, ROI_days))
				}

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

				// Получение данных по сторонам
				FilterMeets := region.TrudFilter(MeetsCouterDay)

				for iMeet := range FilterMeets {
					if iMeet%10 == 0 {

						updmsg.Update(fmt.Sprintf("Парсинг РФ\n> Регион[%d/%d]: %s\n> Суд[%d/%d]: %s\n> %d-й день из %d\n> Сбор трудовых дел %d из %d",
							iRegion+1, len(MapCouter), Region, iCouter+1, len(Couters), Couter.Name, day, ROI_days, iMeet+1, len(FilterMeets)))

					}

					cases, ErrReg := region.ParseCase(FilterMeets[iMeet].Link)
					if ErrReg != nil {
						fmt.Printf("Парсинг дела %s: %v\n", FilterMeets[iMeet].Link, ErrReg)
					}
					time.Sleep(300 * time.Millisecond)
					FilterMeets[iMeet].Case = cases
				}

				// fmt.Println(FilterMeets)

				xx.InputRows(Region, Couter.Name, FilterMeets)
			}
		}
		iRegion++
	}

	xx.Save()
	xx.Close()
	updmsg.Update(("Пропарсил всю Россию"))

	return FileName, nil
}

// Спасить Всё, кроме Московской области
func ParseMO(updmsg *UpdateMassage) (FileName string, Err error) {

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
	for iCouter, Couter := range Couters {
		updmsg.Update(fmt.Sprintf("Парсинг МО\nРегион: %s\nСуд[%d/%d]: %s",
			Region, iCouter+1, len(Couters), Couter.Name))

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

			// Получение данных по сторонам
			FilterMeets := region.TrudFilter(MeetsCouterDay)

			for iMeet := range FilterMeets {
				if iMeet%10 == 0 {

					updmsg.Update(fmt.Sprintf("Парсинг МО\n> Регион: %s\n> Суд[%d/%d]: %s\n> %d-й день из %d\n> Сбор трудовых дел %d из %d",
						Region, iCouter+1, len(Couters), Couter.Name, day, ROI_days, iMeet+1, len(FilterMeets)))

				}

				cases, ErrReg := region.ParseCase(FilterMeets[iMeet].Link)
				if ErrReg != nil {
					fmt.Printf("Парсинг дела %s: %v\n", FilterMeets[iMeet].Link, ErrReg)
				}
				time.Sleep(300 * time.Millisecond)
				FilterMeets[iMeet].Case = cases
			}

			// fmt.Println(FilterMeets)

			xx.InputRows(Region, Couter.Name, FilterMeets)
		}
	}

	xx.Save()
	xx.Close()
	updmsg.Update(("Пропарсил всю Московскую область"))

	return FileName, nil
}

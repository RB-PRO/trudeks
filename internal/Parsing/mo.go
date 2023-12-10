package parsing

import (
	"fmt"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	"github.com/RB-PRO/trudeks/pkg/go-couter/region"
	"github.com/cheggaaa/pb"
)

func MO2() {

	MapCouter, ErrInput := gocouter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		panic(ErrInput)
	}
	if len(MapCouter) == 0 {
		panic("Lenght of MapCouter = 0")
	}
	xx, errxlsx := gocouter.NewSavingFile("regions.xlsx")
	if errxlsx != nil {
		panic(errxlsx)
	}
	xx.SetHead()

	todayDate := time.Now() // Текущая дата
	ROI_days := 60          // Выборка по дням, за которую нужно парсить сайты судов

	ALL := 0

	BarRegion := pb.StartNew(len(MapCouter))
	for Region, Couters := range MapCouter {

		// BarRegion.Prefix(fmt.Sprintf("%s :[%d/%d]", Region, icateg, len(Categorys)))

		for iCouter, Couter := range Couters {
			BarRegion.Prefix(fmt.Sprintf("%s :[%d/%d]", Region, iCouter, len(Couters)))

			URL_couter := Couter.URL
			for day := 0; day < ROI_days; day++ {

				// Добавляем к дате один день
				TecalDay := todayDate.AddDate(0, 0, day)
				// fmt.Println(day, TecalDay.Format("02.01.2006"))

				// Собираем данные со страницы по определённому суду и дате
				MeetsCouterDay, ErrCouterDay := region.Page(URL_couter, TecalDay)
				if ErrCouterDay != nil {
					panic(fmt.Sprintf("Парсинг страницы суда %s судебных дел за %s: %s - %v",
						URL_couter, TecalDay.Format("02.01.2006"), fmt.Sprintf(region.URLpage, URL_couter, TecalDay.Format("02.01.2006")), ErrCouterDay))
				}
				time.Sleep(time.Second)

				ALL += len(MeetsCouterDay)

				xx.InputRows(Region, Couter.Name, region.TrudFilter(MeetsCouterDay))
			}

		}

		BarRegion.Increment()
	}
	BarRegion.Finish()
	fmt.Println("ALL", ALL)
	fmt.Println("ALL", ALL)
	fmt.Println("ALL", ALL)
	// URL_couter := "https://himki--mo.sudrf.ru"
	// for day := 0; day < ROI_days; day++ {

	// 	// Добавляем к дате один день
	// 	TecalDay := timeNow.AddDate(0, 0, day)
	// 	fmt.Println(day, TecalDay.Format("02.01.2006"))

	// 	// Собираем данные со страницы по определённому суду и дате
	// 	MeetsCouterDay, ErrCouterDay := region.Page(URL_couter, TecalDay)
	// 	if ErrCouterDay != nil {
	// 		panic(fmt.Sprintf("Парсинг страницы суда %s судебных дел за %s: %s - %v",
	// 			URL_couter, TecalDay.Format("02.01.2006"), fmt.Sprintf(region.URLpage, URL_couter, TecalDay.Format("02.01.2006")), ErrCouterDay))
	// 	}
	// 	time.Sleep(time.Second)

	// 	xx.InputRows("Московская область", "Химкинский городской суд", MeetsCouterDay)
	// }
	xx.Save()
	xx.Close()
}

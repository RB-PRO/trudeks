package parsing

import (
	"fmt"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
	"github.com/RB-PRO/trudeks/pkg/go-couter/mos"
)

func MSK() {
	Codes := mos.Codes()
	var meets []gocouter.Meeting

	DateFrom := time.Now().AddDate(0, -1, 0)
	DateTo := time.Now()

	for code, CourtName := range Codes {
		fmt.Println(CourtName)
		meetPages, ErrPages := mos.Pages(DateFrom, DateTo, code)
		if ErrPages != nil {
			panic(ErrPages)
		}
		meets = append(meets, meetPages...)
	}

	fmt.Println(len(meets))
	gocouter.SaveXlsx("mos.xlsx", meets)
}

func MO() {
	// загрузка данных
	cr, ErrNewCouter := gocouter.NewCouter("PostFix")
	if ErrNewCouter != nil {
		panic(ErrNewCouter)
	}
	pg, ErrLoadPG := gocouter.LoadConfig("2captcha.json")
	if ErrLoadPG != nil {
		panic(ErrLoadPG)
	}
	cr.Ch = pg

	MapCouter, ErrInput := gocouter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		panic(ErrInput)
	}
	if len(MapCouter) == 0 {
		panic("Lenght of MapCouter = 0")
	}

	var mts []gocouter.Meeting
	for Region, mCouter := range MapCouter {
		// if Region != "Город Санкт-Петербург" {
		// 	continue
		// }

		var RegionMeets []gocouter.Meeting
		for nROI := range mCouter {

			// nROI := 0

			fmt.Println("Исследуем регион:", Region)
			fmt.Println("Исследуем суд:", mCouter[nROI].Name, mCouter[nROI].URL)

			meets, ErrPages := cr.Pages(mCouter[nROI].URL)
			if ErrPages != nil {
				fmt.Println(ErrPages)
				continue
			}
			// fmt.Println("Всего судебных дел в суде:", len(meets))
			for imeet := range meets {
				fmt.Println(imeet, "/", len(meets))
				meets[imeet].Link = mCouter[nROI].URL + meets[imeet].Link
				cs, ErrCase := cr.ParseCase(mCouter[nROI].URL + meets[imeet].Link)
				if ErrCase != nil {
					fmt.Println(ErrCase)
					continue
				}
				meets[imeet].Case = cs
				time.Sleep(time.Millisecond * 100)
			}
			fmt.Println("Всего судебных дел в суде", mCouter[nROI].Name, ":", len(meets))

			RegionMeets = append(RegionMeets, meets...)
		}
		fmt.Println("Всего судебных дел в Регионе", RegionMeets, ":", len(RegionMeets))

		mts = append(mts, RegionMeets...)
	}

	fmt.Println(len(mts))
	gocouter.SaveXlsx("mo.xlsx", mts)
}

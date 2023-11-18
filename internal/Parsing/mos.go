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

	for code := range Codes {
		fmt.Println(code)
		meetPages, ErrPages := mos.Pages(DateFrom, DateTo, code)
		if ErrPages != nil {
			panic(ErrPages)
		}
		meets = append(meets, meetPages...)
	}

	fmt.Println(len(meets))
	gocouter.SaveXlsx("mos.xlsx", meets)
}

package region

import (
	"fmt"
	"testing"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
)

// https://himki--mo.sudrf.ru/modules.php?name=sud_delo&srv_num=1&H_date=11.12.2023
func TestPage(t *testing.T) {
	urlcouter := "https://himki--mo.sudrf.ru"
	date := time.Date(2023, time.December, 11, 0, 0, 0, 0, time.Local)
	meets, ErrPage := Page(urlcouter, date)
	if ErrPage != nil {
		t.Error(ErrPage)
	}
	for imeet, meet := range meets {
		fmt.Printf("%d. %+v\n", imeet, meet)
	}
	fmt.Println(len(meets))
	fmt.Println(len(TrudFilter(meets)))

	gocouter.SaveXlsx("TestPage.xlsx", meets)
}

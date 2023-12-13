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

func TestParseCase(t *testing.T) {
	cases, Errcase := ParseCase("https://abinsk--krd.sudrf.ru/modules.php?name=sud_delo&srv_num=1&name_op=case&case_id=306972050&case_uid=9543662b-d2d9-444f-a8a2-d742c1c1c441&result=0&delo_id=1540005&new=")
	if Errcase != nil {
		t.Error(Errcase)
	}
	fmt.Println(cases)
}

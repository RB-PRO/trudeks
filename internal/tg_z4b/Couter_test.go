package tgz4b

import (
	"fmt"
	"sync"
	"testing"

	couter "trudeks/pkg/go-couter"
)

func TestCouter(t *testing.T) {
	// загрузка данных
	cr, ErrNewCouter := couter.NewCouter("..\\..\\PostFix")
	if ErrNewCouter != nil {
		t.Error("couter.NewCouter:", ErrNewCouter)
	}
	pg, ErrLoadPG := couter.LoadConfig("..\\..\\2captcha.json")
	if ErrLoadPG != nil {
		t.Error("couter.LoadConfig:", ErrLoadPG)
	}
	cr.Ch = pg

	MapCouter, ErrInput := couter.InputFileXlsxCouter("..\\..\\суды.xlsx")
	if ErrInput != nil {
		t.Error("couter.InputFileXlsxCouter:", ErrInput)
	}
	if len(MapCouter) == 0 {
		t.Error("couter.InputFileXlsxCouter: Lenght of MapCouter = 0")
	}

	//
	Region := ""
	var Couters []couter.CouterNameLink
	var i int
	for r, c := range MapCouter {
		if i == 17 {
			Region = r
			Couters = c
			break
		}
		i++
	}
	fmt.Println(Region, Couters)

	var MEETS []couter.Meeting
	var wg sync.WaitGroup
	ch := make(chan []couter.Meeting, 1)
	wg.Add(1)
	go cr.ParseRegion(Region, Couters[:1], ch, &wg)
	wg.Wait()
	close(ch)
	MEETS = <-ch
	couter.SaveXlsx("test.xlsx", MEETS)
}

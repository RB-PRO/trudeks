package mos

import (
	"fmt"
	"testing"
	"time"

	gocouter "github.com/RB-PRO/trudeks/pkg/go-couter"
)

func TestPage(t *testing.T) {
	url := "https://mos-gorsud.ru/search?caseDateFrom=01.11.2023&caseDateTo=18.11.2023&category=a8fa044&processType=2&formType=fullForm&page=1"
	meets, next, err := Page(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Next is", next)
	fmt.Println("len(meets)", len(meets))
	fmt.Printf("%+v\n", meets[0])
	gocouter.SaveXlsx("test.xlsx", meets)
}

func TestPages(t *testing.T) {

	DateFrom := time.Now().AddDate(0, -1, 0)
	DateTo := time.Now()
	meets, err := Pages(DateFrom, DateTo, "a8fa044")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("len(meets)", len(meets))
	fmt.Printf("%+v\n", meets[0])
}
func TestParseMeet(t *testing.T) {
	url := "https://mos-gorsud.ru/mgs/services/cases/appeal-civil/details/0e864560-62b4-11ee-b474-fdabb39ff4a4?caseDateFrom=18.10.2022&caseDateTo=18.11.2023&category=68d93bc&formType=fullForm"
	meet, err := ParseMeet(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", meet)
}

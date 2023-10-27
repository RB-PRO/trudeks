package couterparser

import (
	"fmt"
	"log"
	"testing"

	api2captcha "github.com/2captcha/2captcha-go"
)

func TestPages(t *testing.T) {
	pg, ErrLoadPG := LoadConfig("../../2captcha.json")
	if ErrLoadPG != nil {
		t.Error(ErrLoadPG)
	}

	meets, ErrPage := pg.Pages("https://himki--mo.sudrf.ru")
	if ErrPage != nil {
		t.Error(ErrPage)
	}
	fmt.Println(len(meets))
}

func TestPage(t *testing.T) {
	url := `https://himki--mo.sudrf.ru/modules.php?name=sud_delo&srv_num=1&name_op=r&page=3&vnkod=50RS0048&srv_num=1&name_op=r&vnkod=50RS0048&delo_id=1540005&case_type=0&new=0&G1_PARTS__NAMESS=&g1_case__CASE_NUMBERSS=&g1_case__JUDICIAL_UIDSS=&delo_table=g1_case&g1_case__ENTRY_DATE1D=25.10.2022&g1_case__ENTRY_DATE2D=25.10.2023&lawbookarticles%5B0%5D=%D1%EF%EE%F0%FB%2C+%E2%EE%E7%ED%E8%EA%E0%FE%F9%E8%E5+%E8%E7+%F2%F0%F3%E4%EE%E2%FB%F5+%EE%F2%ED%EE%F8%E5%ED%E8%E9&G1_CASE__JUDGE=&g1_case__RESULT_DATE1D=&g1_case__RESULT_DATE2D=&G1_CASE__RESULT=&G1_CASE__BUILDING_ID=&G1_CASE__COURT_STRUCT=&G1_EVENT__EVENT_NAME=&G1_EVENT__EVENT_DATEDD=&G1_PARTS__PARTS_TYPE=&G1_PARTS__INN_STRSS=&G1_PARTS__KPP_STRSS=&G1_PARTS__OGRN_STRSS=&G1_PARTS__OGRNIP_STRSS=&G1_RKN_ACCESS_RESTRICTION__RKN_REASON=&g1_rkn_access_restriction__RKN_RESTRICT_URLSS=&g1_requirement__ACCESSION_DATE1D=&g1_requirement__ACCESSION_DATE2D=&G1_REQUIREMENT__CATEGORY=&g1_requirement__ESSENCESS=&g1_requirement__JOIN_END_DATE1D=&g1_requirement__JOIN_END_DATE2D=&G1_REQUIREMENT__PUBLICATION_ID=&G1_DOCUMENT__PUBL_DATE1D=&G1_DOCUMENT__PUBL_DATE2D=&G1_CASE__VALIDITY_DATE1D=&G1_CASE__VALIDITY_DATE2D=&G1_ORDER_INFO__ORDER_DATE1D=&G1_ORDER_INFO__ORDER_DATE2D=&G1_ORDER_INFO__ORDER_NUMSS=&G1_ORDER_INFO__EXTERNALKEYSS=&G1_ORDER_INFO__STATE_ID=&G1_ORDER_INFO__RECIP_ID=&Submit=%CD%E0%E9%F2%E8`
	meets, Next, ErrPage := Page(url)
	if ErrPage != nil {
		t.Error(ErrPage)
	}
	fmt.Println("Существование следующей страницы:", Next)
	fmt.Printf("%+v\n", meets[0])
}

func TestParseCase(t *testing.T) {
	cs, err := ParseCase(`https://himki--mo.sudrf.ru/modules.php?name=sud_delo&srv_num=1&name_op=case&case_id=652309029&case_uid=059dec2c-6f10-445b-9558-2b8bdada1ba9&delo_id=1540005`)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", cs)
}

func TestComplexParse(t *testing.T) {
	MapCouter, ErrInput := InputFileXlsxCouter("../../суды.xlsx")
	if ErrInput != nil {
		t.Error(ErrInput)
	}
	if len(MapCouter) == 0 {
		t.Error("Lenght of MapCouter = 0")
	}

	pg, ErrLoadPG := LoadConfig("../../2captcha.json")
	if ErrLoadPG != nil {
		t.Error(ErrLoadPG)
	}

	for Region, mCouter := range MapCouter {
		if Region != "Город Санкт-Петербург" {
			continue
		}
		nROI := 0
		fmt.Println("Исследуем регион:", Region)
		fmt.Println("Исследуем суд:", mCouter[nROI].Name, mCouter[nROI].URL)

		meets, ErrPages := pg.Pages(mCouter[nROI].URL)
		if ErrPages != nil {
			t.Error(ErrPages)
		}
		fmt.Println("Всего судебных дел в суде:", len(meets))
		for imeet := range meets {
			fmt.Println(imeet, "/", len(meets))
			meets[imeet].Link = mCouter[nROI].URL + meets[imeet].Link
			cs, ErrCase := ParseCase(mCouter[nROI].URL + meets[imeet].Link)
			if ErrCase != nil {
				t.Error(ErrCase)
			}
			meets[imeet].Case = cs
		}
		fmt.Println("Всего судебных дел в суде:", len(meets))

		break
	}

}

func TestCaptchaParse(t *testing.T) {
	_, IDcaptcha, ErrCaptchaParse := CaptchaParse("https://vbr--spb.sudrf.ru")
	if ErrCaptchaParse != nil {
		t.Error(ErrCaptchaParse)
	}
	fmt.Println("ID капчи:", IDcaptcha)

	pg, ErrLoadPG := LoadConfig("../../2captcha.json")
	if ErrLoadPG != nil {
		t.Error(ErrLoadPG)
	}

	// Получаем капчу
	base64captcha, IDcaptcha, Err := CaptchaParse("https://vbr--spb.sudrf.ru")
	if Err != nil {
		t.Error(Err)
	}

	// Стртуктура запросника
	cap := api2captcha.Normal{
		Base64:   base64captcha, // base64 картинки
		Numberic: 1,             // Только цифры
		MinLen:   5,             // Минимальное к-во символов
		MaxLen:   5,             // Максимальное к-во символов
		Lang:     "en",
		HintText: "5 digits, lines on top of the picture",
	}
	code, err := pg.CaptchaClient.Solve(cap.ToRequest())
	if err != nil {
		if err == api2captcha.ErrTimeout {
			log.Println("ps.CaptchaClient.Solve: Timeout")
		} else if err == api2captcha.ErrApi {
			log.Println("ps.CaptchaClient.Solve: API error")
		} else if err == api2captcha.ErrNetwork {
			log.Println("ps.CaptchaClient.Solve: Network error")
		} else {
			log.Println(err)
		}
	}
	CapthaCode := fmt.Sprintf("&captcha=%s&captchaid=%s", code, IDcaptcha)
	fmt.Println("Капча:", CapthaCode)
}

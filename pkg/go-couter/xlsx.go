package couterparser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type XLSX struct {
	*excelize.File
	fileName  string
	sheetName string
	cout      int
}

func NewSavingFile(filename string) (*XLSX, error) {

	f := excelize.NewFile() // Создать книгу Excel
	sheetName := "main"     // Create a new sheet.
	_, err := f.NewSheet(sheetName)
	if err != nil {
		return nil, fmt.Errorf("f.NewSheet: %v", err)
	}
	f.DeleteSheet("Sheet1")
	f.SaveAs(filename)

	return &XLSX{File: f, fileName: filename, sheetName: sheetName, cout: 2}, nil
}

// Вписать все строки в файл и сохранить
func (xx *XLSX) InputRows(RegionName, CouterName string, meets []Meeting) error {
	for _, meet := range meets {
		xx.inputRow(RegionName, CouterName, meet)
	}
	xx.Save()
	return nil
}

// Вписать всего одну строку
func (xx *XLSX) inputRow(RegionName, CouterName string, meet Meeting) error {
	xx.setCell(xx.cout, 1, RegionName)
	xx.setCell(xx.cout, 2, CouterName)
	xx.setCell(xx.cout, 3, meet.Number)
	xx.setCell(xx.cout, 4, meet.Code)
	xx.setCell(xx.cout, 5, meet.Type)
	xx.setCell(xx.cout, 6, meet.Link)
	if !meet.DateReceipt.IsZero() {
		xx.setCell(xx.cout, 7, meet.DateReceipt.Format("02.01.2006"))
	}
	xx.setCell(xx.cout, 8, strings.Join(meet.Category, ";"))
	xx.setCell(xx.cout, 9, meet.Judge)
	if !meet.DateDone.IsZero() {
		xx.setCell(xx.cout, 10, meet.DateDone.Format("02.01.2006"))
	}
	if !meet.DateCouterProcess.IsZero() {
		xx.setCell(xx.cout, 11, meet.DateCouterProcess.Format("02.01.2006"))
	}
	// setCell(xx.cout, 10, meet.Appealed)
	xx.setCell(xx.cout, 13, meet.DoneReport)
	if !meet.DateEffective.IsZero() {
		xx.setCell(xx.cout, 14, meet.DateEffective.Format("02.01.2006"))
	}
	xx.setCell(xx.cout, 15, meet.CourtActURL)
	xx.setCell(xx.cout, 16, meet.Status)

	sort.Slice(meet.Case.Attack, func(i, j int) (less bool) {
		return meet.Case.Attack[i].INN != ""
	})
	sort.Slice(meet.Case.Defense, func(i, j int) (less bool) {
		return meet.Case.Defense[i].INN != ""
	})

	// Истец
	var Attack []string
	for _, att := range meet.Case.Attack {
		Attack = append(Attack, att.Name)
	}
	xx.setCell(xx.cout, 17, strings.Join(Attack, ";"))
	xx.setCell(xx.cout, 21, "Мосгорсуд")

	for _, def := range meet.Case.Defense {
		// xx.setColorCell(xx.cout, 18, "E0EBF5")

		xx.setCell(xx.cout, 18, def.Name)
		xx.setCell(xx.cout, 19, def.INN)
		xx.setCell(xx.cout, 20, "ИНН "+def.Name)
		xx.cout++
	}
	// // Ответчик
	// var Defense, DefenseINN []string
	// for _, def := range meet.Case.Defense {
	// 	Defense = append(Defense, def.Name)
	// 	DefenseINN = append(DefenseINN, def.INN)
	// }
	// for i := range Defense {
	// 	xx.setCell(xx.cout, 18, Defense[i])
	// 	xx.setCell(xx.cout, 19, DefenseINN[i])
	// 	xx.setCell(xx.cout, 19, "ИНН "+Defense[i])
	// 	xx.cout++
	// }

	// setCell(xx.cout, 14, Attack []Side)
	if len(meet.Case.Defense) == 0 {
		xx.cout++
	}
	return nil
}

// Название колонок
func (xx *XLSX) SetHead() error {
	xx.setHeadCol(1, "Регион")                           //
	xx.setHeadCol(2, "Название суда")                    //
	xx.setHeadCol(3, "Номер дела")                       // Number
	xx.setHeadCol(4, "Код дела")                         // Code
	xx.setHeadCol(5, "Тип дела")                         // Code
	xx.setHeadCol(6, "Ссылка на дело")                   // Link
	xx.setHeadCol(7, "Дата поступления дела")            // DateReceipt
	xx.setHeadCol(8, "Категория спора")                  // Category
	xx.setHeadCol(9, "Судья")                            // Judge
	xx.setHeadCol(10, "Дата решения")                    // DateDone
	xx.setHeadCol(11, "Дата судебного события")          //
	xx.setHeadCol(12, "Обжалуется или нет")              // Appealed
	xx.setHeadCol(13, "Решение")                         // DoneReport
	xx.setHeadCol(14, "Дата вступления в законную силу") // DateEffective
	xx.setHeadCol(15, "Судебные акты(Ссылка)")           // CourtActURL
	xx.setHeadCol(16, "Статус")                          // CourtActURL
	xx.setHeadCol(17, "Истец")                           // CourtActURL
	xx.setHeadCol(18, "Название компании")               // Ответчик
	xx.setHeadCol(19, "ИНН")                             // CourtActURL
	xx.setHeadCol(20, "link")                            // CourtActURL
	xx.setHeadCol(21, "Источник")
	xx.setHeadCol(22, "Ответственный")
	return nil
}

func (xx *XLSX) setCell(y, x int, value interface{}) error {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	return xx.SetCellValue(xx.sheetName, collumnStr+strconv.Itoa(y), value)
}

// Вписать значение в название колонки
func (xx *XLSX) setHeadCol(x int, value interface{}) error {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	return xx.SetCellValue(xx.sheetName, collumnStr+"1", value)
}

func (xx *XLSX) setColorCell(y, x int, valueRGB string) error {

	style, ErrNewStyle := xx.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{valueRGB}, Pattern: 1},
	})
	if ErrNewStyle != nil {
		return fmt.Errorf("ErrNewStyle: %v", ErrNewStyle)
	}

	collumnStr, ErrColName := excelize.ColumnNumberToName(x)
	if ErrColName != nil {
		return fmt.Errorf("ErrColName: %v", ErrColName)
	}
	cellsStr := collumnStr + strconv.Itoa(y)

	return xx.SetCellStyle(xx.sheetName, cellsStr, cellsStr, style)
}

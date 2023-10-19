package zachestnyibiznes

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/xuri/excelize/v2"
)

type XLSX struct {
	f         *excelize.File
	line      int
	SheetName string
	colNumber map[string]int
}

// Создать книгу
func NewXLSX(PathNameFile string) *XLSX {
	f := excelize.NewFile()
	// defer f.Close()

	SheetName := "main"
	f.NewSheet(SheetName)

	f.SetCellValue(SheetName, "A1", "Запрос") // Запрос, по которому проводился поиск

	t := reflect.TypeOf(DocsItem{})
	// v := reflect.ValueOf(DocsItem{})
	xlsx := &XLSX{
		f:         f,
		line:      2,
		SheetName: SheetName,
		colNumber: make(map[string]int),
	}
	for i := 0; i < t.NumField(); i++ {
		// fmt.Println(t.Field(i).Name, t.Field(i).Tag.Get("json"), v.Field(i))
		CollumnName := strings.ReplaceAll(t.Field(i).Tag.Get("json"), ",omitempty", "") // Название колонки
		SymbolCol, _ := excelize.ColumnNumberToName(i + 2)                              // Символ текущей колонки в Excel
		xlsx.colNumber[CollumnName] = i + 2
		xlsx.f.SetCellValue(SheetName, fmt.Sprintf("%s%d", SymbolCol, 1), CollumnName)
	}

	xlsx.f.DeleteSheet("Sheet1")
	xlsx.f.SaveAs(PathNameFile)
	return xlsx
}

// Вписать данные по товару в книгу
func (x *XLSX) WriteXLSX(ID string, Products Сontacts) {
	// Записать данные
	x.f.SetCellValue(x.SheetName, fmt.Sprintf("%s%d", "A", x.line), ID)
	if len(Products.Body.Docs) > 0 {
		t := reflect.TypeOf(Products.Body.Docs[0])
		value := reflect.ValueOf(Products.Body.Docs[0])
		for i := 0; i < value.NumField(); i++ {
			CollumnName := strings.ReplaceAll(t.Field(i).Tag.Get("json"), ",omitempty", "") // Название колонки

			SymbolCol, _ := excelize.ColumnNumberToName(x.colNumber[CollumnName]) // Символ текущей колонки в Excel
			x.f.SetCellValue(x.SheetName, fmt.Sprintf("%s%d", SymbolCol, x.line), value.Field(i))
		}
		x.f.Save()
		x.line++ // Иттерирование по строкам
	}
}

// Закрыть и сохранить файл
func (x *XLSX) CloceAndSaveXLSX() {
	x.f.Save()
	x.f.Close()
}

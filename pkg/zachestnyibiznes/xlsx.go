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
	for i := 0; i < t.NumField(); i++ {
		// fmt.Println(t.Field(i).Name, t.Field(i).Tag.Get("json"), v.Field(i))
		CollumnName := strings.ReplaceAll(t.Field(i).Tag.Get("json"), ",omitempty", "") // Название колонки
		SymbolCol, _ := excelize.ColumnNumberToName(i + 2)                              // Символ текущей колонки в Excel
		f.SetCellValue(SheetName, fmt.Sprintf("%s%d", SymbolCol, 1), CollumnName)
	}

	//////////////
	// iterateStructFields(DocsItem{})

	f.DeleteSheet("Sheet1")
	f.SaveAs(PathNameFile)
	return &XLSX{f: f, line: 2, SheetName: SheetName}
}

// Вписать данные по товару в книгу
func (x *XLSX) WriteXLSX(ID string, Products Сontacts) {
	// Записать данные
	// t := reflect.TypeOf(DocsItem{})
	x.f.SetCellValue(x.SheetName, fmt.Sprintf("%s%d", "A", x.line), ID)
	value := reflect.ValueOf(Products.Body.Docs[0])
	for i := 0; i < value.NumField(); i++ {
		SymbolCol, _ := excelize.ColumnNumberToName(i + 2) // Символ текущей колонки в Excel
		x.f.SetCellValue(x.SheetName, fmt.Sprintf("%s%d", SymbolCol, x.line), value.Field(i))
	}
	x.f.Save()
	x.line++ // Иттерирование по строкам
}

// Закрыть и сохранить файл
func (x *XLSX) CloceAndSaveXLSX() {
	x.f.Save()
	x.f.Close()
}

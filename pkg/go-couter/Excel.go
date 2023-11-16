package couterparser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Сохранить данные в Xlsx
func SaveXlsx(FileName string, Meetings []Meeting) error {
	// Создать книгу
	book, makeBookError := MakeWorkBook()
	if makeBookError != nil {
		return makeBookError
	}

	wotkSheet := "main"

	setHead(book, wotkSheet, 1, "Номер дела")                       // Number
	setHead(book, wotkSheet, 2, "Код дела")                         // Code
	setHead(book, wotkSheet, 3, "Ссылка на дело")                   // Link
	setHead(book, wotkSheet, 4, "Дата поступления дела")            // DateReceipt
	setHead(book, wotkSheet, 5, "Категория")                        // Category
	setHead(book, wotkSheet, 6, "Судья")                            // Judge
	setHead(book, wotkSheet, 7, "Дата решения")                     // DateDone
	setHead(book, wotkSheet, 8, "Обжалуется или нет")               // Appealed
	setHead(book, wotkSheet, 9, "Решение")                          // DoneReport
	setHead(book, wotkSheet, 10, "Дата вступления в законную силу") // DateEffective
	setHead(book, wotkSheet, 11, "Судебные акты(Ссылка)")           // CourtActURL

	for indexItem, valItem := range Meetings {
		row := indexItem + 2
		setCell(book, wotkSheet, row, 1, valItem.Number)
		setCell(book, wotkSheet, row, 2, valItem.Code)
		setCell(book, wotkSheet, row, 3, valItem.Link)
		setCell(book, wotkSheet, row, 4, valItem.DateReceipt)
		setCell(book, wotkSheet, row, 5, strings.Join(valItem.Category, ";"))
		setCell(book, wotkSheet, row, 6, valItem.Judge)
		setCell(book, wotkSheet, row, 7, valItem.DateDone)
		setCell(book, wotkSheet, row, 8, valItem.Appealed)
		setCell(book, wotkSheet, row, 9, valItem.DoneReport)
		setCell(book, wotkSheet, row, 10, valItem.DateEffective)
		setCell(book, wotkSheet, row, 11, valItem.CourtActURL)

		setCell(book, wotkSheet, row, 12, valItem.Case.Idntifier)
		setCell(book, wotkSheet, row, 13, valItem.Case.IdntifierLink)

		fmt.Printf("A: %+v\n", valItem.Case.Attack)
		sort.Slice(valItem.Case.Attack, func(i, j int) (less bool) {
			return valItem.Case.Attack[i].INN != ""
		})
		fmt.Printf("A: %+v\n", valItem.Case.Attack)
		// setCell(book, wotkSheet, row, 14, Attack []Side)

	}

	// Закрыть книгу
	closeBookError := CloseXlsx(book, FileName)
	if closeBookError != nil {
		return closeBookError
	}

	return nil
}

// Создать книгу
func MakeWorkBook() (*excelize.File, error) {
	// Создать книгу Excel
	f := excelize.NewFile()
	// Create a new sheet.
	_, err := f.NewSheet("main")
	if err != nil {
		return f, err
	}
	f.DeleteSheet("Sheet1")
	return f, nil
}

// Сохранить и закрыть файл
func CloseXlsx(f *excelize.File, filename string) error {
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// Вписать значение в ячейку
func setCell(file *excelize.File, wotkSheet string, y, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+strconv.Itoa(y), value)
}

// Вписать значение в название колонки
func setHead(file *excelize.File, wotkSheet string, x int, value interface{}) {
	collumnStr, _ := excelize.ColumnNumberToName(x)
	file.SetCellValue(wotkSheet, collumnStr+"1", value)
}

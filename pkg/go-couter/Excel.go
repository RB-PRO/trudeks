package couterparser

import (
	"sort"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Сохранить данные в Xlsx
func SaveXlsx(FileName string, Meetings []Meeting) error {
	book, makeBookError := MakeWorkBook() // Создать книгу
	if makeBookError != nil {
		return makeBookError
	}
	wotkSheet := "main"
	setHead(book, wotkSheet, 1, "Номер дела")                       // Number
	setHead(book, wotkSheet, 2, "Код дела")                         // Code
	setHead(book, wotkSheet, 3, "Тип дела")                         // Code
	setHead(book, wotkSheet, 4, "Ссылка на дело")                   // Link
	setHead(book, wotkSheet, 5, "Дата поступления дела")            // DateReceipt
	setHead(book, wotkSheet, 6, "Категория спора")                  // Category
	setHead(book, wotkSheet, 7, "Судья")                            // Judge
	setHead(book, wotkSheet, 8, "Дата решения")                     // DateDone
	setHead(book, wotkSheet, 9, "Дата судебного события")           //
	setHead(book, wotkSheet, 10, "Обжалуется или нет")              // Appealed
	setHead(book, wotkSheet, 11, "Решение")                         // DoneReport
	setHead(book, wotkSheet, 12, "Дата вступления в законную силу") // DateEffective
	setHead(book, wotkSheet, 13, "Судебные акты(Ссылка)")           // CourtActURL
	setHead(book, wotkSheet, 14, "Статус")                          // CourtActURL
	setHead(book, wotkSheet, 15, "Истец")                           // CourtActURL
	setHead(book, wotkSheet, 16, "Название компании")               // Ответчик
	setHead(book, wotkSheet, 17, "ИНН")                             // CourtActURL
	setHead(book, wotkSheet, 18, "link")                            // CourtActURL
	setHead(book, wotkSheet, 19, "Источник")
	setHead(book, wotkSheet, 20, "Ответственный")

	var row int = 2

	for _, valItem := range Meetings {
		// row = indexItem + 2
		setCell(book, wotkSheet, row, 1, valItem.Number)
		setCell(book, wotkSheet, row, 2, valItem.Code)
		setCell(book, wotkSheet, row, 3, valItem.Type)
		setCell(book, wotkSheet, row, 4, valItem.Link)
		if !valItem.DateReceipt.IsZero() {
			setCell(book, wotkSheet, row, 5, valItem.DateReceipt.Format("02.01.2006"))
		}
		setCell(book, wotkSheet, row, 6, strings.Join(valItem.Category, ";"))
		setCell(book, wotkSheet, row, 7, valItem.Judge)
		if !valItem.DateDone.IsZero() {
			setCell(book, wotkSheet, row, 8, valItem.DateDone.Format("02.01.2006"))
		}
		if !valItem.DateCouterProcess.IsZero() {
			setCell(book, wotkSheet, row, 9, valItem.DateCouterProcess.Format("02.01.2006"))
		}
		// setCell(book, wotkSheet, row, 8, valItem.Appealed)
		setCell(book, wotkSheet, row, 11, valItem.DoneReport)
		if !valItem.DateEffective.IsZero() {
			setCell(book, wotkSheet, row, 12, valItem.DateEffective.Format("02.01.2006"))
		}
		setCell(book, wotkSheet, row, 13, valItem.CourtActURL)
		setCell(book, wotkSheet, row, 14, valItem.Status)

		// setCell(book, wotkSheet, row, 12, valItem.Case.Idntifier)
		// setCell(book, wotkSheet, row, 13, valItem.Case.IdntifierLink)

		// fmt.Printf("A: %+v\n", valItem.Case.Attack)
		sort.Slice(valItem.Case.Attack, func(i, j int) (less bool) {
			return valItem.Case.Attack[i].INN != ""
		})
		sort.Slice(valItem.Case.Defense, func(i, j int) (less bool) {
			return valItem.Case.Defense[i].INN != ""
		})
		// fmt.Printf("A: %+v\n", valItem.Case.Attack)

		// Истец
		var Attack []string
		for _, att := range valItem.Case.Attack {
			Attack = append(Attack, att.Name)
		}
		setCell(book, wotkSheet, row, 15, strings.Join(Attack, ";"))
		setCell(book, wotkSheet, row, 19, "Мосгорсуд")

		// Ответчик
		var Defense, DefenseINN []string
		for _, def := range valItem.Case.Defense {
			Defense = append(Defense, def.Name)
			DefenseINN = append(DefenseINN, def.INN)
		}
		for i := range Defense {
			setCell(book, wotkSheet, row, 16, Defense[i])
			setCell(book, wotkSheet, row, 17, DefenseINN[i])
			setCell(book, wotkSheet, row, 18, "ИНН "+Defense[i])
			row++
		}

		// setCell(book, wotkSheet, row, 14, Attack []Side)
		if len(Defense) == 0 {
			row++
		}
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

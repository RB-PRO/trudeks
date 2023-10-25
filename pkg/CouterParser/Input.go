package couterparser

import (
	"github.com/xuri/excelize/v2"
)

// Структура суда из файла источника
type CouterNameLink struct {
	Name string // Название суда
	URL  string // Ссылка на суд
}

// Загрузить файл и получить список судов из файла судов
//
//	MapCouter, ErrInput := InputFileXlsxCouter("суды.xlsx")
func InputFileXlsxCouter(FileNamePath string) (map[string][]CouterNameLink, error) {

	// Открываем файл
	f, ErrOpenFile := excelize.OpenFile(FileNamePath)
	if ErrOpenFile != nil {
		return nil, ErrOpenFile
	}
	defer f.Close()

	// Мапа, где в качестве ключа - регион суда
	CoutersMap := make(map[string][]CouterNameLink)

	// Получить все строки на первом листе
	rows, ErrGetRows := f.GetRows(f.GetSheetName(0))
	if ErrGetRows != nil {
		return nil, ErrGetRows
	}

	// Цикл по всем строкам
	for iRows := 1; iRows < len(rows); iRows++ {
		row := rows[iRows] // Текущая строка
		if len(row) >= 3 { // Если к-во колонок больше или равно 3s
			CoutersMap[row[2]] = append(CoutersMap[row[2]], CouterNameLink{Name: row[1], URL: row[0]})
		}
	}

	return CoutersMap, nil
}

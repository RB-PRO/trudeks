package tgz4b

import (
	"fmt"
	"sync"
	"time"

	couter "github.com/RB-PRO/trudeks/pkg/go-couter"
)

// Спарсить суды по списку
func ParsingCounterGoRoutines() (FileName string, Err error) {
	// загрузка данных
	cr, ErrNewCouter := couter.NewCouter("PostFix")
	if ErrNewCouter != nil {
		return "", fmt.Errorf("couter.NewCouter: %w", ErrNewCouter)
	}
	pg, ErrLoadPG := couter.LoadConfig("2captcha.json")
	if ErrLoadPG != nil {
		return "", fmt.Errorf("couter.LoadConfig: %w", ErrLoadPG)
	}
	cr.Ch = pg

	MapCouter, ErrInput := couter.InputFileXlsxCouter("суды.xlsx")
	if ErrInput != nil {
		return "", fmt.Errorf("couter.InputFileXlsxCouter: %w", ErrInput)
	}
	if len(MapCouter) == 0 {
		return "", fmt.Errorf("couter.InputFileXlsxCouter: Lenght of MapCouter = 0")
	}

	var wg sync.WaitGroup             // Объект ожтдания отработки парсеров
	ch := make(chan []couter.Meeting) // Канал для общения горутин и сохранения результатов
	var MEETS []couter.Meeting        // Слайс со всеми найденными судебными делами
	var i int
	for Region, Couters := range MapCouter { // Цикл по всем регионам Регионы
		if i == 0 {
			i++
			continue
		}
		fmt.Println("Исследуем регион", Region)

		wg.Add(1) // Добавляем счётчик горутин

		// Запустить горутину парсинга региона
		go cr.ParseRegion(Region, Couters[:2], ch, &wg)

		i++
		if i == 2 {
			break
		}
	}

	wg.Wait()

	for result := range ch {
		MEETS = append(MEETS, result...)
	}
	fmt.Println("Всего найдено судебных дел:", len(MEETS))

	// Название файла
	FileName = fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02_15-04"))

	// Сохрание данных в файл
	Err = couter.SaveXlsx(FileName, MEETS)

	return FileName, Err
}

package couterparser

import (
	"fmt"
	"sync"
)

// Спасить данные по всему региону подряд по каждому суду
func (cr *Couter) ParseRegion(Region string, mCouter []CouterNameLink, cn chan []Meeting, wg *sync.WaitGroup) { //(meets []Meeting, Err error) {

	var MeetsRegion []Meeting
	for nROI := range mCouter { // Суды региона
		fmt.Printf("%+v\n\n\n", mCouter[nROI])

		fmt.Printf("Исследуем суд %s по ссылке %s\n", mCouter[nROI].Name, mCouter[nROI].URL)

		// Парсинг данныен судов
		MeetsCouter, ErrParseCouter := cr.ParseCouter(mCouter[nROI].URL)
		if ErrParseCouter != nil {
			//return nil, fmt.Errorf("ParseCouter: %w for url couter %s", ErrParseCouter, mCouter[nROI].URL)
			fmt.Printf("ParseCouter: %s for url couter %s\n", ErrParseCouter.Error(), mCouter[nROI].URL)
			continue
		}

		MeetsRegion = append(MeetsRegion, MeetsCouter...)
	}
	cn <- MeetsRegion
	wg.Done()
}

// спарсить все судебные дела по определённому суду
func (cr *Couter) ParseCouter(url string) ([]Meeting, error) {

	fmt.Println()
	fmt.Println()
	fmt.Println("meets[imeet].Link", url)
	fmt.Println()
	fmt.Println()
	//  Получаем список всех судебных дел
	meets, ErrPages := cr.Pages(url)
	if ErrPages != nil {
		return nil, fmt.Errorf("couter.Pages: %w for url couter %s", ErrPages, url)
	}
	fmt.Printf("Всего найдено %d судебных дел в суде по ссылке %s\n", len(meets), url)

	// Подробности по каждому судебному деву
	for imeet := range meets {
		fmt.Println(imeet, "/", len(meets))
		fmt.Println()
		fmt.Println()
		fmt.Println("meets[imeet].Link", url, meets[imeet].Link)
		fmt.Println()
		fmt.Println()
		meets[imeet].Link = url + meets[imeet].Link
		cs, ErrCase := cr.ParseCase(meets[imeet].Link)
		if ErrCase != nil {
			return nil, fmt.Errorf("couter.ParseCase: %w for url couter %s", ErrCase, url+meets[imeet].Link)
		}
		meets[imeet].Case = cs
	}

	return meets, nil
}

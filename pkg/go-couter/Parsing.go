package couterparser

import "fmt"

// Спасить данные по всему региону подряд по каждому суду
func (cr *Couter) ParseRegion(Region string, mCouter []CouterNameLink, cn chan<- []Meeting) { //(meets []Meeting, Err error) {
	var meets []Meeting
	for nROI := 0; nROI < len(mCouter); nROI++ { // Суды региона

		fmt.Println("Исследуем суд:", mCouter[nROI].Name, mCouter[nROI].URL)

		//  Получаем список всех судебных дел
		meets, ErrPages := cr.Pages(mCouter[nROI].URL)
		if ErrPages != nil {
			// return nil, fmt.Errorf("couter.Pages: %w for url couter %s", ErrPages, mCouter[nROI].URL)
		}
		fmt.Println("Всего судебных дел в суде", mCouter[nROI].Name, ":", len(meets))

		// Подробности по каждому судебному деву
		for imeet := range meets {
			fmt.Println(imeet, "/", len(meets))
			meets[imeet].Link = mCouter[nROI].URL + meets[imeet].Link
			cs, ErrCase := cr.ParseCase(mCouter[nROI].URL + meets[imeet].Link)
			if ErrCase != nil {
				// return nil, fmt.Errorf("couter.ParseCase: %w for url couter %s", ErrCase, mCouter[nROI].URL+meets[imeet].Link)
			}
			meets[imeet].Case = cs
		}
	}
	cn <- meets
	// return meets, Err
}

package couterparser

import "time"

// Судебное дело
type Meeting struct {
	Number        string    // Номер дела
	Code          string    // Код дела - Номер дела в суде нижестоящей инстанции
	Link          string    // Ссылка на дело
	CouterName    string    // Название суда
	DateReceipt   time.Time // Дата поступления дела
	Category      []string  // Категория
	Judge         string    // Судья
	DateDone      time.Time // Дата решения
	Appealed      bool      // Обжалуется или нет
	DoneReport    string    // Решение
	DateEffective time.Time // Дата вступления в законную силу
	CourtActURL   string    // Судебные акты(Ссылка)
	Case          Case      // Содержимое по каждому делу

	// Тип дела:
	//	- Гражданские дела - первая инстанция
	//	- Гражданские дела - апелляция
	//	- Административные дела (КАC РФ) - первая инстанция
	//	- Дела об административных правонарушениях - первая инстанция
	//	- Дела об административных правонарушениях - жалобы на постановления
	//	- Уголовные дела - первая инстанция
	//	- Производство по материалам
	Type              string
	DateCouterProcess time.Time // Дата суздебного события

	Status string // Текущее состояние
	// DateApp time.Time // Дата поступления дела в апелляционную инстанцию
}

// Информация, которая получается из карточки судебного дела
type Case struct {
	Idntifier     string // Уникальный идентификатор дела
	IdntifierLink string // Ссылка на уникальный идентификатор дела
	Attack        []Side // ИСТЕЦ(ЗАЯВИТЕЛЬ)
	Defense       []Side // ОТВЕТЧИК
}

// Сторона дела
type Side struct {
	Name string // Имя, название компании стороны
	INN  string // ИНН стороны (Скорее всего не будет заполнено)
	KPP  string // КПП (Скорее всего не будет заполнено)
	OGRN string // ОГРН (Скорее всего не будет заполнено)
}

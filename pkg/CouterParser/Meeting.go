package couterparser

import "time"

// Судебное дело
type Meeting struct {
	Number        string    // Номер дела
	Code          string    // Код дела
	Link          string    // Ссылка на дело
	DateReceipt   time.Time // Дата поступления дела
	Category      []string  // Категория
	Attack        []Side    // ИСТЕЦ(ЗАЯВИТЕЛЬ)
	Defense       []Side    // ОТВЕТЧИК
	Judge         string    // Судья
	DateDone      time.Time // Дата решения
	Appealed      bool      // Обжалуется или нет
	DoneReport    string    // Решение
	DateEffective time.Time // Дата вступления в законную силу
	CourtActURL   string    // Судебные акты(Ссылка)
}

// Сторона дела
type Side struct {
	Name string // Имя, название компании стороны
	INN  string // ИНН стороны(Скорее всего не будет заполнено)
}

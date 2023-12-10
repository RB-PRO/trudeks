package tgcouter

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cron "github.com/robfig/cron/v3"
)

// Отправить уведомление админам
func (TG *Telegram) Crones() error {

	// https://russianblogs.com/article/50842746150/
	c := cron.New()

	// Выступать один раз в день в 18 часов в воскресенье
	c.AddFunc("0 0 18 * * 0,1,2,3,4", func() {
		TG.Send(tgbotapi.NewMessage(TG.ChatNotificationID, "Начинаю парсить всю Московскую область"))
		filename, Err := ParseMO()
		if Err != nil {
			msg := tgbotapi.NewMessage(TG.ChatNotificationID, "")
			msg.Text = "Ошибка в парсинге судов московской области: " + Err.Error()
			TG.Send(msg)
		} else {
			TG.Send(tgbotapi.NewDocument(TG.ChatNotificationID, tgbotapi.FilePath(filename)))
		}
	})

	c.AddFunc("0 0 0 * * 5", func() {
		TG.Send(tgbotapi.NewMessage(TG.ChatNotificationID, "Начинаю парсить всю Россию"))
		filename, Err := ParseRussia()
		if Err != nil {
			msg := tgbotapi.NewMessage(TG.ChatNotificationID, "")
			msg.Text = "Ошибка в парсинге судов России: " + Err.Error()
			TG.Send(msg)
		} else {
			TG.Send(tgbotapi.NewDocument(TG.ChatNotificationID, tgbotapi.FilePath(filename)))
		}
	})

	c.Start()
	return nil
	// msg := tgbotapi.NewMessage(TG.ChatNotificationID, str)
	// _, err := TG.Send(msg)
	// return err
}

// Отправить уведомление админам
func (TG *Telegram) Message(str string) error {
	msg := tgbotapi.NewMessage(TG.ChatNotificationID, str)
	_, err := TG.Send(msg)
	return err
}
func (TG *Telegram) Watch() error {
	TG.Crones()
	// Создайте новую структуру конфигурации обновления со смещением 0.
	// Смещения используются для того, чтобы убедиться, что Telegram знает,
	// что мы обработали предыдущие значения, и нам не нужно их повторять.
	updateConfig := tgbotapi.NewUpdate(0)

	// Сообщите Telegram, что мы должны ждать обновления до 30 секунд при каждом запросе.
	// Таким образом, мы можем получать информацию так же быстро,
	// как и при выполнении множества частых запросов,
	// без необходимости отправлять почти столько же.
	updateConfig.Timeout = 30

	// Начните опрос Telegram на предмет обновлений
	updates := TG.GetUpdatesChan(updateConfig)

	// map:=map[string]string
	// Давайте рассмотрим каждое обновление, которое мы получаем от Telegram
	for update := range updates {
		// Telegram может отправлять множество типов обновлений в зависимости от того,
		// чем занимается ваш бот.
		// Пока мы хотим просмотреть только сообщения,
		// чтобы исключить любые другие обновления
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			// Create a new MessageConfig. We don't have text yet,
			// so we leave it empty.

			// Extract the command from the Message.
			switch update.Message.Command() {
			case "ping", "ping@Geo987_bot":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = time.Now().Format("2006-01-02T15:04:05")
				TG.Send(msg)
			case "region", "region@RUcouters_bot":

				go func(MessageChatID int64, MessageID int) {
					msg := tgbotapi.NewMessage(MessageChatID, "ок, начинаю парсить всю Россию")
					msg.ReplyToMessageID = MessageID
					TG.Send(msg)
					filename, Err := ParseRussia()
					if Err != nil {
						msg := tgbotapi.NewMessage(MessageChatID, "")
						msg.Text = "Ошибка в парсинге судов России: " + Err.Error()
						msg.ReplyToMessageID = MessageID
						TG.Send(msg)
					} else {
						doc := tgbotapi.NewDocument(MessageChatID, tgbotapi.FilePath(filename))
						doc.ReplyToMessageID = MessageID
						TG.Send(doc)
					}
				}(update.Message.Chat.ID, update.Message.MessageID)

			case "mo", "mo@RUcouters_bot":
				go func(MessageChatID int64, MessageID int) {
					msg := tgbotapi.NewMessage(MessageChatID, "ок, начинаю парсить всю Московскую область")
					msg.ReplyToMessageID = MessageID
					TG.Send(msg)
					filename, Err := ParseMO()
					if Err != nil {
						msg := tgbotapi.NewMessage(MessageChatID, "")
						msg.Text = "Ошибка в парсинге судов Московской области: " + Err.Error()
						msg.ReplyToMessageID = MessageID
						TG.Send(msg)
					} else {
						doc := tgbotapi.NewDocument(MessageChatID, tgbotapi.FilePath(filename))
						doc.ReplyToMessageID = MessageID
						TG.Send(doc)
					}
				}(update.Message.Chat.ID, update.Message.MessageID)
			case "msk", "msk@RUcouters_bot":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "я не умею(((((((")
				// msg.Text = time.Now().Format("2006-01-02T15:04:05")
				msg.ReplyToMessageID = update.Message.MessageID
				TG.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = "I don't know that command"
				TG.Send(msg)
			}
			continue
		}

	}
	return nil
}

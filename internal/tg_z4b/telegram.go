package tgz4b

import (
	"log"
	"strings"
	"time"

	zachestnyibiznes "github.com/RB-PRO/trudeks/pkg/zachestnyibiznes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {

	TG_config, ErrConfigTg := LoadConfig("telegram.json")
	if ErrConfigTg != nil {
		panic(ErrConfigTg)
	}
	Z4B, ErrConfigZ4B := zachestnyibiznes.LoadConfig("zachestnyibiznes.json")
	if ErrConfigZ4B != nil {
		panic(ErrConfigZ4B)
	}

	bot, err := tgbotapi.NewBotAPI(TG_config.API_KEY_TG)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Произошла авторизация %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		log.Println(update.Message.Chat.UserName, "-", update.Message.Text, ">", update.Message.Caption)

		// Игнорируем НЕкоманды
		if !update.Message.IsCommand() {
			// Проверка наличия текста в сообщении
			if update.Message.Text == "" {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Не вижу текста.\nНужно отправить фотографию вместе с текстом."))
				continue
			}

			IDS := strings.Split(update.Message.Text, "\n")

			filename := "z4b от " + time.Now().Format("15h04m 01-02-2006") + ".xlsx"

			xlsx := zachestnyibiznes.NewXLSX(filename)
			for _, ID := range IDS {
				contacts, ErrCont := Z4B.Contacts(ID)
				if ErrCont != nil {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка:\n"+ErrCont.Error()))
				}
				xlsx.WriteXLSX(ID, contacts)
				time.Sleep(time.Millisecond * 300)
			}
			xlsx.CloceAndSaveXLSX()

			// отправляем файл
			file := tgbotapi.FilePath(filename)
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, file))

			continue
		}

		switch update.Message.Command() {
		case "example":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, `9715255412
7721294705`))

		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду\nПопробуй /start"))
			continue
		}
	}

	//

}

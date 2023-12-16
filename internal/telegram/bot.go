package telegram

import (
	"github.com/jmoiron/sqlx"
	tele "gopkg.in/telebot.v3"
	"log"
	"regexp"
	"strings"
	"time"
)

func StartTelegramBot(api string, db *sqlx.DB) {
	pref := tele.Settings{
		Token:  api,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tele.OnText, func(c tele.Context) error {
		var infoTarget string

		typeData, data := IdentifyData(c.Text())
		switch {
		case typeData == "Номер сотового телефона":
			infoTarget = GetTargetInfoToPhone(data, db)
		case typeData == "ИНН":
			infoTarget = "Метод в разработке"
		case typeData == "Номер авто":
			infoTarget = "Метод в разработке"
		case typeData == "Номер паспорта":
			infoTarget = "Метод в разработке"
		case typeData == "СНИЛС":
			infoTarget = "Метод в разработке"
		case typeData == "Штрих-код":
			infoTarget = "Метод в разработке"
		case typeData == "Электронная почта":
			infoTarget = "Метод в разработке"
		default:
			infoTarget = "Неизвестные данные"
		}
		return c.Send(infoTarget)
	})

	bot.Start()
}

func GetTargetInfoToPhone(phoneNumber string, db *sqlx.DB) string {
	//phoneRegexFone := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	wonCifr := phoneNumber[0:1]
	if len(phoneNumber) == 10 && wonCifr == "9" {
		phoneNumber = "7" + phoneNumber
	}

	var result []string
	finishedName := "Имя предпологаемое: \r\n"
	finishedAdres := "\r\nАдрес:  \r\n"

	var nameTarget, _ = GetTargetName(phoneNumber, db)
	finishedName = finishedName + nameTarget

	result = append(result, finishedName)

	adresTarget, _ := GetTargetAdres(phoneNumber, db)
	finishedAdres = finishedAdres + adresTarget
	result = append(result, finishedAdres)

	target := strings.Join(result, "\r\n ")
	return target
}
func IdentifyData(data string) (string, string) {
	phoneRegex := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	innRegex := regexp.MustCompile(`^\d{10}$`)
	carNumberRegex := regexp.MustCompile(`^[А-Я]{1}\d{3}[А-Я]{2}\d{2,3}$`)
	passportNumberRegex := regexp.MustCompile(`^\d{4}\s?\d{6}$`)
	snilsRegex := regexp.MustCompile(`^\d{3}-\d{3}-\d{3}\s?\d{2}$`)
	htRegex := regexp.MustCompile(`^[A-Za-z0-9]+$`)
	emailRegex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

	switch {
	case phoneRegex.MatchString(data):
		return "Номер сотового телефона", data
	case innRegex.MatchString(data):
		return "ИНН", data
	case carNumberRegex.MatchString(data):
		return "Номер авто", data
	case passportNumberRegex.MatchString(data):
		return "Номер паспорта", data
	case snilsRegex.MatchString(data):
		return "СНИЛС", data
	case htRegex.MatchString(data):
		return "Штрих-код", data
	case emailRegex.MatchString(data):
		return "Электронная почта", data
	default:
		return "Неизвестные данные", data
	}
}

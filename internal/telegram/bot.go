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
		lenText := strings.Count(c.Text(), "") - 1
		if c.Text()[0:1] == "7" && len(c.Text()) == 11 {
			infoTarget = GetTargetInfoToPhone(c.Text(), db)
		} else if lenText == 9 || lenText == 8 {
			infoTarget, _ = GetTargetAvtoToGosNomer(c.Text(), db)
		} else {
			infoTarget = "Мы не поняли что вы нам прислали. Номер телефона в формате 79109999999 или номеравто а111пр11"

		}
		return c.Send(infoTarget)
	})

	bot.Start()
}

func GetTargetInfoToPhone(message string, db *sqlx.DB) string {
	phoneRegexFone := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	phoneNumberFone := phoneRegexFone.FindString(message)

	var result []string
	finishedName := "Имя предпологаемое: \r\n"
	finishedAdres := "\r\nАдрес:  \r\n"

	var nameTarget, _ = GetTargetName(phoneNumberFone, db)
	finishedName = finishedName + nameTarget

	result = append(result, finishedName)

	adresTarget, _ := GetTargetAdres(phoneNumberFone, db)
	finishedAdres = finishedAdres + adresTarget
	result = append(result, finishedAdres)

	target := strings.Join(result, "\r\n ")
	return target
}

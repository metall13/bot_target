package telegram

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

func GetTargetAvtoToGosNomer(message string, db *sqlx.DB) (string, error) {

	message = strings.ToUpper(message)

	//phoneRegexpNombeAvto := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	//phoneNumberNombeAvto := phoneRegexpNombeAvto.FindString(message)

	qwery := `SELECT       
    fio,          
    place_birth, 
    date_of_birth,
    identification,
    address,       
    phone_number,
    inn,        
    gos_nomer,    
    old_gos_nomer,
    vin,          
    info         
    from public.gipdd_nn_2020 where gos_nomer = $1`

	var phone_number int64
	var fio, place_birth, date_of_birth, identification, address,
		inn, gos_nomer, old_gos_nomer, info, vin, target string
	//Counting number of users

	rows, err := db.Query(qwery, message)

	var adres []string
	for rows.Next() {
		err = rows.Scan(&fio, &date_of_birth, &place_birth, &identification, &address, &phone_number, &inn, &gos_nomer, &old_gos_nomer, &vin, &info)
		if err != nil {
			return "", err
		}
		phone_numberStr := strconv.Itoa(int(phone_number))
		adres = append(adres,
			fio+"\r\n"+place_birth+"\r\n "+identification+"\r\n"+address+"\r\n"+phone_numberStr+"\r\n"+inn+"\r\n"+gos_nomer+"\r\n"+old_gos_nomer+"\r\n"+info+"\r\n"+vin+"\r\n")
	}
	rows.Close()
	target = strings.Join(adres, "\r\n")

	if target == "" {
		target = "Нет Данных"
	}

	return target, nil
}

func GetTargetName(phoneNumber string, db *sqlx.DB) (string, error) {
	qwery := `SELECT avito_user_name from public.avito_full where phone_number = $1
	UNION ALL SELECT beeline_full_name from public.beeline_full where phone_number = $1
	UNION ALL SELECT cdek_full_name from public.cdek_full where phone_number = $1
	UNION ALL SELECT delivery2_name from public.delivery2_full where phone_number = $1
	UNION ALL SELECT delivery_name from public.delivery_full where phone_number = $1
	UNION ALL SELECT fb_full_name from public.facebook_full where phone_number = $1
	UNION ALL SELECT gibdd2_name from public.gibdd2_full where phone_number = $1
	UNION ALL SELECT gibdd_name from public.gibdd_full where phone_number = $1
	UNION ALL SELECT linkedin_name from public.linkedin_full where phone_number = $1
	UNION ALL SELECT mailru_full_name from public.mailru_full where phone_number = $1
	UNION ALL SELECT miltor_name from public.miltor_full where phone_number = $1
	UNION ALL SELECT okrug_pib from public.okrug_full where phone_number = $1
	UNION ALL SELECT pikabu_username from public.pikabu_full where phone_number = $1
	UNION ALL SELECT rfcont_name from public.rfcont_full where phone_number = $1
	UNION ALL SELECT sushi_name from public.sushi_full where phone_number = $1
	UNION ALL SELECT vk_first_name  from public.vk_full where phone_number = $1
	UNION ALL SELECT wildberries_name from public.wildberries_full where phone_number = $1`
	//UNION ALL SELECT fio from public.gipdd_nn_2020 where phone_number = $1

	queryAnswer := "Предпологаемое имя: "
	//Counting number of users

	rows, err := db.Query(qwery, phoneNumber)

	var names []string
	for rows.Next() {
		err = rows.Scan(&queryAnswer)
		if err != nil {
			return "", err
		}
		names = append(names, queryAnswer)
	}
	rows.Close()
	target := strings.Join(names, "\n\r ")

	if target == "" {
		target = "Нет Данных"
	}

	return target, nil
}

func GetTargetAdres(phoneNumber string, db *sqlx.DB) (string, error) {
	qwery := `SELECT delivery2_address_full from public.delivery2_full where phone_number = $1
	UNION ALL SELECT delivery_address from public.delivery_full where phone_number = $1
	UNION ALL SELECT fb_address1 from public.facebook_full where phone_number = $1
	UNION ALL SELECT gibdd2_address from public.gibdd2_full where phone_number = $1
	UNION ALL SELECT sushi_address_city || ' ' || sushi_address_street || ' ' ||sushi_address_home from public.sushi_full where phone_number = $1
	UNION ALL SELECT vk_first_name || ' ' || vk_last_name from public.vk_full where phone_number = $1
	UNION ALL SELECT wildberries_address from public.wildberries_full where phone_number = $1
	UNION ALL SELECT yandex_address_city || ' ' ||yandex_address_street|| ' ' ||yandex_address_house from public.yandex_full where phone_number = $1
	--UNION ALL SELECT address from public.gipdd_nn_2020 where phone_number = $1
	UNION ALL SELECT address_city || ' ' || actual_address from public.delivery_full where mobile_phone = $1`

	var target string

	var queryAnswer string
	//Counting number of users

	rows, err := db.Query(qwery, phoneNumber)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var adres []string
	for rows.Next() {
		err = rows.Scan(&queryAnswer)
		if err != nil {
			return "", err
		}
		adres = append(adres, queryAnswer)
	}
	rows.Close()
	target = strings.Join(adres, "\n\r")

	if target == "" {
		target = "Нет Данных"
	}

	return target, nil
}

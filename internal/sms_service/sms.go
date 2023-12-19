package sms_service

import (
	"os"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwdth     string
	ResponseTime string
	Provider     string
}

var provider = map[string]bool{"Topolo": true, "Rond": true, "Kildy": true}

func GetSMSData(filePath string) ([]SMSData, error) { //Читаем содержимое файла.
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	result := make([]SMSData, 0, len(lines))

	for _, line := range lines {
		fields := strings.Split(line, ";") //Разбиваем строку на поля.
		if len(fields) != 4 {              //Проверяем количество полей и игнорируем некорректные строки.
			continue
		}

		if validCountry(fields[0]) && provider[fields[3]] { // Проверяем корректность провайдера.
			result = append(result, SMSData{
				Country:      fields[0],
				Bandwdth:     fields[1],
				ResponseTime: fields[2],
				Provider:     fields[3],
			})
		}
	}
	return result, nil
}

func validCountry(countryCode string) bool {
	countryMap := map[string]bool{
		"RU": true,
		"US": true,
		"GB": true,
		"FR": true,
		"BL": true,
		"AT": true,
		"BG": true,
		"DK": true,
		"CA": true,
		"ES": true,
		"CH": true,
		"TR": true,
		"PE": true,
		"NZ": true,
		"MC": true,
	}
	return countryMap[countryCode]
}

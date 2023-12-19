package email_service

import (
	"os"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

func ReadEmailData(fileDataPath string) ([]EmailData, error) {
	var result []EmailData

	// Чтение содержимого файла
	file, err := os.ReadFile(fileDataPath)
	if err != nil {
		return nil, err
	}

	// Разбиение файла на строки
	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		// Разбиение строки на поля с использованием точки с запятой
		fields := strings.Split(line, ";")

		// Проверка количества полей
		if len(fields) != 3 {
			continue
		}

		// Преобразование времени доставки в тип int
		deliveryTime, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		// Проверка провайдера
		validProviders := map[string]bool{
			"Gmail":      true,
			"Yahoo":      true,
			"Hotmail":    true,
			"MSN":        true,
			"Orange":     true,
			"Comcast":    true,
			"AOL":        true,
			"Live":       true,
			"RediffMail": true,
			"GMX":        true,
			"Protonmail": true,
			"Yandex":     true,
			"Mail.ru":    true,
		}
		if !validProviders[fields[1]] {
			continue
		}

		// Создание структуры и добавление в результат
		emailFileData := EmailData{
			Country:      fields[0],
			Provider:     fields[1],
			DeliveryTime: deliveryTime,
		}

		result = append(result, emailFileData)
	}

	return result, nil
}

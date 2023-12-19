package billing_service

import (
	"math"
	"os"
)

// Определение структуры BillingData для хранения состояния систем
type BillingData struct {
	CreateCustomer bool // Создание клиента
	Purchase       bool // Оплата
	Payout         bool // Выплата
	Recurring      bool // Платежи по подписке
	FraudControl   bool // Контроль мошенничества
	CheckoutPage   bool // Страница оплаты
}

// Функция для чтения битовой маски из файла
func ReadMaskFromFile(filePath string) (BillingData, error) {
	// Чтение файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return BillingData{}, err // Возвращаем ошибку, если чтение не удалось
	}

	var value uint8
	// Проходим по строке с битовой маской
	for i, bit := range data {
		// Если текущий бит равен '1', добавляем его вес к общему значению
		if bit == '1' {
			value += uint8(math.Pow(2, float64(len(data)-i-1)))
		}
	}

	return getBillingData(value), nil // Возвращаем итоговое значение и nil (нет ошибки)
}

// Функция для получения данных о системе из битовой маски
func getBillingData(mask uint8) BillingData {
	// Заполняем структуру BillingData, проверяя каждый бит
	return BillingData{
		CreateCustomer: mask&(1<<0) != 0, // Проверяем 0-й бит
		Purchase:       mask&(1<<1) != 0, // Проверяем 1-й бит
		Payout:         mask&(1<<2) != 0, // Проверяем 2-й бит
		Recurring:      mask&(1<<3) != 0, // Проверяем 3-й бит
		FraudControl:   mask&(1<<4) != 0, // Проверяем 4-й бит
		CheckoutPage:   mask&(1<<5) != 0, // Проверяем 5-й бит
	}
}

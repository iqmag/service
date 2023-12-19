package sort_filt_service

import "sort"

type SMSData struct {
	Provider  string
	Country   string
	AlphaCode string
}

type MMSData struct {
	Provider  string
	Country   string
	AlphaCode string
}

type VoiceCall struct {
	ResultSetT
}

type EmailData struct {
	Provider     string
	Country      string
	DeliveryTime float64
}

type Billing struct {
}

type Support struct {
	Incidents []Incident
}

type Incident struct {
	Status        string
	ActiveTickets int
}

type ResultSetT struct {
	SMS     [][]SMSData
	MMS     [][]MMSData
	Voice   []VoiceCall
	Email   map[string][]EmailData
	Billing Billing
	Support Support
}

func GetResultData() ResultSetT {
	result := ResultSetT{}

	// Получение данных по SMS
	smsData := []SMSData{
		{"Topolo", "RU", "AlphaCode"},
		{"Rond", "US", "AlphaCode"},
		{"Kildy", "BL", "AlphaCode"},
	}

	// Сортировка данных по провайдеру
	sort.Slice(smsData, func(i, j int) bool {
		return smsData[i].Provider < smsData[j].Provider
	})

	// Сортировка данных по стране
	sort.Slice(smsData, func(i, j int) bool {
		return smsData[i].Country < smsData[j].Country
	})

	result.SMS = append(result.SMS, smsData)

	// Получение данных по MMS
	mmsData := []MMSData{
		{"Topolo", "RU", "AlphaCode"},
		{"Rond", "US", "AlphaCode"},
		{"Kildy", "BL", "AlphaCode"},
		// остальные данные по MMS
	}

	// Сортировка данных по провайдеру
	sort.Slice(mmsData, func(i, j int) bool {
		return mmsData[i].Provider < mmsData[j].Provider
	})

	// Сортировка данных по стране
	sort.Slice(mmsData, func(i, j int) bool {
		return mmsData[i].Country < mmsData[j].Country
	})

	result.MMS = append(result.MMS, mmsData)

	// Получение данных по Voice Call
	voiceData := []VoiceCall{
		// данные по голосовым вызовам
	}

	result.Voice = voiceData

	// Получение данных по Email
	emailData := map[string][]EmailData{
		"Country1": {
			{"Topolo", "RU", 2.7},
			{"Rond", "US", 9.1},
			{"Kildy", "BL", 5.7},
			// остальные данные по Email для Country1
		},
		"Country2": {
			{"Topolo", "RU", 7.3},
			{"Rond", "US", 4.4},
			{"Kildy", "BL", 3.2},
			// остальные данные по Email для Country2
		},
		// остальные данные по странам
	}

	// Сортировка провайдеров по времени доставки письма
	for _, emails := range emailData {
		sort.Slice(emails, func(i, j int) bool {
			return emails[i].DeliveryTime < emails[j].DeliveryTime
		})
	}

	result.Email = emailData

	// Получение данных о системе Billing
	billingData := Billing{
		// данные о системе Billing
	}

	result.Billing = billingData

	// Получение данных о системе Support
	supportData := Support{
		Incidents: []Incident{
			{Status: "active"},
			{Status: "resolved"},
			// остальные данные об инцидентах
		},
	}

	// Сортировка инцидентов по статусу
	sort.Slice(supportData.Incidents, func(i, j int) bool {
		return supportData.Incidents[i].Status == "active"
	})

	result.Support = supportData

	return result
}

package struct_data

import (
	"encoding/json"
	"net/http"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing_service"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

type SMSData struct {
	Country      string `json:"country"`
	Bandwdth     string `json:"bandwdth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// Функция для выполнения GET запроса и разбора ответа

func GetData() (ResultT, error) {

	// Выполнение GET запроса
	resp, err := http.Get("http://127.0.0.1:8585")
	if err != nil {
		return ResultT{}, err
	}
	defer resp.Body.Close()

	// Декодирование ответа
	var result ResultT
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ResultT{}, err
	}

	return result, nil
}

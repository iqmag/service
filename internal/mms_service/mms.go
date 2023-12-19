package mms_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func GetMMSData() ([]MMSData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/mms")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []MMSData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	checkData := make([]MMSData, 0)
	for _, d := range data {
		if theValidCountry(d.Country) && theValidProvider(d.Provider) {
			checkData = append(checkData, d)
		}
	}
	return checkData, nil
}

func theValidCountry(country string) bool {
	validCountry := []string{
		"RU",
		"US",
		"GB",
		"FR",
		"BL",
		"AT",
		"BG",
		"DK",
		"CA",
		"ES",
		"CH",
		"TR",
		"PE",
		"NZ",
		"MC",
	}
	for _, c := range validCountry {
		if c == country {
			return true
		}
	}
	return false
}

func theValidProvider(provider string) bool {
	validProvider := []string{
		"Topolo",
		"Rond",
		"Kildy"}
	for _, p := range validProvider {
		if p == provider {
			return true
		}
	}
	return false
}

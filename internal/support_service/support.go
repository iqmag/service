package support_service

import (
	"encoding/json"
	"io"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func GetSupportData() ([]SupportData, error) {
	resp, err := http.Get("http://127.0.0.1:8383/support")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []SupportData{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []SupportData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []SupportData{}, nil
	}
	return data, nil
}

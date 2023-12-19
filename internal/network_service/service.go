package network_service

import (
	"encoding/json"
	"net/http"
)

type ServiceData struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	var response []ServiceData

	response = append(response, ServiceData{"Value1", 1})
	response = append(response, ServiceData{"Value2", 2})
	response = append(response, ServiceData{"Value3", 3})

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

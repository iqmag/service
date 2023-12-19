package voice_call_service

import (
	"os"
	"strconv"
	"strings"
)

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

func FiltrationVoiceCallData(filePath string) ([]VoiceCallData, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n") //Разделение на строки
	var theVoiceCallData []VoiceCallData

	for _, line := range lines {
		fields := strings.Split(line, ";")
		if len(fields) != 8 {
			continue
		}

		country := fields[0]
		bandwidth := fields[1]
		responseTime := fields[2]
		provider := fields[3]
		connectionStability, err := strconv.ParseFloat(fields[4], 32)
		if err != nil {
			continue
		}
		ttfb, err := strconv.Atoi(fields[5])
		if err != nil {
			continue
		}
		voicePurity, err := strconv.Atoi(fields[6])
		if err != nil {
			continue
		}
		medianOfCallsTime, err := strconv.Atoi(fields[7])
		if err != nil {
			continue
		}

		if !validCountry(country) || !validProvider(provider) {
			continue
		}

		theVoiceCallData = append(theVoiceCallData, VoiceCallData{
			Country:             country,
			Bandwidth:           bandwidth,
			ResponseTime:        responseTime,
			Provider:            provider,
			ConnectionStability: float32(connectionStability),
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		})
	}
	return theVoiceCallData, nil
}

func validCountry(country string) bool {
	voiceProviderMap := map[string]bool{
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
	return voiceProviderMap[country]
}

func validProvider(provider string) bool {
	providerMap := map[string]bool{
		"TransparentCalls": true,
		"E-Voice":          true,
		"JustPhone":        true,
	}

	return providerMap[provider]
}

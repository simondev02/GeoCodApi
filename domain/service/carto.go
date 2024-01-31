package carto

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getGeocodingReverse(lat float64, lon float64) (string, error) {
	encodedToken := os.Getenv("TOKEN_CARTO")
	encodedApiBaseURL := os.Getenv("ApiBaseUrlCarto")

	token, err := base64.StdEncoding.DecodeString(encodedToken)
	if err != nil {
		return "", err
	}

	apiBaseURL, err := base64.StdEncoding.DecodeString(encodedApiBaseURL)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/v3/lds/geocoding/reverse?lat=%f&lon=%f&language=STRING", string(apiBaseURL), lat, lon)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+string(token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

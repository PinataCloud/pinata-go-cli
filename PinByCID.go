package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func PinByCID(cid string, name string) (PinByCIDResponse, error) {
	jwt, err := findToken()
	if err != nil {
		return PinByCIDResponse{}, err
	}
	host := GetHost()
	url := fmt.Sprintf("https://%s/pinning/pinByHash", host)

	// Create the request body
	requestBody := map[string]interface{}{
		"hashToPin": cid,
		"pinataMetadata": map[string]string{
			"name": name,
		},
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return PinByCIDResponse{}, errors.New("failed to create JSON body")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return PinByCIDResponse{}, errors.New("failed to create the request")
	}
	req.Header.Set("Authorization", "Bearer "+string(jwt))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PinByCIDResponse{}, errors.New("failed to send the request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return PinByCIDResponse{}, fmt.Errorf("server returned an error %s", resp.Status)
	}

	var response PinByCIDResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return PinByCIDResponse{}, err
	}

	fmt.Println("Pin by CID Request Started")
	fmt.Println("Request ID:", response.Id)
	fmt.Println("CID:", response.CID)
	fmt.Println("Status:", response.Status)
	fmt.Println("Name:", response.Name)

	return response, nil
}


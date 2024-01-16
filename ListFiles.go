package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ListFiles(amount string, cid string, name string, status string, offset string) (ListResponse, error) {
	jwt, err := findToken()
	if err != nil {
		return ListResponse{}, err
	}
	host := GetHost()
	url := fmt.Sprintf("https://%s/data/pinList?includesCount=false&pageLimit=%s&status=%s", host, amount, status)

	if cid != "null" {
		url += "&hashContains=" + cid
	}
	if name != "null" {
		url += "&metadata[name]=" + name
	}
	if offset != "null" {
		url += "&pageOffset=" + offset
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ListResponse{}, errors.Join(err, errors.New("failed to create the request"))
	}
	req.Header.Set("Authorization", "Bearer "+string(jwt))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ListResponse{}, errors.Join(err, errors.New("failed to send the request"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return ListResponse{}, fmt.Errorf("server Returned an error %d", resp.StatusCode)
	}

	var response ListResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return ListResponse{}, err
	}
	formattedJSON, err := json.MarshalIndent(response.Rows, "", "    ")
	if err != nil {
		return ListResponse{}, errors.New("failed to format JSON")
	}

	fmt.Println(string(formattedJSON))

	return response, nil

}

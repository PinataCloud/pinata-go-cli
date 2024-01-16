package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func Requests(cid string, status string, offset string) (RequestsResponse, error) {
	jwt, err := findToken()
	if err != nil {
		return RequestsResponse{}, err
	}
	host := GetHost()
	url := fmt.Sprintf("https://%s/pinning/pinJobs?sort=DESC", host)

	if cid != "null" {
		url += "&ipfs_pin_hash=" + cid
	}
	if offset != "null" {
		url += "&offset=" + offset
	}
	if status != "null" {
		url += "&status=" + status
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RequestsResponse{}, errors.Join(err, errors.New("failed to create the request"))
	}
	req.Header.Set("Authorization", "Bearer "+string(jwt))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return RequestsResponse{}, errors.Join(err, errors.New("failed to send the request"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RequestsResponse{}, fmt.Errorf("server Returned an error %d", resp.StatusCode)
	}

	var response RequestsResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return RequestsResponse{}, err
	}

	for i := 0; i < len(response.Rows); i++ {
		fmt.Println("Request ID:", response.Rows[i].Id)
		fmt.Println("CID:", response.Rows[i].CID)
		fmt.Println("Date Started:", response.Rows[i].StartDate)
		fmt.Println("Name:", response.Rows[i].Name)
		fmt.Println("Status:", response.Rows[i].Status)
		fmt.Println()
	}

	return response, nil

}

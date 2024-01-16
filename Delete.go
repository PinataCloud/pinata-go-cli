package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Delete(cid string) error {
	jwt, err := findToken()
	if err != nil {
		return err
	}
	host := GetHost()
	url := fmt.Sprintf("https://%s/pinning/unpin/%s", host, cid)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return errors.Join(err, errors.New("failed to create the request"))
	}
	req.Header.Set("Authorization", "Bearer "+string(jwt))
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Join(err, errors.New("failed to send the request"))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("server Returned an error %d, check CID", resp.StatusCode)
	}

	fmt.Println("File Deleted")

	return nil

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	statsURL = "http://localhost:8080/v1/strum/stats"
)

func postToAcksin(conf *config) {
	var err error

	if conf.apiKey == "" {
		fmt.Fprintln(os.Stderr, "Provide the -api-key flag or set the ACKSIN_API_KEY.")
		fmt.Fprintln(os.Stderr, "The API Key can be gathered at https://www.acksin.com/console/credentials")
		os.Exit(-1)
	}

	jsonStr, err := json.Marshal(conf.stats)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse json")
		return
	}

	req, err := http.NewRequest("POST", statsURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Acksin-API-Key", conf.apiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "1 An error occured", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var respForm struct {
		ID string
	}

	if err = json.Unmarshal(body, &respForm); err != nil {
		fmt.Fprintln(os.Stderr, "An error occured", err)
		return
	}

	fmt.Printf("https://www.acksin.com/console/strum/#/%s\n", respForm.ID)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

	resp, err := http.Post(fmt.Sprintf("https://%s:@api.acksin.com/v1/strum/stats", conf.apiKey), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occured", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var respForm struct {
		ID string
	}

	err = json.Unmarshal(body, &respForm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occured", err)
		return
	}

	fmt.Printf("https://www.acksin.com/console/strum/%s\n", respForm.ID)
}

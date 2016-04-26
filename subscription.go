package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Subscription int

const (
	OpenSubscription Subscription = iota
	StartupSubscription

	// No yet released.
	ProSubscription
	PremiumSubscription
	EnterpriseSubscription
)

const (
	subscriptionAPI = "https://bridge-api.acksin.com/subscription/autotune"
)

var (
	currentSubscription = OpenSubscription
)

func loadSubscription(apiKey string) {
	if apiKey == "" {
		currentSubscription = OpenSubscription
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", subscriptionAPI, nil)
	req.SetBasicAuth(apiKey, "")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var sub struct {
		Product string
		Plan    string
	}

	err = json.Unmarshal(body, &sub)
	if err != nil {
		return
	}

	switch sub.Plan {
	case "autotune-startup":
		currentSubscription = StartupSubscription
	case "autotune-pro":
		currentSubscription = ProSubscription
	case "autotune-premium":
		currentSubscription = PremiumSubscription
	}
}

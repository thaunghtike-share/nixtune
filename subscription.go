package main

import (
	"log"

	"github.com/abhiyerra/gumroad"
)

type Subscription int

const (
	OpenSubscription Subscription = iota
	ProSubscription
	PremiumSubscription
	EnterpriseSubscription
)

var (
	currentSubscription Subscription
)

func setSubscription(apiKey string) {
	currentSubscription = OpenSubscription

	if apiKey == "" {
		return
	}

	if err := gumroad.VerifyLicense("autotune-pro", apiKey, true); err == nil {
		currentSubscription = ProSubscription
	}

	if err := gumroad.VerifyLicense("autotune-premium", apiKey, true); err == nil {
		currentSubscription = PremiumSubscription
	}
}

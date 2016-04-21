package main

import (
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
	currentSubscription = OpenSubscription
)

func init() {
}

func loadSubscription(apiKey string) {
	if apiKey == "" {
		currentSubscription = OpenSubscription
		return
	}

	if err := gumroad.VerifyLicense("autotune-pro", apiKey, true); err == nil {
		currentSubscription = ProSubscription
	}

	if err := gumroad.VerifyLicense("autotune-premium", apiKey, true); err == nil {
		currentSubscription = PremiumSubscription
	}
}

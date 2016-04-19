package main

import (
	"log"

	"github.com/acksin/utils-go/credentials"
)

type Subscription int

const (
	OpenSubscription Subscription = iota
	ProSubscription
	PremiumSubscription
	EnterpriseSubscription
)

func setSubscription(apiKey string) (s Subscription) {
	s = OpenSubscription

	if apiKey == "" {
		return
	}

	session, err := fugue_credentials.GetSessionID(apiKey)
	if err != nil {
		log.Println("Invalid API Key")
		return
	}

	subscription, err := fugue_credentials.GetSubscription("Autotune", session)
	if err != nil {
		return
	}

	if subscription == nil {
		return
	}

	switch subscription.Name {
	case "Open":
		s = OpenSubscription
	case "Pro":
		s = ProSubscription
	case "Premium":
		s = PremiumSubscription
	case "Enterprise":
		s = EnterpriseSubscription
	}

	return s
}

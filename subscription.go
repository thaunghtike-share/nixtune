package main

import (
	"log"

	"github.com/acksin/fugue/client"
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

	session, err := fugueapi.GetSessionID(apiKey)
	if err != nil {
		log.Println("Invalid API Key")
		return
	}

	subscription, err := fugueapi.GetSubscription("Autotune", session)
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

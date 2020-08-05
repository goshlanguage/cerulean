package subscriptions

import (
	"github.com/goshlanguage/cerulean/services"
	"github.com/labstack/echo/v4"
)

// SubscriptionService satisfies the Service interface, and is used to start and maintain the Subscription Service
type SubscriptionService struct {
	// Subscriptions is our statebag
	Subscriptions []Subscription
}

// NewSubscriptionService is a factory for the SubscriptionService, which satisfies the services.Service interface
func NewSubscriptionService() services.Service {
	initSub := NewSubscription()
	return &SubscriptionService{
		Subscriptions: []Subscription{initSub},
	}
}

// GetHandlers returns the echo handlers that the service needs in order to operate
func (svc *SubscriptionService) GetHandlers() map[string]echo.HandlerFunc {
	svcMap := make(map[string]echo.HandlerFunc)
	svcMap["/subscriptions"] = svc.GetSubscriptionsHandler()
	return svcMap
}

// GetBaseSubscriptionID is a SubscriptionService specific helper that returns the initial subscriptionID
func (svc *SubscriptionService) GetBaseSubscriptionID() string {
	return svc.Subscriptions[0].SubscriptionID
}

package subscriptions

import (
	"net/http"

	"github.com/goshlanguage/cerulean/services"
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

// GetHandlers returns the HTTP GET Echo handlers that the service needs in order to operate
func (svc *SubscriptionService) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/subscriptions"] = services.Handler{http.MethodGet, svc.GetSubscriptionsHandler()}
	return svcMap
}

// GetBaseSubscriptionID is a SubscriptionService specific helper that returns the initial subscriptionID
func (svc *SubscriptionService) GetBaseSubscriptionID() string {
	return svc.Subscriptions[0].SubscriptionID
}

package subscriptions

import (
	"encoding/json"

	"net/http"

	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/goshlanguage/cerulean/services"
)

const serviceKey = "subscriptions"

// SubscriptionService satisfies the Service interface, and is used to start and maintain the Subscription Service
type SubscriptionService struct {
	Store *lightdb.Store
}

// NewSubscriptionService is a factory for the SubscriptionService, which satisfies the services.Service interface and provides a default sub
// TODO Error handling
func NewSubscriptionService(s *lightdb.Store) *SubscriptionService {
	service := &SubscriptionService{
		Store: s,
	}
	initSub := NewSubscription()
	service.AddSubscription(initSub)

	return service
}

// GetHandlers returns a map of all HTTP Echo handlers that the service needs in order to operate
func (svc *SubscriptionService) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/subscriptions"] = services.Handler{http.MethodGet, svc.GetSubscriptionsHandler()}
	return svcMap
}

// GetBaseSubscriptionID is a SubscriptionService specific helper that returns the initial subscriptionID
func (svc *SubscriptionService) GetBaseSubscriptionID() string {
	subsString, err := svc.Store.Get(serviceKey)
	if err != nil {
		panic(err)
	}

	var subs []Subscription
	err = json.Unmarshal([]byte(subsString), &subs)
	if err != nil {
		panic(err)
	}

	return subs[0].SubscriptionID
}

// GetSubscriptions returns the Stores state
func (svc *SubscriptionService) GetSubscriptions() ([]Subscription, error) {
	var subs []Subscription
	subsString, err := svc.Store.Get(serviceKey)
	if err != nil {
		return subs, err
	}

	err = json.Unmarshal([]byte(subsString), &subs)
	if err != nil {
		return subs, err
	}

	return subs, nil
}

// AddSubscription takes a subscription and adds it to the store
func (svc *SubscriptionService) AddSubscription(s Subscription) error {
	subsString, err := svc.Store.Get(serviceKey)
	if err != nil {
		return err
	}

	// if there are existing subs, be sure to deserialize the response and append
	var subs []Subscription
	var subsBytes []byte
	if subsString != "" {
		err := json.Unmarshal([]byte(subsString), &subs)
		if err != nil {
			return err
		}

	}
	subs = append(subs, s)

	subsBytes, err = json.Marshal(subs)
	if err != nil {
		return err
	}
	subsString = string(subsBytes)

	svc.Store.Put(serviceKey, subsString)
	return nil
}

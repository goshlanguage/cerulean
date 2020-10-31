package subscriptions

import (
	"encoding/json"

	"net/http"

	"github.com/goshlanguage/cerulean/internal/services"
	"github.com/goshlanguage/cerulean/pkg/lightdb"
)

const serviceKey = "subscriptions"

// Service satisfies the Service interface, and is used to start and maintain the Subscription Service
type Service struct {
	Store *lightdb.Store
}

// NewService is a factory for the SubscriptionService, which satisfies the services.Service interface and provides a default sub
// TODO Error handling
func NewService(s *lightdb.Store) *Service {
	service := &Service{
		Store: s,
	}
	initSub := NewSubscription()
	service.AddSubscription(initSub)

	return service
}

// GetHandlers returns a map of all HTTP Echo handlers that the service needs in order to operate
func (svc *Service) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/subscriptions"] = services.Handler{
		Verb: http.MethodGet,
		Func: svc.GetHandler(),
	}
	return svcMap
}

// GetBaseSubscriptionID is a SubscriptionService specific helper that returns the initial subscriptionID
func (svc *Service) GetBaseSubscriptionID() string {
	subsString := svc.Store.Get(serviceKey)

	var subs []Subscription
	err := json.Unmarshal([]byte(subsString), &subs)
	if err != nil {
		panic(err)
	}

	return subs[0].SubscriptionID
}

// GetSubscriptions returns the Stores state
func (svc *Service) GetSubscriptions() ([]Subscription, error) {
	var subs []Subscription
	subsString := svc.Store.Get(serviceKey)

	err := json.Unmarshal([]byte(subsString), &subs)
	if err != nil {
		return subs, err
	}

	return subs, nil
}

// AddSubscription takes a subscription and adds it to the store
func (svc *Service) AddSubscription(s Subscription) error {
	subsString := svc.Store.Get(serviceKey)

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

	subsBytes, err := json.Marshal(subs)
	if err != nil {
		return err
	}
	subsString = string(subsBytes)

	svc.Store.Put(serviceKey, subsString)
	return nil
}

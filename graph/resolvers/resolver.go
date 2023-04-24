package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import m "gqlgen-subscriptions/graph/model"

type Resolver struct {
	EventChannel chan *m.Event
}

package resolvers

import (
	"gqlgen-subscriptions/utils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Pubsub *utils.Pubsub
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/oluwakeye-john/wallet-alert/graph/generated"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/handlers"
)

func (r *mutationResolver) CreateSubscription(ctx context.Context, input model.SubscriptionInput) (*model.SubscriptionStatus, error) {
	return handlers.CreateSubscription(ctx, input)
}

func (r *mutationResolver) CancelSubscription(ctx context.Context, address string) (*model.SubscriptionStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SupportedCurrencies(ctx context.Context) ([]*model.Currency, error) {
	return handlers.GetSupportedCurrencies(ctx)
}

func (r *queryResolver) GetSubscriptionStatus(ctx context.Context, email string) (*model.SubscriptionStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

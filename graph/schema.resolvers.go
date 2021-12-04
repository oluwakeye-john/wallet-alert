package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/oluwakeye-john/wallet-alert/graph/generated"
	"github.com/oluwakeye-john/wallet-alert/graph/model"
	"github.com/oluwakeye-john/wallet-alert/handlers"
)

func (r *mutationResolver) CreateTestAddress(ctx context.Context) (*model.Key, error) {
	return handlers.CreateTestAddress()
}

func (r *mutationResolver) FundTestAddress(ctx context.Context, input model.FundTestAddressInput) (*model.Transaction, error) {
	return handlers.FundTestAddress(input)
}

func (r *mutationResolver) DeleteAddressHook(ctx context.Context, input model.DeleteHookInput) (bool, error) {
	return handlers.DeleteAddressHook(input)
}

func (r *mutationResolver) CreateSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.SubscriptionStatus, error) {
	return handlers.CreateSubscription(ctx, input)
}

func (r *mutationResolver) CancelSubscription(ctx context.Context, input model.CancelSubscriptionInput) (*model.SubscriptionStatus, error) {
	return handlers.CancelSubscription(ctx, input)
}

func (r *queryResolver) SupportedCurrencies(ctx context.Context) ([]*model.Currency, error) {
	return handlers.GetSupportedCurrencies(ctx)
}

func (r *queryResolver) GetSubscriptionStatus(ctx context.Context, input model.GetStatusInput) (*model.SubscriptionStatus, error) {
	return handlers.GetSubscriptionStatus(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

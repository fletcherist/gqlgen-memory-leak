package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gqlgen-memory-leak/graph/generated"
	"gqlgen-memory-leak/graph/model"
	"time"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Test(ctx context.Context, name string) (<-chan *model.Todo, error) {
	// log.Println("connected to subscription Test")
	newChan := make(chan *model.Todo, 1)

	go func() {
		count := 0
		for {
			time.Sleep(time.Second * 1)
			newChan <- &model.Todo{
				ID:   fmt.Sprint(count),
				Text: "test",
			}
			count += 1
		}
	}()

	return newChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

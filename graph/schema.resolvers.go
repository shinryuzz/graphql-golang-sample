package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/shinryuzz/graphql-golang-sample/graph/model"
	"github.com/shinryuzz/graphql-golang-sample/internal"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := []*model.Todo{
		{
			ID:   "TODO-1",
			Text: "My Todo 1",
			User: &model.User{
				ID:   "User-1",
				Name: "user1",
			},
			Done: false,
		},
		{
			ID:   "TODO-2",
			Text: "My Todo 2",
			User: &model.User{
				ID:   "User-2",
				Name: "user2",
			},
			Done: true,
		},
	}
	return todos, nil
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.User.ID, Name: "user " + obj.User.ID}, nil
}

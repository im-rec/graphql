package graphql

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"github.com/rianekacahya/graphql/internal/v1/module/hello"
	"github.com/rianekacahya/graphql/internal/v1/entity"
)

type Resolver struct{
	usecase hello.Usecase
}

func (r *queryResolver) GetHello(ctx context.Context) (*entity.Hello, error) {
	return r.usecase.GetHello()
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

package hello

import (
	"github.com/rianekacahya/graphql/internal/v1/invoker/event"
	"github.com/rianekacahya/graphql/internal/v1/invoker/grpc"
	"github.com/rianekacahya/graphql/internal/v1/invoker/rest"
	"github.com/rianekacahya/graphql/internal/v1/module/hello/repository"
)

type hello struct {
	repository Repository
	invoker    invoker
}

type invoker struct {
	rest  rest.Handler
	grpc  grpc.Handler
	event event.Handler
}

func Initialize() *hello {
	return &hello{
		repository: repository.NewHello(),
		invoker: invoker{
			rest:  rest.Initialize(),
			grpc:  grpc.Initialize(),
			event: event.Initialize(),
		},
	}
}

type Usecase interface{}

type Repository interface{}

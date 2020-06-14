package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo"
	"github.com/rianekacahya/graphql/internal/v1/module/hello"
	"github.com/rianekacahya/graphql/pkg/echoserver"
)

func NewGraphQL(usecase hello.Usecase) {
	transport := handler.NewDefaultServer(NewExecutableSchema(Config{
		Resolvers: &Resolver{
			usecase: usecase,
		},
	}))

	attribute := echoserver.GetServer().Group("/v1/hello")
	attribute.POST("", echo.WrapHandler(transport))
}

package rest

import (
	"github.com/labstack/echo"
	"github.com/rianekacahya/graphql/internal/v1/module/hello"
	"github.com/rianekacahya/graphql/pkg/echoserver"
	"github.com/rianekacahya/graphql/pkg/echoserver/response"
)

type rest struct {
	usecase hello.Usecase
}

func NewRest(usecase hello.Usecase) {
	transport := rest{usecase}

	attribute := echoserver.GetServer().Group("/v1/hello")

	attribute.GET("", transport.list)
}

func (DI *rest) list(c echo.Context) error {

	return response.Render(c, "Hello world.")
}
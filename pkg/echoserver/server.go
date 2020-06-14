package echoserver

import (
	"context"
	"fmt"
	"github.com/rianekacahya/graphql/internal/v1/config"
	"github.com/rianekacahya/graphql/pkg/echoserver/adapter"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	server *echo.Echo
	mutex  sync.Once
)

func GetServer() *echo.Echo {
	mutex.Do(func() {
		server = echo.New()
	})

	return server
}

func InitServer(debug bool) {

	// Hide banner
	GetServer().HideBanner = true

	// Set debug status parameter
	GetServer().Debug = config.GetConfig().AppDebug

	// init default adapter
	GetServer().Use(
		middleware.RequestID(),
		middleware.RecoverWithConfig(middleware.RecoverConfig{
			DisableStackAll:   true,
			DisablePrintStack: true,
		}),
		adapter.CORS(),
		adapter.Headers(),
		middleware.BodyDump(adapter.Logger),
	)

	// custom error handler
	//GetServer().HTTPErrorHandler = Handler

	// Custom binder
	GetServer().Binder = &CustomBinder{bind: &echo.DefaultBinder{}}
}

func StartServer(port, writetimeout, readtimeout, idletimeout int) {
	if err := GetServer().StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: time.Duration(writetimeout) * time.Second,
		ReadTimeout:  time.Duration(readtimeout) * time.Second,
		IdleTimeout:  time.Duration(idletimeout) * time.Second,
	}); err != nil {
		log.Fatal(err.Error())
	}
}

func Shutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := GetServer().Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

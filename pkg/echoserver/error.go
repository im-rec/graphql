package echoserver

import (
	"github.com/labstack/echo"
)

func Handler(err error, c echo.Context) {
	//var (
	//	code        = http.StatusInternalServerError
	//	msg, logmsg interface{}
	//)
	//
	//logmsg = err.Error()
	//if GetServer().Debug {
	//	switch e := err.(type) {
	//	case *echo.HTTPError:
	//		code = e.Code
	//		msg = e.Message
	//	default:
	//		msg = e.Error()
	//	}
	//} else {
	//
	//}

	//if !c.Response().Committed {
	//	switch e := err.(type) {
	//	case *echo.HTTPError:
	//		msg := e.Message
	//		if e.Internal != nil {
	//			msg = fmt.Sprintf("%v, %v", err, e.Internal)
	//		}
	//
	//		if c.Request().Method == "HEAD" {
	//			err = c.NoContent(e.Code)
	//		} else {
	//			err = c.JSON(e.Code, response.Response{
	//				Success: false,
	//				Errors:&response.Customerror{
	//					Status: 0,
	//					Message: fmt.Sprint(msg),
	//				},
	//			})
	//		}
	//	//default:
	//	//	msg = http.StatusText(code)
	//	//	logmsg = e.Error()
	//	}

	//err = response.Error(c, err)
	//if err != nil {
	//	log.Fatalf("got an error while serve data, error: %s", err)
	//}
}

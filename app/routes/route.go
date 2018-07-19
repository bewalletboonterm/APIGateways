package route

import (
	"net/http"

	"APIGateways/app/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//Init route
func Init() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "Welcome mvc echo with mysql app using Golang")
	})

	gateways := e.Group("/gateways")
	{
		member := gateways.Group("/member")
		{
			member.POST("/sendOTPRegister", MemberController.SendOTPRegister)
			member.POST("/validateOTPByPhone", MemberController.ValidateOTPByPhone)
			member.POST("/validateBeforeRegister", MemberController.ValidateBeforeRegister)
			member.POST("/register", MemberController.Register)
		}

	}

	return e
}

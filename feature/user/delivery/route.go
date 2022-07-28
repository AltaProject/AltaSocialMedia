package delivery

import (
	"github.com/AltaProject/AltaSocialMedia/config"
	"github.com/AltaProject/AltaSocialMedia/domain"
	"github.com/AltaProject/AltaSocialMedia/feature/user/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, bu domain.UserHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/login", bu.Login())
	e.POST("/register", bu.Register())
	e.GET("/user/:id", bu.GetSpecificUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/user/:id", bu.UpdateUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/user/id", bu.DeleteUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}

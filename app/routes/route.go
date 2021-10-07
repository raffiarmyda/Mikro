package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mikro/app/middlewares"
	"mikro/deliveries/users"
)

type ControllerList struct {
	LoggerMiddleware   middlewares.MongoConfig
	JWTMiddleware      middleware.JWTConfig
	JWTMiddlewareAdmin middleware.JWTConfig
	UserController     users.Controller
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	cl.LoggerMiddleware.Start(e)
	//AUTH
	v1.POST("auth/login", cl.UserController.LoginUserController)

	//USERS
	//middleware.JWTWithConfig(cl.JWTMiddleware)
	v1.GET("users", cl.UserController.GetUsersController)
	v1.GET("users/:userId", cl.UserController.GetDetailUserController)
	v1.POST("users", cl.UserController.CreateUserController)
	v1.DELETE("users/:userId", cl.UserController.DeleteUserController)
	v1.PUT("users/:userId", cl.UserController.UpdateUserController)
}

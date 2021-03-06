package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mikro/app/middlewares"
	"mikro/deliveries/products"
	"mikro/deliveries/transactions"
	"mikro/deliveries/users"
)

type ControllerList struct {
	LoggerMiddleware      middlewares.MongoConfig
	JWTMiddleware         middleware.JWTConfig
	JWTMiddlewareAdmin    middleware.JWTConfig
	UserController        users.Controller
	ProductController     products.Controller
	TransactionController transactions.Controller
}

func (cl *ControllerList) RouteUsers(e *echo.Echo) {
	v1 := e.Group("api/v1/")
	cl.LoggerMiddleware.Start(e)
	//AUTH
	v1.POST("auth/login", cl.UserController.LoginUserController)

	//USERS
	//middleware.JWTWithConfig(cl.JWTMiddleware)
	v1.GET("users", cl.UserController.GetUsersController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.GET("users/:userId", cl.UserController.GetDetailUserController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.POST("users", cl.UserController.CreateUserController)
	v1.DELETE("users/:userId", cl.UserController.DeleteUserController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.PUT("users/:userId", cl.UserController.UpdateUserController, middleware.JWTWithConfig(cl.JWTMiddleware))

	v1.GET("products", cl.ProductController.GetProductsController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.GET("products/:productId", cl.ProductController.GetDetailProductController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.POST("products", cl.ProductController.CreateProductController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.DELETE("products/:productId", cl.ProductController.DeleteProductController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.PUT("products/:productId", cl.ProductController.UpdateProductController, middleware.JWTWithConfig(cl.JWTMiddleware))

	v1.GET("transactions", cl.TransactionController.GetTransactionsController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.GET("transactions/:transactionId", cl.TransactionController.GetDetailTransactionController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.POST("transactions", cl.TransactionController.CreateTransactionController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.DELETE("transactions/:transactionId", cl.TransactionController.DeleteTransactionController, middleware.JWTWithConfig(cl.JWTMiddleware))
	v1.PUT("transactions/:transactionId", cl.TransactionController.UpdateTransactionController, middleware.JWTWithConfig(cl.JWTMiddleware))
}

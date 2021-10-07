package users

import (
	"github.com/labstack/echo/v4"
	"mikro/business/users"
	"mikro/deliveries"
	"mikro/deliveries/users/requests"
	"mikro/deliveries/users/responses"
	"net/http"
	"strconv"
)

type Controller struct {
	userUsecase users.Usecase
}

func NewUserController(uc users.Usecase) *Controller {
	return &Controller{
		userUsecase: uc,
	}
}

func (ctrl *Controller) GetUsersController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.userUsecase.GetAll(ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersListDomain(data))
}

func (ctrl *Controller) GetDetailUserController(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("userId")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	user := users.Domain{
		ID: convInt,
	}
	data, err := ctrl.userUsecase.GetById(ctxNative, user)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (ctrl *Controller) CreateUserController(c echo.Context) error {
	request := requests.UserRegister{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.userUsecase.Create(ctx, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (ctrl *Controller) LoginUserController(c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctxNative := c.Request().Context()

	request := requests.UserLogin{}
	err = c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	login, err = ctrl.userUsecase.Login(ctxNative, request.ToDomain())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomainToLogin(login, token))
}

func (cl *Controller) UpdateUserController(c echo.Context) error {
	id := c.Param("userId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.UsersRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	req.ID = convId
	data, err := cl.userUsecase.Update(ctx, req.ToDomainUser())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromUsersDomain(data))
}

func (cl *Controller) DeleteUserController(c echo.Context) error {
	id := c.Param("userId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	user := users.Domain{
		ID: convId,
	}
	err = cl.userUsecase.Delete(ctx, user)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, convId)
}

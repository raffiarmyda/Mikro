package transactions

import (
	"github.com/labstack/echo/v4"
	"mikro/business/transactions"
	"mikro/deliveries"
	"mikro/deliveries/transactions/requests"
	"mikro/deliveries/transactions/responses"
	"net/http"
	"strconv"
)

type Controller struct {
	productUsecase transactions.Usecase
}

func NewTransactionController(uc transactions.Usecase) *Controller {
	return &Controller{
		productUsecase: uc,
	}
}

func (ctrl *Controller) GetTransactionsController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.productUsecase.GetAll(ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductListDomain(data))
}

func (ctrl *Controller) GetDetailTransactionController(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("transactionId")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	product := transactions.Domain{
		ID: convInt,
	}
	data, err := ctrl.productUsecase.GetById(ctxNative, product)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (ctrl *Controller) CreateTransactionController(c echo.Context) error {
	request := requests.TransactionRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.productUsecase.Create(ctx, request.ToDomainUser())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (cl *Controller) UpdateTransactionController(c echo.Context) error {
	id := c.Param("transactionId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.TransactionRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	domain := transactions.Domain{
		ID:        convId,
		ProductId: req.ProductId,
		BuyerId:   req.BuyerId,
	}
	data, err := cl.productUsecase.Update(ctx, domain)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (cl *Controller) DeleteTransactionController(c echo.Context) error {
	id := c.Param("transactionId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	product := transactions.Domain{
		ID: convId,
	}
	err = cl.productUsecase.Delete(ctx, product)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, convId)
}

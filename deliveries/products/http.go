package products

import (
	"github.com/labstack/echo/v4"
	"mikro/business/products"
	"mikro/deliveries"
	"mikro/deliveries/products/requests"
	"mikro/deliveries/products/responses"
	"net/http"
	"strconv"
)

type Controller struct {
	productUsecase products.Usecase
}

func NewProductController(uc products.Usecase) *Controller {
	return &Controller{
		productUsecase: uc,
	}
}

func (ctrl *Controller) GetProductsController(c echo.Context) error {
	ctxNative := c.Request().Context()
	data, err := ctrl.productUsecase.GetAll(ctxNative)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductListDomain(data))
}

func (ctrl *Controller) GetDetailProductController(c echo.Context) error {
	ctxNative := c.Request().Context()
	id := c.Param("productId")
	convInt, errConvInt := strconv.Atoi(id)
	if errConvInt != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, errConvInt)
	}
	product := products.Domain{
		ID: convInt,
	}
	data, err := ctrl.productUsecase.GetById(ctxNative, product)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (ctrl *Controller) CreateProductController(c echo.Context) error {
	request := requests.ProductsRequest{}
	err := c.Bind(&request)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctx := c.Request().Context()

	data, err := ctrl.productUsecase.Create(ctx, request.ToDomainProduct())
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (cl *Controller) UpdateProductController(c echo.Context) error {
	id := c.Param("productId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := requests.ProductsRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	domain := products.Domain{
		ID:       convId,
		SellerId: req.SellerId,
		Name:     req.Name,
		Price:    req.Price,
	}
	data, err := cl.productUsecase.Update(ctx, domain)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, responses.FromProductDomain(data))
}

func (cl *Controller) DeleteProductController(c echo.Context) error {
	id := c.Param("productId")
	convId, err := strconv.Atoi(id)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	product := products.Domain{
		ID: convId,
	}
	err = cl.productUsecase.Delete(ctx, product)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return deliveries.NewSuccessResponse(c, convId)
}

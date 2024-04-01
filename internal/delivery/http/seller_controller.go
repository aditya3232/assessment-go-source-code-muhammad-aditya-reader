package http

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
	"assessment-go-source-code-muhammad-aditya-reader/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SellerController struct {
	UseCase *usecase.SellerUseCase
	Log     *logrus.Logger
}

func NewSellerController(useCase *usecase.SellerUseCase, log *logrus.Logger) *SellerController {
	return &SellerController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *SellerController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateSellerRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating seller")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SellerResponse]{Data: response})
}

func (c *SellerController) List(ctx *fiber.Ctx) error {
	request := &model.SearchSellerRequest{
		NationalId: ctx.QueryInt("national_id"),
		Name:       ctx.Query("name", ""),
		Email:      ctx.Query("email", ""),
		Phone:      ctx.Query("phone", ""),
		Page:       ctx.QueryInt("page", 1),
		Size:       ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching seller")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.SellerResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *SellerController) Get(ctx *fiber.Ctx) error {
	request := &model.GetSellerRequest{
		ID: ctx.Params("sellerId"),
	}

	response, err := c.UseCase.Get(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting seller")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SellerResponse]{Data: response})
}

func (c *SellerController) Update(ctx *fiber.Ctx) error {
	request := new(model.UpdateSellerRequest)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("sellerId")

	response, err := c.UseCase.Update(ctx.Context(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating seller")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.SellerResponse]{Data: response})
}

func (c *SellerController) Delete(ctx *fiber.Ctx) error {
	sellerId := ctx.Params("sellerId")

	request := &model.DeleteSellerRequest{
		ID: sellerId,
	}

	if err := c.UseCase.Delete(ctx.Context(), request); err != nil {
		c.Log.WithError(err).Error("error deleting seller")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}

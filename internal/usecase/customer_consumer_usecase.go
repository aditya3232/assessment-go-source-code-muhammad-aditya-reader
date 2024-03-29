package usecase

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/entity"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model/converter"
	"assessment-go-source-code-muhammad-aditya-reader/internal/repository"
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerConsumerUseCase struct {
	DB                         *gorm.DB
	Log                        *logrus.Logger
	Validate                   *validator.Validate
	CustomerConsumerRepository *repository.CustomerConsumerRepository
}

func NewCustomerConsumerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	customerConsumerRepository *repository.CustomerConsumerRepository) *CustomerConsumerUseCase {
	return &CustomerConsumerUseCase{
		DB:                         db,
		Log:                        logger,
		Validate:                   validate,
		CustomerConsumerRepository: customerConsumerRepository,
	}
}

func (c *CustomerConsumerUseCase) Create(ctx context.Context, request *model.CreateCustomerConsumerRequest) (*model.CustomerConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customerConsumer := &entity.CustomerConsumer{
		ID:            uuid.New().String(),
		NationalId:    request.NationalId,
		Name:          request.Name,
		DetailAddress: request.DetailAddress,
		CreatedAt:     time.Now().UnixNano() / int64(time.Millisecond),
		UpdatedAt:     time.Now().UnixNano() / int64(time.Millisecond),
	}

	totalNationalId, err := c.CustomerConsumerRepository.CountByNationalId(tx, customerConsumer)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if totalNationalId > 0 {
		c.Log.Warnf("Customer already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	if err := c.CustomerConsumerRepository.Create(tx, customerConsumer); err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerConsumerToResponse(customerConsumer), nil
}

func (c *CustomerConsumerUseCase) Update(ctx context.Context, request *model.UpdateCustomerConsumerRequest) (*model.CustomerConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customerConsumer := new(entity.CustomerConsumer)
	if err := c.CustomerConsumerRepository.FindById(tx, customerConsumer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer consumer")
		return nil, fiber.ErrNotFound
	}

	customerConsumer.Name = request.Name
	customerConsumer.DetailAddress = request.DetailAddress

	if err := c.CustomerConsumerRepository.Update(tx, customerConsumer); err != nil {
		c.Log.WithError(err).Error("error updating customer consumer")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerConsumerToResponse(customerConsumer), nil
}

func (c *CustomerConsumerUseCase) Get(ctx context.Context, request *model.GetCustomerConsumerRequest) (*model.CustomerConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customerConsumer := new(entity.CustomerConsumer)
	if err := c.CustomerConsumerRepository.FindById(tx, customerConsumer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer consumer")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerConsumerToResponse(customerConsumer), nil
}

func (c *CustomerConsumerUseCase) Delete(ctx context.Context, request *model.DeleteCustomerConsumerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	customerConsumer := new(entity.CustomerConsumer)
	if err := c.CustomerConsumerRepository.FindById(tx, customerConsumer, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting customer consumer")
		return fiber.ErrNotFound
	}

	if err := c.CustomerConsumerRepository.Delete(tx, customerConsumer); err != nil {
		c.Log.WithError(err).Error("error deleting customer consumer")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting customer consumer")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CustomerConsumerUseCase) Search(ctx context.Context, request *model.SearchCustomerConsumerRequest) ([]model.CustomerConsumerResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	customerConsumers, total, err := c.CustomerConsumerRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting customer consumers")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting customer consumers")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CustomerConsumerResponse, len(customerConsumers))
	for i, customerConsumer := range customerConsumers {
		responses[i] = *converter.CustomerConsumerToResponse(&customerConsumer)
	}

	return responses, total, nil
}

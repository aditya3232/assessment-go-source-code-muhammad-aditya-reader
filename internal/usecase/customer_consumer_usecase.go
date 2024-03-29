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

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

type SellerUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	SellerRepository *repository.SellerRepository
}

func NewSellerUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate,
	sellerRepository *repository.SellerRepository) *SellerUseCase {
	return &SellerUseCase{
		DB:               db,
		Log:              log,
		Validate:         validate,
		SellerRepository: sellerRepository,
	}
}

func (u *SellerUseCase) Create(ctx context.Context, request *model.CreateSellerRequest) (*model.SellerResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// semuer err, akan dikirim ke log
	// sedangkan error yang direturn adalah error dari fiber
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	seller := &entity.Seller{
		ID:            uuid.New().String(),
		NationalId:    request.NationalId,
		Name:          request.Name,
		Email:         request.Email,
		Phone:         request.Phone,
		DetailAddress: request.DetailAddress,
		CreatedAt:     time.Now().UnixNano() / int64(time.Millisecond),
		UpdatedAt:     time.Now().UnixNano() / int64(time.Millisecond),
	}

	totalNationalId, err := u.SellerRepository.CountByNationalId(tx, seller)
	if err != nil {
		u.Log.Warnf("Failed count user from database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if totalNationalId > 0 {
		u.Log.Warnf("Seller already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	if err := u.SellerRepository.Create(tx, seller); err != nil {
		u.Log.WithError(err).Error("error create seller")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SellerToResponse(seller), nil
}

func (u *SellerUseCase) Update(ctx context.Context, request *model.UpdateSellerRequest) (*model.SellerResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	seller := new(entity.Seller)
	if err := u.SellerRepository.FindById(tx, seller, request.ID); err != nil {
		u.Log.WithError(err).Error("error find seller by id")
		return nil, fiber.ErrNotFound
	}

	seller.Name = request.Name
	seller.Email = request.Email
	seller.Phone = request.Phone
	seller.DetailAddress = request.DetailAddress
	seller.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)

	if err := u.SellerRepository.Update(tx, seller); err != nil {
		u.Log.WithError(err).Error("error update seller")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SellerToResponse(seller), nil
}

func (u *SellerUseCase) Get(ctx context.Context, request *model.GetSellerRequest) (*model.SellerResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	seller := new(entity.Seller)
	if err := u.SellerRepository.FindById(tx, seller, request.ID); err != nil {
		u.Log.WithError(err).Error("error find seller by id")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SellerToResponse(seller), nil
}

func (u *SellerUseCase) Delete(ctx context.Context, request *model.DeleteSellerRequest) error {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	seller := new(entity.Seller)
	if err := u.SellerRepository.FindById(tx, seller, request.ID); err != nil {
		u.Log.WithError(err).Error("error find seller by id")
		return fiber.ErrNotFound
	}

	if err := u.SellerRepository.Delete(tx, seller); err != nil {
		u.Log.WithError(err).Error("error delete seller")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *SellerUseCase) Search(ctx context.Context, request *model.SearchSellerRequest) ([]model.SellerResponse, int64, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	sellers, total, err := u.SellerRepository.Search(tx, request)
	if err != nil {
		u.Log.WithError(err).Error("error getting sellers")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.SellerResponse, len(sellers))
	for i, seller := range sellers {
		responses[i] = *converter.SellerToResponse(&seller)
	}

	return responses, total, nil
}

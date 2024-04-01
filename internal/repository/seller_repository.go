package repository

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/entity"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SellerRepository struct {
	Repository[entity.Seller]
	Log *logrus.Logger
}

func NewSellerRepository(log *logrus.Logger) *SellerRepository {
	return &SellerRepository{
		Log: log,
	}
}

func (r *SellerRepository) CountByNationalId(db *gorm.DB, seller *entity.Seller) (int64, error) {
	var total int64
	err := db.Model(seller).Where("national_id = ?", seller.NationalId).Count(&total).Error
	return total, err
}

func (r *SellerRepository) Search(db *gorm.DB, request *model.SearchSellerRequest) ([]entity.Seller, int64, error) {
	var sellers []entity.Seller
	if err := db.Scopes(r.FilterSeller(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&sellers).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Seller{}).Scopes(r.FilterSeller(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return sellers, total, nil
}

func (r *SellerRepository) FilterSeller(request *model.SearchSellerRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if request.NationalId > 0 {
			tx = tx.Where("national_id = ?", request.NationalId)
		}
		if request.Name != "" {
			tx = tx.Where("name LIKE ?", "%"+request.Name+"%")
		}
		if request.Email != "" {
			tx = tx.Where("email LIKE ?", "%"+request.Email+"%")
		}
		if request.Phone != "" {
			tx = tx.Where("phone LIKE ?", "%"+request.Phone+"%")
		}
		return tx
	}
}

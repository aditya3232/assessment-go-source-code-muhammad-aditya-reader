package repository

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/entity"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerConsumerRepository struct {
	Repository[entity.CustomerConsumer]
	Log *logrus.Logger
}

func NewCustomerConsumerRepository(log *logrus.Logger) *CustomerConsumerRepository {
	return &CustomerConsumerRepository{
		Log: log,
	}
}

func (r *CustomerConsumerRepository) CountByNationalId(db *gorm.DB, customerConsumer *entity.CustomerConsumer) (int64, error) {
	var total int64
	err := db.Model(customerConsumer).Where("national_id = ?", customerConsumer.NationalId).Count(&total).Error
	return total, err
}

func (r *CustomerConsumerRepository) Search(db *gorm.DB, request *model.SearchCustomerConsumerRequest) ([]entity.CustomerConsumer, int64, error) {
	var customerConsumers []entity.CustomerConsumer
	if err := db.Scopes(r.FilterCustomer(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&customerConsumers).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.CustomerConsumer{}).Scopes(r.FilterCustomer(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return customerConsumers, total, nil
}

func (r *CustomerConsumerRepository) FilterCustomer(request *model.SearchCustomerConsumerRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if national_id := request.NationalId; national_id != 0 {
			tx = tx.Where("national_id = ?", national_id)
		}

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("name LIKE ?", name)
		}

		return tx
	}
}

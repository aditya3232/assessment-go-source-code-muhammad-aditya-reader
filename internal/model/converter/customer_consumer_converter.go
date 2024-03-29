package converter

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/entity"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
)

func CustomerConsumerToResponse(customer *entity.CustomerConsumer) *model.CustomerConsumerResponse {
	return &model.CustomerConsumerResponse{
		ID:            customer.ID,
		NationalId:    customer.NationalId,
		Name:          customer.Name,
		DetailAddress: customer.DetailAddress,
		CreatedAt:     customer.CreatedAt,
		UpdatedAt:     customer.UpdatedAt,
	}
}

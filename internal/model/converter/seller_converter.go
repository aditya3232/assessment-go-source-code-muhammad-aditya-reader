package converter

import (
	"assessment-go-source-code-muhammad-aditya-reader/internal/entity"
	"assessment-go-source-code-muhammad-aditya-reader/internal/model"
)

func SellerToResponse(seller *entity.Seller) *model.SellerResponse {
	return &model.SellerResponse{
		ID:            seller.ID,
		NationalId:    seller.NationalId,
		Name:          seller.Name,
		Email:         seller.Email,
		Phone:         seller.Phone,
		DetailAddress: seller.DetailAddress,
		CreatedAt:     seller.CreatedAt,
		UpdatedAt:     seller.UpdatedAt,
	}
}

package model

type SellerResponse struct {
	ID            string `json:"id"`
	NationalId    int64  `json:"national_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	DetailAddress string `json:"detail_address"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type CreateSellerRequest struct {
	NationalId    int64  `json:"national_id" validate:"required"`
	Name          string `json:"name" validate:"required,max=255"`
	Email         string `json:"email" validate:"required,email,max=255"`
	Phone         string `json:"phone" validate:"required,max=15"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type UpdateSellerRequest struct {
	ID            string `json:"-" validate:"required,max=100,uuid"`
	Name          string `json:"name" validate:"required,max=255"`
	Email         string `json:"email" validate:"required,email,max=255"`
	Phone         string `json:"phone" validate:"required,max=15"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type SearchSellerRequest struct {
	NationalId int    `json:"national_id"`
	Name       string `json:"name" validate:"max=255"`
	Email      string `json:"email" validate:"max=255"`
	Phone      string `json:"phone" validate:"max=15"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type GetSellerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteSellerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

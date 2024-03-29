package model

type CustomerConsumerResponse struct {
	ID            string `json:"id"`
	NationalId    int64  `json:"national_id"`
	Name          string `json:"name"`
	DetailAddress string `json:"detail_address"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

type CreateCustomerConsumerRequest struct {
	NationalId    int64  `json:"national_id" validate:"required"`
	Name          string `json:"name" validate:"required,max=255"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type UpdateCustomerConsumerRequest struct {
	ID            string `json:"-" validate:"required,max=100,uuid"`
	Name          string `json:"name" validate:"required,max=255"`
	DetailAddress string `json:"detail_address" validate:"required,max=255"`
}

type SearchCustomerConsumerRequest struct {
	NationalId int    `json:"national_id"`
	Name       string `json:"name" validate:"max=255"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type GetCustomerConsumerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteCustomerConsumerRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

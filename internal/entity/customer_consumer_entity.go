package entity

type CustomerConsumer struct {
	ID            string `gorm:"column:id;primaryKey"`
	NationalId    int64  `gorm:"column:national_id"`
	Name          string `gorm:"column:name"`
	DetailAddress string `gorm:"column:detail_address"`
	CreatedAt     int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt     int64  `gorm:"column:updated_at;autoUpdateTime:milli"`
}

func (c *CustomerConsumer) TableName() string {
	return "customer_consumers"
}

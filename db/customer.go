package db

type CustomerType string

const (
	CustomerTypeIndividual CustomerType = "INDIVIDUAL"
	CustomerTypeBusiness   CustomerType = "BUSINESS"
)

type Customer struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       string       `gorm:"unique"`
	Type        CustomerType `gorm:"type:varchar(20);default:'INDIVIDUAL'"`
	CompanyName *string      `gorm:"type:varchar(255)"` // For business customers only
}

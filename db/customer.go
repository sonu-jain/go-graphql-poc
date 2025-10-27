package db

import "time"

type CustomerType string
type CustomerStatus string

const (
	CustomerTypeIndividual CustomerType = "INDIVIDUAL"
	CustomerTypeBusiness   CustomerType = "BUSINESS"
	CustomerTypePremium    CustomerType = "PREMIUM"
)

const (
	CustomerStatusActive    CustomerStatus = "ACTIVE"
	CustomerStatusInactive  CustomerStatus = "INACTIVE"
	CustomerStatusSuspended CustomerStatus = "SUSPENDED"
	CustomerStatusPending   CustomerStatus = "PENDING"
)

type Customer struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       string         `gorm:"unique"`
	Password    string         `gorm:"type:varchar(255)"` // Hashed password
	Type        CustomerType   `gorm:"type:varchar(20);default:'INDIVIDUAL'"`
	Status      CustomerStatus `gorm:"type:varchar(20);default:'ACTIVE'"`
	CompanyName *string        `gorm:"type:varchar(255)"` // For business customers
	PremiumTier *string        `gorm:"type:varchar(50)"`  // For premium customers

	// Personal info for individual customers
	Phone       *string `gorm:"type:varchar(20)"`
	Address     *string `gorm:"type:text"`
	DateOfBirth *string `gorm:"type:date"`

	// Business info for business customers
	TaxID         *string `gorm:"type:varchar(50)"`
	Industry      *string `gorm:"type:varchar(100)"`
	EmployeeCount *int    `gorm:"type:int"`
	Website       *string `gorm:"type:varchar(255)"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

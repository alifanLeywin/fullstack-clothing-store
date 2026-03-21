package models

type User struct{
	BaseModel

	Name              string `gorm:"not null" json:"name"`
	Email             string `gorm:"uniqueIndex;not null" json:"email"`
	Password          string `gorm:"not null" json:"-"`
	Role              string `gorm:"type:varchar(20);default:'user'" json:"role"`
	IsVerified        bool   `gorm:"default:false" json:"is_verified"`
	VerificationCode  string `gorm:"type:varchar(6);" json:"-"`
	
}
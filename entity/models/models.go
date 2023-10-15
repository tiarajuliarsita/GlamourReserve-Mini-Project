package models

import (
	"errors"
	"glamour_reserve/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"not null;primary key"`
	UserName  string `gorm:"not null" valid:"required~your user name is required"`
	Email     string `gorm:"not null;unique" valid:"required~your email is required, email~invalid email format"`
	Password  string `gorm:"not null" valid:"required~your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Phone     string `gorm:"not null" valid:"required~your phone is required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	u.ID = newUuid.String()
	if len(u.Password) < 6 {
		return errors.New("password must have a minimum length of 6 characters")
	}
	u.Password, _ = helpers.HassPass(u.Password)
	return nil
}

type Service struct {
	ID          string `gorm:"not null;primary key" json:"id" form:"id"`
	Name        string `gorm:"not null;unique" valid:"required~your name is required" json:"name" form:"name"`
	Description string `gorm:"not null" valid:"required~your description is required" json:"description" form:"description"`
	// Image       string `gorm:"not null" valid:"required~your image is required" json:"image" form:"image"`
	Variants  []Variant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"variants"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	s.ID = newUuid.String()

	return nil
}

type Variant struct {
	ID          string `gorm:"not null;primary key"`
	Name        string `gorm:"not null" valid:"required~your name is required"`
	Description string `gorm:"not null;unique" valid:"required~your description is required"`
	Price       int    `gorm:"not null" valid:"required~your price is required"`
	ServiceID   string //`gorm:"not null" valid:"required~your service id is required"`
	Service     *Service
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	ID              string `gorm:"not null;primary key"`
	UserID          string `gorm:"not null" valid:"required~your user id is required"`
	InvoiceNumb     string `gorm:"not null;unique" valid:"required~your invoice is required"`
	details_booking []DetailBooking
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
type DetailBooking struct {
	ID        string `gorm:"not null;primary key"`
	Date      string `gorm:"not null" valid:"required~your date  is required"`
	Time      string `gorm:"not null" valid:"required~your time is required"`
	BookingID string
	VariantID string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

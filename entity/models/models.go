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
	Role      string `gorm:"type:ENUM('user', 'admin');not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Service struct {
	ID          string `gorm:"not null;type:varchar(255);primary key" json:"id" form:"id"`
	Name        string `gorm:"not null;unique" valid:"required~your name is required" json:"name" form:"name"`
	Description string `gorm:"not null" valid:"required~your description is required" json:"description" form:"description"`
	// Image       string `gorm:"not null" valid:"required~your image is required" json:"image" form:"image"`
	Price     int `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	ID             string `gorm:"not null;primary key; type:varchar(255)"`
	UserID         string `valid:"required~your user id is required"`
	InvoiceNumb    string `gorm:"not null;unique" valid:"required~your invoice is required"`
	Total          int    `gorm:"not null"`
	Status         string `gorm:"type:ENUM('pending', 'done');not null;default:'pending'"`
	DetailsBooking []DetailBooking
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type DetailBooking struct {
	ID        string `gorm:"not null;primary key"`
	Date      string `gorm:"not null" valid:"required~your date  is required"`
	Time      string `gorm:"not null" valid:"required~your time is required"`
	BookingID string `gorm:"type:varchar(255)"`
	ServiceID string `gorm:"type:varchar(255)"`
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

func (s *Service) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	s.ID = newUuid.String()

	return nil
}

func (b *Booking) BeforeCreate(tx *gorm.DB) (err error) {

	_, err = govalidator.ValidateStruct(b)
	if err != nil {
		return err
	}
	newUuid := uuid.New()
	b.ID = newUuid.String()

	return nil
}

func (d *DetailBooking) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(d)
	if err != nil {
		return err
	}

	newUuid := uuid.New()
	d.ID = newUuid.String()

	return nil
}

package models

import (
	"errors"
	"glamour_reserve/utils/helpers"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
	d.DateTime = d.DateTime.Local()
	return nil
}

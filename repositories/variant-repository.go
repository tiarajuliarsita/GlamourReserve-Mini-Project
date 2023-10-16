package repositories

import (
	"database/sql"
	"errors"
	"glamour_reserve/entity/core"
	"glamour_reserve/entity/models"

	"gorm.io/gorm"
)

type VariantRepoInterface interface {
	Create(variant core.VariantCore) (core.VariantCore, error)
}

type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) *variantRepository {
	return &variantRepository{db}
}

func (r *variantRepository) Create(variant core.VariantCore) (core.VariantCore, error) {
	dataVariant := core.VariantCore{}
	service := models.Service{}

	err := r.db.Where("id = ?", variant.ServiceID).Table("services").First(&service).Error
	// err := r.db.Where("id = ?", variant.ServiceID).Table("services").First(&service).Error

	if err != nil {
		if err == sql.ErrNoRows {
			return dataVariant,errors.New("service not found")
		}

		return dataVariant,err
	}

	variantInsert := core.VariantCoreToVariantModel(variant)

	err = r.db.Create(&variantInsert).Error
	if err != nil {
		return dataVariant, err
	}

	dataVariant = core.VariantModelToVariantCore(variantInsert)
	return dataVariant, nil
}

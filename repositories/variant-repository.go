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
	FindByID(id string) (core.VariantCore, error)
	FindAll() ([]core.VariantCore, error)
	Delete(id string) error
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
	if err != nil {
		if err == sql.ErrNoRows {
			return dataVariant, errors.New("service not found")
		}
		return dataVariant, err
	}

	variantInsert := core.VariantCoreToVariantModel(variant)
	err = r.db.Create(&variantInsert).Error
	if err != nil {
		return dataVariant, err
	}

	dataVariant = core.VariantModelToVariantCore(variantInsert)
	return dataVariant, nil
}

func (r *variantRepository) FindByID(id string) (core.VariantCore, error) {
	variant := models.Variant{}
	dataVariant := core.VariantCore{}

	err := r.db.Where("id = ?", id).First(&variant).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return dataVariant, errors.New("variant not found")
		}
		return dataVariant, err
	}

	dataVariant = core.VariantModelToVariantCore(variant)
	return dataVariant, err
}

func (r *variantRepository) FindAll() ([]core.VariantCore, error) {
	variants := []models.Variant{}
	dataVariant := []core.VariantCore{}

	err := r.db.Find(&variants).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return dataVariant, errors.New("variant not found")
		}
		return dataVariant, err
	}

	for _, v := range variants {
		variant := core.VariantModelToVariantCore(v)
		dataVariant = append(dataVariant, variant)
	}

	return dataVariant, err
}

func (r *variantRepository) Delete(id string) error {
	variant := models.Variant{}
	err := r.db.Where("id = ?", id).Delete(&variant).Error
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("variant not found")
		}
		return  err
	}

	return nil
}

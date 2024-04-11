package model

import (
	"github.com/go-playground/validator/v10"
)

type BannerPostRequest struct {

	// Идентификаторы тэгов
	TagIds []int32 `json:"tag_ids,omitempty" validate:"required,min=1"`

	// Идентификатор фичи
	FeatureId int32 `json:"feature_id,omitempty" validate:"required,min=1"`

	// Содержимое баннера
	Content map[string]interface{} `json:"content,omitempty"`

	// Флаг активности баннера
	IsActive bool `json:"is_active,omitempty"`
}

// AssertBannerPostRequestRequired checks if the required fields are not zero-ed
func AssertBannerPostRequestRequired(obj BannerPostRequest) error {
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return err
	}
	return nil
}

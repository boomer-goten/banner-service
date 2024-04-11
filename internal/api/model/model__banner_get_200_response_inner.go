package model

import (
	"time"
)

type BannerGet200ResponseInner struct {

	// Идентификатор баннера
	BannerId int32 `json:"banner_id,omitempty"`

	// Идентификаторы тэгов
	TagIds []int32 `json:"tag_ids,omitempty"`

	// Идентификатор фичи
	FeatureId int32 `json:"feature_id,omitempty"`

	// Содержимое баннера
	Content map[string]interface{} `json:"content,omitempty"`

	// Флаг активности баннера
	IsActive bool `json:"is_active,omitempty"`

	// Дата создания баннера
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Дата обновления баннера
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BannerGet200ResponseContent struct {
	Content map[string]interface{} `json:"content,omitempty"`
}

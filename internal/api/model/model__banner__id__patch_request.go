package model

import "errors"

type BannerIdPatchRequest struct {

	// Идентификаторы тэгов
	TagIds *[]int32 `json:"tag_ids,omitempty"`

	// Идентификатор фичи
	FeatureId *int32 `json:"feature_id,omitempty"`

	// Содержимое баннера
	Content *map[string]interface{} `json:"content,omitempty"`

	// Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`
}

// AssertBannerIdPatchRequestRequired checks if the required fields are not zero-ed
func AssertBannerIdPatchRequestRequired(obj BannerIdPatchRequest) error {
	if obj.TagIds == nil && obj.FeatureId == nil && obj.Content == nil && obj.IsActive == nil {
		return errors.New("you must specify at least one required parameter")
	}
	return nil
}

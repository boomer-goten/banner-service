package model

import (
	"banner-server/internal/api/model"
	"banner-server/pkg/repository"

	"gorm.io/gorm"
)

type BannerTag struct {
	Tags      Tag     `gorm:"foreignKey:TagID"`
	Features  Feature `gorm:"foreignKey:FeatureID"`
	Banners   Banner  `gorm:"foreignKey:BannerID"`
	BannerID  int     `gorm:"index"`
	TagID     int     `gorm:"primaryKey;autoIncrement:false"`
	FeatureID int     `gorm:"primaryKey;autoIncrement:false"`
}

func (BannerTag) TableName() string {
	return "banner_tags"
}

func ConvertPostRequestTags(data *model.BannerPostRequest) []BannerTag {
	slice := make([]BannerTag, 0, len(data.TagIds))
	for _, val := range data.TagIds {
		slice = append(slice, BannerTag{
			TagID:     int(val),
			FeatureID: int(data.FeatureId),
		})
	}
	return slice
}

func ConvertPatchRequestTags(id int, data *model.BannerIdPatchRequest) ([]BannerTag, error) {
	if data.TagIds == nil && data.FeatureId == nil {
		return nil, repository.ErrPatchBannerTags
	}
	if data.TagIds == nil {
		slice := make([]BannerTag, 0, 1)
		slice = append(slice, BannerTag{
			BannerID:  id,
			TagID:     0,
			FeatureID: int(*data.FeatureId),
		})
		data.FeatureId = nil
		return slice, repository.ErrPatchTags
	}
	var err error
	slice := make([]BannerTag, 0, len(*data.TagIds))
	for _, val := range *data.TagIds {
		if data.FeatureId == nil {
			slice = append(slice, BannerTag{
				BannerID:  id,
				TagID:     int(val),
				FeatureID: 0,
			})
			err = repository.ErrPatchFeature
		} else {
			slice = append(slice, BannerTag{
				BannerID:  id,
				TagID:     int(val),
				FeatureID: int(*data.FeatureId),
			})
			err = nil
		}
	}
	data.FeatureId = nil
	data.TagIds = nil
	return slice, err
}

func (b *BannerTag) AfterDelete(tx *gorm.DB) (err error) {
	tx.Model(Banner{}).Where("banner_id = ?", b.BannerID).Delete(Banner{})
	return nil
}

package repository

import (
	request "banner-server/internal/api/model"
	"errors"
)

var (
	ErrCreateItem      = errors.New("faile to create banner")
	ErrIndex           = errors.New("current tag or feature doesn't exist or banner with current tag and feature already exist")
	ErrDb              = errors.New("internal database error")
	ErrDeleteItem      = errors.New("faile to delete banner")
	ErrDeleteFind      = errors.New("banner to delete doesn't exist")
	ErrSelectItem      = errors.New("current banner doesn't exist")
	ErrPatchItem       = errors.New("faile to update banner")
	ErrPatchBannerTags = errors.New("do not update banner_tags")
	ErrPatchTags       = errors.New("update only featureID")
	ErrPatchFeature    = errors.New("update only tags")
	ErrFoundItem       = errors.New("banner doens't exist")
)

type Repository interface {
	BannerGet(tagID, featureID, offset, limit int, role string) ([]request.BannerGet200ResponseInner, error)
	BannerIdDelete(bannerID int) error
	BannerIdPatch(id int, data *request.BannerIdPatchRequest) error
	BannerPost(data *request.BannerPostRequest) (int, error)
	UserBannerGet(tag, feature int) ([]byte, bool, error)
}

package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"banner-server/internal/api/middleware/auth"
	"banner-server/internal/api/model"
	"banner-server/pkg/cache"
	key "banner-server/pkg/cache/cacheInMemory/model"
	"banner-server/pkg/repository"

	"github.com/golang-jwt/jwt"
)

type DefaultAPIService struct {
	db    repository.Repository
	cache cache.Cache
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService(db repository.Repository, cache cache.Cache) DefaultAPIServicer {
	return &DefaultAPIService{
		db:    db,
		cache: cache,
	}
}

// BannerGet - Получение всех баннеров c фильтрацией по фиче и/или тегу
func (s *DefaultAPIService) BannerGet(ctx context.Context, featureId int32, tagId int32, limit int32, offset int32) (ImplResponse, error) {
	values, err := s.db.BannerGet(int(tagId), int(featureId), int(offset), int(limit), ctx.Value(auth.ContextRoleKey).(string))
	if err != nil {
		return Response(http.StatusInternalServerError, model.UserBannerGet400Response{Error: err.Error()}), nil
	}
	return Response(http.StatusOK, values), nil
}

// BannerIdDelete - Удаление баннера по идентификатору
func (s *DefaultAPIService) BannerIdDelete(ctx context.Context, id int32) (ImplResponse, error) {
	err := s.db.BannerIdDelete(int(id))
	if err == repository.ErrDb {
		return Response(http.StatusInternalServerError, model.UserBannerGet400Response{
			Error: err.Error(),
		}), nil
	} else if err == repository.ErrDeleteFind {
		return Response(http.StatusNotFound, nil), nil
	} else if err != nil {
		return Response(http.StatusBadRequest, model.UserBannerGet400Response{Error: err.Error()}), err
	}
	return Response(http.StatusNoContent, nil), nil
}

// BannerIdPatch - Обновление содержимого баннера
func (s *DefaultAPIService) BannerIdPatch(ctx context.Context, id int32, bannerIdPatchRequest model.BannerIdPatchRequest) (ImplResponse, error) {
	err := s.db.BannerIdPatch(int(id), &bannerIdPatchRequest)
	if err == repository.ErrDb {
		return Response(http.StatusInternalServerError, model.UserBannerGet400Response{
			Error: err.Error(),
		}), nil
	} else if err == repository.ErrFoundItem {
		return Response(http.StatusNotFound, nil), nil
	} else if err != nil {
		return Response(http.StatusBadRequest, model.UserBannerGet400Response{Error: err.Error()}), nil
	}
	return Response(http.StatusOK, nil), nil
}

// BannerPost - Создание нового баннера
func (s *DefaultAPIService) BannerPost(ctx context.Context, bannerPostRequest model.BannerPostRequest) (ImplResponse, error) {
	id, err := s.db.BannerPost(&bannerPostRequest)
	if err != nil {
		return Response(http.StatusBadRequest, err), err
	}
	s.cache.Set(key.KeyCache{
		FeatureID: int(bannerPostRequest.FeatureId),
		TagID:     int(bannerPostRequest.TagIds[0]),
	}, key.ValueCache{
		IsActive: bannerPostRequest.IsActive,
		Banner: key.ContentBanner{
			Content: bannerPostRequest.Content,
		},
	})
	return Response(http.StatusCreated, model.BannerPost201Response{
		BannerId: int32(id),
	}), nil
}

// UserBannerGet - Получение баннера для пользователя
func (s *DefaultAPIService) UserBannerGet(ctx context.Context, tagId int32, featureId int32, useLastRevision bool) (ImplResponse, error) {
	if !useLastRevision {
		request, ok := s.cache.Get(key.KeyCache{
			FeatureID: int(featureId),
			TagID:     int(tagId),
		})
		if ok {
			content := request.(key.ValueCache)
			if ctx.Value(auth.ContextRoleKey) != "admin" && !content.IsActive {
				return Response(http.StatusForbidden, nil), nil
			}
			return Response(http.StatusOK, content.Banner), nil
		}
	}
	content, isActive, err := s.db.UserBannerGet(int(tagId), int(featureId))
	if ctx.Value(auth.ContextRoleKey) != "admin" && !isActive {
		return Response(http.StatusForbidden, nil), nil
	}
	if err == repository.ErrDb {
		return Response(http.StatusInternalServerError, errors.New("error db processing")), nil
	}
	if err != nil {
		return Response(http.StatusNotFound, nil), nil
	}
	banner := model.BannerGet200ResponseContent{}
	json.Unmarshal(content, &banner.Content)
	s.cache.Set(key.KeyCache{
		FeatureID: int(featureId),
		TagID:     int(tagId),
	}, key.ValueCache{
		IsActive: isActive,
		Banner: key.ContentBanner{
			Content: banner.Content,
		},
	})
	return Response(http.StatusOK, banner), nil
}

// TokenGet - Получение токена
func (s *DefaultAPIService) TokenGet(ctx context.Context) (ImplResponse, error) {
	admin, errAd := createToken("admin")
	user, errUs := createToken("user")
	if errAd != nil || errUs != nil {
		return Response(http.StatusInternalServerError, errors.New("error creating token")), nil
	}
	return Response(http.StatusCreated, model.TokenGet201Response{
		Admin: admin,
		User:  user,
	}), nil
}

func createToken(role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["role"] = role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

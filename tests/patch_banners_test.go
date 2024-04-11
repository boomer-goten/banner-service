package tests

import (
	"banner-server/internal/api/model"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestPatchStatusOK(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "PATCH with status 200",
	})
	var featureID int32 = 5
	active := true
	e.PATCH("/banner/{id}").WithPath("id", 1).
		WithJSON(model.BannerIdPatchRequest{
			TagIds:    &[]int32{2},
			FeatureId: &featureID,
			IsActive:  &active,
		}).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK)
}

func TestPatchStatusUnauthorized(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "PATCH with status 401",
	})
	var featureID int32 = 5
	active := true
	e.PATCH("/banner/{id}").WithPath("id", 1).
		WithJSON(model.BannerIdPatchRequest{
			TagIds:    &[]int32{2},
			FeatureId: &featureID,
			IsActive:  &active,
		}).
		WithHeader("token", "hello world").
		Expect().Status(http.StatusUnauthorized)
}

func TestPatchStatusNotAccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "PATCH with status 403",
	})
	var featureID int32 = 5
	active := true
	e.PATCH("/banner/{id}").WithPath("id", 1).
		WithJSON(model.BannerIdPatchRequest{
			TagIds:    &[]int32{2},
			FeatureId: &featureID,
			IsActive:  &active,
		}).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk").
		Expect().Status(http.StatusForbidden)
}

func TestPatchStatusNotFound(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "PATCH with status 404",
	})

	e.PATCH("/banner/{id}").WithPath("id", 100).
		WithJSON(model.BannerIdPatchRequest{
			TagIds: &[]int32{2},
		}).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusNotFound)
}

func TestPatchStatusIncorrectData(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "PATCH with status 403",
	})
	e.PATCH("/banner/{id}").WithPath("id", 100).
		WithQuery("tag_id", 5).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusBadRequest)
}

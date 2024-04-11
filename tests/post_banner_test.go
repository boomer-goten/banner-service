package tests

import (
	"banner-server/internal/api/model"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestPostStatusOK_1(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 201",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 2, 3, 4},
		FeatureId: 1,
		Content: map[string]interface{}{
			"title": "hello, i'm title from 1 post test",
			"text":  "i have four tags",
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsKey("banner_id")
}

func TestPostStatusOK_2(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 201",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{5, 6},
		FeatureId: 1,
		Content: map[string]interface{}{
			"title": "Alan Rickman",
			"film":  "Harry Potter",
			"part":  "and the prisoner of azkaban",
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsKey("banner_id")
}

func TestPostStatusOK_3(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 201",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 2, 3, 4},
		FeatureId: 2,
		Content: map[string]interface{}{
			"title": "Warcraft 3",
			"text":  "Frozen throne",
			"rate":  10,
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsKey("banner_id")
}

func TestPostStatusOK_4(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 201",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{5, 6},
		FeatureId: 2,
		Content: map[string]interface{}{
			"group":    "The doors",
			"song":     "Riders on the storm",
			"describe": "classic",
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsKey("banner_id")
}

func TestPostStatusOK_5(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 201",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 3, 4},
		FeatureId: 3,
		Content: map[string]interface{}{
			"text": "some_text",
		},
		IsActive: false,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusCreated).JSON().
		Object().ContainsKey("banner_id")
}

func TestPostStatusInvalidData(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 400",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 3, 4},
		FeatureId: 3,
		Content: map[string]interface{}{
			"text": "some_text",
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusBadRequest)
}

func TestPostStatusUnauthorized(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 401",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 3, 4},
		FeatureId: 3,
		Content: map[string]interface{}{
			"text": "some_text",
		},
		IsActive: true,
	}).WithHeader("token", "bad token").Expect().Status(http.StatusUnauthorized)
}

func TestPostStatusNoAccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "POST with status 403",
	})

	e.POST("/banner").WithJSON(model.BannerPostRequest{
		TagIds:    []int32{1, 3, 4},
		FeatureId: 3,
		Content: map[string]interface{}{
			"text": "some_text",
		},
		IsActive: true,
	}).WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk").
		Expect().Status(http.StatusForbidden)
}

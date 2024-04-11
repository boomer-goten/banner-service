package tests

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestGetBannersStatusOK_1(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("tag_id", 2).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(2)
}

func TestGetBannersStatusOK_2(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("tag_id", 2).
		WithQuery("limit", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_3(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("tag_id", 2).
		WithQuery("offset", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_4(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("tag_id", 2).
		WithQuery("limit", 5).
		WithQuery("offset", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_5(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(2)
}

func TestGetBannersStatusOK_6(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("limit", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_7(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("limit", 100).
		WithQuery("offset", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_8(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("tag_id", 2).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_9(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("tag_id", 2).
		WithQuery("limit", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(1)
}

func TestGetBannersStatusOK_10(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("tag_id", 2).
		WithQuery("limit", 1).
		WithQuery("offset", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().Array().Length().IsEqual(0)
}

func TestGetBannersStatusUnauthorized(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 401",
	})

	e.GET("/banner").
		WithQuery("feature_id", 2).
		WithQuery("tag_id", 2).
		WithQuery("limit", 1).
		WithQuery("offset", 1).
		WithHeader("token", "bad").
		Expect().Status(http.StatusUnauthorized)
}

package tests

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestGetStatusOK(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 200",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 4).
		WithQuery("feature_id", 2).
		WithQuery("use_last_revision", true).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusOK).JSON().
		Object().ContainsKey("content")
}

func TestGetStatusInvalidTag(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 400",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", []int{1, 2}).
		WithQuery("feature_id", 2000).
		WithQuery("use_last_revision", true).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusBadRequest).JSON().IsEqual("invalid Data")
}

func TestGetStatusInvalidFeature(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 400",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 2).
		WithQuery("feature_id", false).
		WithQuery("use_last_revision", true).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusBadRequest).JSON().IsEqual("invalid Data")
}

func TestGetStatusInvalidRevision(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 400",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 2).
		WithQuery("feature_id", 5).
		WithQuery("use_last_revision", 2.5).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusBadRequest).JSON().IsEqual("invalid Data")
}

func TestGetStatusInvalidToken(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 401",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 2).
		WithQuery("feature_id", 5).
		WithQuery("use_last_revision", false).
		WithHeader("token", "hello world").
		Expect().Status(http.StatusUnauthorized)
}

func TestGetStatusWeakToken(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 403",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 1).
		WithQuery("feature_id", 3).
		WithQuery("use_last_revision", true).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk").
		Expect().Status(http.StatusForbidden)
}

func TestGetStatusNotFound(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "GET with status 404",
	})

	e.GET("/user_banner").
		WithQuery("tag_id", 2).
		WithQuery("feature_id", 5).
		WithQuery("use_last_revision", true).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk").
		Expect().Status(http.StatusNotFound)
}

// TODO POST method after all others

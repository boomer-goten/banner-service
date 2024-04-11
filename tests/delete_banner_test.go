package tests

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestDeleteStatusOK(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "DELETE with status 204",
	})
	e.DELETE("/banner/{id}").WithPath("id", 1).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusNoContent)
}

func TestDeleteStatusNotFound(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "DELETE with status 404",
	})
	e.DELETE("/banner/{id}").WithPath("id", 101).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.124iY15VoTfoX5zua5CKorLT5Kjl-jxW5B7fB8tENWI").
		Expect().Status(http.StatusNotFound)
}

func TestDeleteStatusUnathorized(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "DELETE with status 401",
	})
	e.DELETE("/banner/{id}").WithPath("id", 101).
		WithHeader("token", "admin").
		Expect().Status(http.StatusUnauthorized)
}

func TestDeleteStatusNotAccess(t *testing.T) {
	e := httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost:8080",
		Reporter: httpexpect.NewAssertReporter(t),
		TestName: "DELETE with status 403",
	})
	e.DELETE("/banner/{id}").WithPath("id", 101).
		WithHeader("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoidXNlciJ9.9tRuwXmCywT0xIJMAFfgSW9XjsGzaFsxIbbgV3gg9Nk").
		Expect().Status(http.StatusForbidden)
}

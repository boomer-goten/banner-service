package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"banner-server/internal/api/model"

	"github.com/gorilla/mux"
)

var ErrData = errors.New("invalid Data")

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service      DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) Router {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"BannerGet": Route{
			strings.ToUpper("Get"),
			"/banner",
			c.BannerGet,
		},
		"BannerIdDelete": Route{
			strings.ToUpper("Delete"),
			"/banner/{id}",
			c.BannerIdDelete,
		},
		"BannerIdPatch": Route{
			strings.ToUpper("Patch"),
			"/banner/{id}",
			c.BannerIdPatch,
		},
		"BannerPost": Route{
			strings.ToUpper("Post"),
			"/banner",
			c.BannerPost,
		},
		"UserBannerGet": Route{
			strings.ToUpper("Get"),
			"/user_banner",
			c.UserBannerGet,
		},
		"TokenGet": Route{
			strings.ToUpper("Get"),
			"/token",
			c.TokenGet,
		},
	}
}

// BannerGet - Получение всех баннеров c фильтрацией по фиче и/или тегу
func (c *DefaultAPIController) BannerGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var featureIdParam int32
	if query.Has("feature_id") {
		param, err := parseNumericParameter[int32](
			query.Get("feature_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		featureIdParam = param
	}
	var tagIdParam int32
	if query.Has("tag_id") {
		param, err := parseNumericParameter[int32](
			query.Get("tag_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		tagIdParam = param
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}
		offsetParam = param
	}
	result, err := c.service.BannerGet(r.Context(), featureIdParam, tagIdParam, limitParam, offsetParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// BannerIdDelete - Удаление баннера по идентификатору
func (c *DefaultAPIController) BannerIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseNumericParameter[int32](
		params["id"],
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: errors.New("to remove a banner you need to specify the ID as a number")}, nil)
		return
	}
	result, err := c.service.BannerIdDelete(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// BannerIdPatch - Обновление содержимого баннера
func (c *DefaultAPIController) BannerIdPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam, err := parseNumericParameter[int32](
		params["id"],
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bannerIdPatchRequestParam := model.BannerIdPatchRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bannerIdPatchRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: errors.New("you must specify at least one parameter")}, nil)
		return
	}
	if err := model.AssertBannerIdPatchRequestRequired(bannerIdPatchRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.BannerIdPatch(r.Context(), idParam, bannerIdPatchRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// BannerPost - Создание нового баннера
func (c *DefaultAPIController) BannerPost(w http.ResponseWriter, r *http.Request) {
	bannerPostRequestParam := model.BannerPostRequest{}
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&bannerPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: errors.New("need to pass the banner parameters in the body of the request")}, nil)
		return
	}
	if err := model.AssertBannerPostRequestRequired(bannerPostRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: errors.New("data not validated")}, nil)
		return
	}
	result, err := c.service.BannerPost(r.Context(), bannerPostRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UserBannerGet - Получение баннера для пользователя
func (c *DefaultAPIController) UserBannerGet(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var tagIdParam int32
	if query.Has("tag_id") {
		param, err := parseNumericParameter[int32](
			query.Get("tag_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: ErrData}, nil)
			return
		}

		tagIdParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "tag_id"}, nil)
		return
	}
	var featureIdParam int32
	if query.Has("feature_id") {
		param, err := parseNumericParameter[int32](
			query.Get("feature_id"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: ErrData}, nil)
			return
		}

		featureIdParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "feature_id"}, nil)
		return
	}
	var useLastRevisionParam bool
	if query.Has("use_last_revision") {
		param, err := parseBoolParameter(
			query.Get("use_last_revision"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: ErrData}, nil)
			return
		}

		useLastRevisionParam = param
	} else {
		var param bool = false
		useLastRevisionParam = param
	}
	result, err := c.service.UserBannerGet(r.Context(), tagIdParam, featureIdParam, useLastRevisionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// TokenGet - Получение токена
func (c *DefaultAPIController) TokenGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.TokenGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

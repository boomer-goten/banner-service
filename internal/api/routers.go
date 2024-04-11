package api

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"banner-server/internal/api/middleware/auth"
	"banner-server/internal/api/middleware/logger"

	"github.com/gorilla/mux"
)

// A Route defines the parameters for an api endpoint
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a map of defined api endpoints
type Routes map[string]Route

// Router defines the required methods for retrieving api routes
type Router interface {
	Routes() Routes
}

const errMsgRequiredMissing = "required parameter is missing"
const errMsgMinValueConstraint = "provided parameter is not respecting minimum value constraint"
const errMsgMaxValueConstraint = "provided parameter is not respecting maximum value constraint"

// NewRouter creates a new router for any number of api routers
func NewRouter(log *slog.Logger, routers ...Router) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logger.LoggerMiddleware(log))
	router.Use(auth.AuthMiddleware)
	log.Info("auth middleware enabled")
	for _, api := range routers {
		for name, route := range api.Routes() {
			handler := route.HandlerFunc
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(name).
				Handler(handler)
		}
	}

	return router
}

// EncodeJSONResponse uses the json encoder to write an interface to the http response with an optional status code
func EncodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	wHeader := w.Header()

	f, ok := i.(*os.File)

	if ok {
		data, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		wHeader.Set("Content-Type", http.DetectContentType(data))
		wHeader.Set("Content-Disposition", "attachment; filename="+f.Name())
		if status != nil {
			w.WriteHeader(*status)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		_, err = w.Write(data)
		return err
	}
	wHeader.Set("Content-Type", "application/json; charset=UTF-8")

	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if i != nil {
		return json.NewEncoder(w).Encode(i)
	}

	return nil
}

type Number interface {
	~int32 | ~int64 | ~float32 | ~float64
}

type ParseString[T Number | string | bool] func(v string) (T, error)

// parseInt32 parses a string parameter to an int32.
func parseInt32(param string) (int32, error) {
	if param == "" {
		return 0, nil
	}

	val, err := strconv.ParseInt(param, 10, 32)
	return int32(val), err
}

// parseBool parses a string parameter to an bool.
func parseBool(param string) (bool, error) {
	if param == "" {
		return false, nil
	}

	return strconv.ParseBool(param)
}

type Operation[T Number | string | bool] func(actual string) (T, bool, error)

func WithRequire[T Number | string | bool](parse ParseString[T]) Operation[T] {
	var empty T
	return func(actual string) (T, bool, error) {
		if actual == "" {
			return empty, false, errors.New(errMsgRequiredMissing)
		}

		v, err := parse(actual)
		return v, false, err
	}
}

func WithDefaultOrParse[T Number | string | bool](def T, parse ParseString[T]) Operation[T] {
	return func(actual string) (T, bool, error) {
		if actual == "" {
			return def, true, nil
		}

		v, err := parse(actual)
		return v, false, err
	}
}

func WithParse[T Number | string | bool](parse ParseString[T]) Operation[T] {
	return func(actual string) (T, bool, error) {
		v, err := parse(actual)
		return v, false, err
	}
}

type Constraint[T Number | string | bool] func(actual T) error

func WithMinimum[T Number](expected T) Constraint[T] {
	return func(actual T) error {
		if actual < expected {
			return errors.New(errMsgMinValueConstraint)
		}

		return nil
	}
}

func WithMaximum[T Number](expected T) Constraint[T] {
	return func(actual T) error {
		if actual > expected {
			return errors.New(errMsgMaxValueConstraint)
		}

		return nil
	}
}

// parseNumericParameter parses a numeric parameter to its respective type.
func parseNumericParameter[T Number](param string, fn Operation[T], checks ...Constraint[T]) (T, error) {
	v, ok, err := fn(param)
	if err != nil {
		return 0, err
	}

	if !ok {
		for _, check := range checks {
			if err := check(v); err != nil {
				return 0, err
			}
		}
	}

	return v, nil
}

// parseBoolParameter parses a string parameter to a bool
func parseBoolParameter(param string, fn Operation[bool]) (bool, error) {
	v, _, err := fn(param)
	return v, err
}

// parseQuery parses query paramaters and returns an error if any malformed value pairs are encountered.
func parseQuery(rawQuery string) (url.Values, error) {
	return url.ParseQuery(rawQuery)
}

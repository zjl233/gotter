// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /API/user/)
	Refresh(ctx echo.Context, params RefreshParams) error

	// (POST /API/user/)
	SignUp(ctx echo.Context) error

	// (GET /API/user/info)
	Info(ctx echo.Context, params InfoParams) error

	// (POST /API/user/login)
	Login(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Refresh converts echo context to params.
func (w *ServerInterfaceWrapper) Refresh(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RefreshParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "x-auth" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("x-auth")]; found {
		var XAuth XAuth
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for x-auth, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "x-auth", valueList[0], &XAuth)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter x-auth: %s", err))
		}

		params.XAuth = XAuth
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter x-auth is required, but not found"))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Refresh(ctx, params)
	return err
}

// SignUp converts echo context to params.
func (w *ServerInterfaceWrapper) SignUp(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SignUp(ctx)
	return err
}

// Info converts echo context to params.
func (w *ServerInterfaceWrapper) Info(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params InfoParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "x-auth" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("x-auth")]; found {
		var XAuth XAuth
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for x-auth, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "x-auth", valueList[0], &XAuth)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter x-auth: %s", err))
		}

		params.XAuth = XAuth
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter x-auth is required, but not found"))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Info(ctx, params)
	return err
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.Login(ctx)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/API/user/", wrapper.Refresh)
	router.POST("/API/user/", wrapper.SignUp)
	router.GET("/API/user/info", wrapper.Info)
	router.POST("/API/user/login", wrapper.Login)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+RXX2vkNhD/KkIt9EW33ttSCH7q9ghlIZeGbkMfjqUo8qytnC3pRuP4QvB3L5LtrHft",
	"LA2Bg6NvtqWZ+c1v/vqJK1s5a8CQ5+kTdxJlBQQY376+kzUV4SkDr1A70tbwlN83xOKJ4Dq8FyAzQC64",
	"kRXwdJATHOFLrREynhLWILhXBVQyKKRHF256Qm1y3rZtuOydNR6i6UtEi+FBWUNgKDxK50qtZACR3PuA",
	"5Gmk0aF1gKQ7eUAcmbF396CItyJ8/+jzqUvRHqvAe5kDF6f4Irq6pJHOO2tLkIZ30Ac/Pw0Xn03FB74T",
	"p1iC3DEGZavKGgYRipNEEjv1nZPRr2tobj3g1GGplK07no61hussBEawW6O/1CCYksZYYqqQJjpbya9X",
	"YPIQ6ve/CF5p8/w6w0QX5DkzP3lmtPocL7xWq5PeNxazFxwYjgWTxEqQnthFcACloph6e4uVJJ4eFB0h",
	"WC2PEFycQbCaQvhYe2JeVsCkZyMDbzV6kjo9cUMoxVjtAd2uFXw+B/7Rc/TFoDOdMbtnVABz1tMhw7Uh",
	"yAEDAaMUmpBz9zn/W5blpspnj/e2LG3TYdIElR/dGlnov0hE+XgQCzpeJzek4DSG1neN7BXKHNq9LmHe",
	"tZMIBYbHAeojdvBjRMWA5sjCEZO72AS02duh00kV6YdK6jJgdJpAVr/6RuY54ELbQ5Pddt/Y+mbD/gJZ",
	"ccFrDEIFkUuTZCQz6TRrRo0mAmSqtAaYdI4LXmoFxkdiextrJ1UBbLVYHmn3aZI0TbOQ8XhhMU96WZ9c",
	"bT5cXm8v360Wy0VBVRkZB6z8H/st4INWMAcxiVeSkJaaynDldxvxrW82XPAHQN8hf79YLpZBp3VgpNM8",
	"5T/HT6FCqIghT9Y3m6T2gEl4y2GmKSLsEXzB7psQxVBDcaxsMp7yP7uzqPEwDT898R8R9jzlPySHmZkc",
	"riT91Gt3J6NstVy+YZC9PHcEr/suMAesnxlJ7BRt+x/Gz4caEQyxoJVJk7Fjkrohf7IXnBnoUf9e9uDn",
	"ID6zlHTTPsrE7jSJl9e5YbVjBpqIbxK0rc7NretXDvD0m80eX8X6OQ6HoTtD2vYEmGBZ3ZmBjPV9goVh",
	"K2NXyCZLUTvJlvffRbZc9w4zhRCc/VYZ0opRfQ+t84UapxoNU6O0nqTNJij4Pxc6f3sQSpvr6Nl86V6F",
	"43n249EbavbFDfjsdjm3tJ2f+jML2W5uf0edBV7JsnLk2bla/z4Spguhr5UC779ZoYcfH8CHoSiPlpvS",
	"KlkW1lN6sbxY8nbX/hsAAP//qTpsRcoOAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

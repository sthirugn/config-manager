// Package controllers provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package controllers

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get archive of state changes for requesting account
	// (GET /states)
	GetStates(ctx echo.Context, params GetStatesParams) error
	// Update and roll out configuration state for requesting account
	// (POST /states)
	UpdateStates(ctx echo.Context) error
	// Get the current state for requesting account
	// (GET /states/current)
	GetCurrentState(ctx echo.Context) error
	// Get a preview of the playbook built from the provided state map
	// (POST /states/preview)
	GetPlaybookPreview(ctx echo.Context) error
	// Get single state change for requesting account
	// (GET /states/{id})
	GetStateById(ctx echo.Context, id StateIDParam) error
	// Get ansible playbook for current state configuration
	// (GET /states/{id}/playbook)
	GetPlaybookById(ctx echo.Context, id StateIDParam) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetStates converts echo context to params.
func (w *ServerInterfaceWrapper) GetStates(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetStatesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStates(ctx, params)
	return err
}

// UpdateStates converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateStates(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateStates(ctx)
	return err
}

// GetCurrentState converts echo context to params.
func (w *ServerInterfaceWrapper) GetCurrentState(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCurrentState(ctx)
	return err
}

// GetPlaybookPreview converts echo context to params.
func (w *ServerInterfaceWrapper) GetPlaybookPreview(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPlaybookPreview(ctx)
	return err
}

// GetStateById converts echo context to params.
func (w *ServerInterfaceWrapper) GetStateById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id StateIDParam

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStateById(ctx, id)
	return err
}

// GetPlaybookById converts echo context to params.
func (w *ServerInterfaceWrapper) GetPlaybookById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id StateIDParam

	err = runtime.BindStyledParameter("simple", false, "id", ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPlaybookById(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/states", wrapper.GetStates)
	router.POST(baseURL+"/states", wrapper.UpdateStates)
	router.GET(baseURL+"/states/current", wrapper.GetCurrentState)
	router.POST(baseURL+"/states/preview", wrapper.GetPlaybookPreview)
	router.GET(baseURL+"/states/:id", wrapper.GetStateById)
	router.GET(baseURL+"/states/:id/playbook", wrapper.GetPlaybookById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RYW4/bNhP9KwS/71FZedPmxW+5FIHRLbJt2qdgEYzFkc1UvGQ4dNdY6L8XJOVbpL2l",
	"aRZ9s8SZ4ZkzZzi0bmTjjHcWLQc5v5EeCAwyUn660EZz+qEwNKQ9a2flXP4C19pEI2w0SyThWkEYYsdB",
	"sBOEHMnKSupk+jkibWUlLRiUc9nlgJUMzRoNlMgtxI7l/MWskqYElvPns/SkbXk6ryRvffLXlnGFJPu+",
	"ku/aNuAEuoVVugHGIHiNIjAQa7sS3gWdLBLctJCRCcIOWG8wIU9vExsdMoqAnCw1o0mBgIUBbtYH11sy",
	"dAXVZIrHOc0mc3rPwLh4c5mqMM4spFUB1KwTYq3Qsm410g6KB14fkGglK0n4OWpCJedMEY9R/Z+wlXP5",
	"v/oggLqshnqAIfuEaXiZfF42jYs2c27g+gLtitdyfl7y2j/uMwtM2q5kX+0cc9wsM3IeiTXmsHAIexeq",
	"3e59lXJ7aAqV7GCJ3X32F9kopbvDeG/0XLEhVbf8hE2GtrCaNbCjFGNExMUOy2hlzw0olYUK3eURSy10",
	"AasviMtq1WAb/Og82tCAT6/RJoV9kGhh2WGSgdKh/LyaKI62Qa/Wpf8f6UpoUGlIcB/v3h/L88MBRjWZ",
	"1xebXU1Qnyl8WdrjW6isIQRG9RHudfpdGwwMxj9anPpYLnf5HHT13SV9zGsYE7un9fS0enk0H9AyaQzD",
	"cEAlWnJG7M/KL0/CahgUEzENXE/EPZ4742DulkmxR0joCUM6Tu1KnA4NbVm0LpmknTbp5bDp5FbDIBzv",
	"9eswbspyJfNgeVBNdoI+VAaIYJufHUM3lVdeELkwiabWRavuwH1r0RdvUvTWkUk9IGPMQ2V0DBzkf2yt",
	"gPEZa4Njlz4rv3Vj7K+dbfVKGLCwQhIBaaObHEFzhyMDWckNUii+s7PZ2XkuuEcLXsu5/CG/qvJozHTX",
	"uRnyz1XRRJJyPlQWSs7lWyxTKmSnw2Xow3SlDiZ1uSz11b2Gw8Wlv8py8c6Gguf5bFbayTKWhgLvu3SX",
	"0c7Wn0LK8eYxE3zfspnvU57f/ZyI+rFsebr0CpT4DT9HDFxszickFnmd+qWAE0aHkFojtzUdfF9MxV9Y",
	"RrLQifdIGyTxE5ErIgzRGKBtKcP+ouNaUW4+zRrsCsPQj3mTtCkcTmzvwkRN//BJivuyDq6vnNp+W76n",
	"eM4LwoAvl0iLqPJxtUQRM6zxLa3/F3VxchH7R7L46tKWagiwSpDrOuEiiyY3dSw1G6p9W5X7atfEdROJ",
	"Bkpua+bXxaQk/J8g9sn6Lf//KXQ9vASecKPxr3wrmGy+t8iXHWyXzv15Odg+WQvGgCo13wptgogCxIBf",
	"+AHjV3Qj4zXXG6vOtmC6U7wTY+/rZDE+HPfIh/+zuwTEMuqOizLye3IbrVANJTXgT+p3o1V/7zB8tV2o",
	"R8/Dk/+y323YfVOSU491eDJ9HtITidN6L6g7yN01xnfh92lkaoNedkfqTPSdHjInZ//wySEfVIWFSJ2c",
	"yzWzD/O6bjoX1RmhWgOfNc7U4HVdIjwbroT1Jl34TkF6cio2ZYNqF/OBvvtPRPnzSn/V/x0AAP//hMAA",
	"dzQTAAA=",
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


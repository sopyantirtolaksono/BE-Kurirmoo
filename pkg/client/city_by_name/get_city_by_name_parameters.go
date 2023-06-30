// Code generated by go-swagger; DO NOT EDIT.

package city_by_name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetCityByNameParams creates a new GetCityByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCityByNameParams() *GetCityByNameParams {
	return &GetCityByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCityByNameParamsWithTimeout creates a new GetCityByNameParams object
// with the ability to set a timeout on a request.
func NewGetCityByNameParamsWithTimeout(timeout time.Duration) *GetCityByNameParams {
	return &GetCityByNameParams{
		timeout: timeout,
	}
}

// NewGetCityByNameParamsWithContext creates a new GetCityByNameParams object
// with the ability to set a context for a request.
func NewGetCityByNameParamsWithContext(ctx context.Context) *GetCityByNameParams {
	return &GetCityByNameParams{
		Context: ctx,
	}
}

// NewGetCityByNameParamsWithHTTPClient creates a new GetCityByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCityByNameParamsWithHTTPClient(client *http.Client) *GetCityByNameParams {
	return &GetCityByNameParams{
		HTTPClient: client,
	}
}

/*
GetCityByNameParams contains all the parameters to send to the API endpoint

	for the get city by name operation.

	Typically these are written to a http.Request.
*/
type GetCityByNameParams struct {

	// CityName.
	CityName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get city by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCityByNameParams) WithDefaults() *GetCityByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get city by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCityByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get city by name params
func (o *GetCityByNameParams) WithTimeout(timeout time.Duration) *GetCityByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get city by name params
func (o *GetCityByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get city by name params
func (o *GetCityByNameParams) WithContext(ctx context.Context) *GetCityByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get city by name params
func (o *GetCityByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get city by name params
func (o *GetCityByNameParams) WithHTTPClient(client *http.Client) *GetCityByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get city by name params
func (o *GetCityByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCityName adds the cityName to the get city by name params
func (o *GetCityByNameParams) WithCityName(cityName string) *GetCityByNameParams {
	o.SetCityName(cityName)
	return o
}

// SetCityName adds the cityName to the get city by name params
func (o *GetCityByNameParams) SetCityName(cityName string) {
	o.CityName = cityName
}

// WriteToRequest writes these params to a swagger request
func (o *GetCityByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param city_name
	if err := r.SetPathParam("city_name", o.CityName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
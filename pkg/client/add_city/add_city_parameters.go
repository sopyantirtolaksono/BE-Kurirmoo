// Code generated by go-swagger; DO NOT EDIT.

package add_city

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

// NewAddCityParams creates a new AddCityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddCityParams() *AddCityParams {
	return &AddCityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddCityParamsWithTimeout creates a new AddCityParams object
// with the ability to set a timeout on a request.
func NewAddCityParamsWithTimeout(timeout time.Duration) *AddCityParams {
	return &AddCityParams{
		timeout: timeout,
	}
}

// NewAddCityParamsWithContext creates a new AddCityParams object
// with the ability to set a context for a request.
func NewAddCityParamsWithContext(ctx context.Context) *AddCityParams {
	return &AddCityParams{
		Context: ctx,
	}
}

// NewAddCityParamsWithHTTPClient creates a new AddCityParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddCityParamsWithHTTPClient(client *http.Client) *AddCityParams {
	return &AddCityParams{
		HTTPClient: client,
	}
}

/*
AddCityParams contains all the parameters to send to the API endpoint

	for the add city operation.

	Typically these are written to a http.Request.
*/
type AddCityParams struct {

	// Data.
	Data AddCityBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add city params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCityParams) WithDefaults() *AddCityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add city params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add city params
func (o *AddCityParams) WithTimeout(timeout time.Duration) *AddCityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add city params
func (o *AddCityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add city params
func (o *AddCityParams) WithContext(ctx context.Context) *AddCityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add city params
func (o *AddCityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add city params
func (o *AddCityParams) WithHTTPClient(client *http.Client) *AddCityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add city params
func (o *AddCityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the add city params
func (o *AddCityParams) WithData(data AddCityBody) *AddCityParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the add city params
func (o *AddCityParams) SetData(data AddCityBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *AddCityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
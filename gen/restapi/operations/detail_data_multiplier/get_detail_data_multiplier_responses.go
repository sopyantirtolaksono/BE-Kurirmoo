// Code generated by go-swagger; DO NOT EDIT.

package detail_data_multiplier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"kurirmoo/gen/models"
)

// GetDetailDataMultiplierOKCode is the HTTP code returned for type GetDetailDataMultiplierOK
const GetDetailDataMultiplierOKCode int = 200

/*
GetDetailDataMultiplierOK A JSON of detail data multiplier

swagger:response getDetailDataMultiplierOK
*/
type GetDetailDataMultiplierOK struct {

	/*
	  In: Body
	*/
	Payload *GetDetailDataMultiplierOKBody `json:"body,omitempty"`
}

// NewGetDetailDataMultiplierOK creates GetDetailDataMultiplierOK with default headers values
func NewGetDetailDataMultiplierOK() *GetDetailDataMultiplierOK {

	return &GetDetailDataMultiplierOK{}
}

// WithPayload adds the payload to the get detail data multiplier o k response
func (o *GetDetailDataMultiplierOK) WithPayload(payload *GetDetailDataMultiplierOKBody) *GetDetailDataMultiplierOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get detail data multiplier o k response
func (o *GetDetailDataMultiplierOK) SetPayload(payload *GetDetailDataMultiplierOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDetailDataMultiplierOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDetailDataMultiplierBadRequestCode is the HTTP code returned for type GetDetailDataMultiplierBadRequest
const GetDetailDataMultiplierBadRequestCode int = 400

/*
GetDetailDataMultiplierBadRequest Bad Request

swagger:response getDetailDataMultiplierBadRequest
*/
type GetDetailDataMultiplierBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDetailDataMultiplierBadRequest creates GetDetailDataMultiplierBadRequest with default headers values
func NewGetDetailDataMultiplierBadRequest() *GetDetailDataMultiplierBadRequest {

	return &GetDetailDataMultiplierBadRequest{}
}

// WithPayload adds the payload to the get detail data multiplier bad request response
func (o *GetDetailDataMultiplierBadRequest) WithPayload(payload *models.Error) *GetDetailDataMultiplierBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get detail data multiplier bad request response
func (o *GetDetailDataMultiplierBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDetailDataMultiplierBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDetailDataMultiplierNotFoundCode is the HTTP code returned for type GetDetailDataMultiplierNotFound
const GetDetailDataMultiplierNotFoundCode int = 404

/*
GetDetailDataMultiplierNotFound The specified resource was not found

swagger:response getDetailDataMultiplierNotFound
*/
type GetDetailDataMultiplierNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDetailDataMultiplierNotFound creates GetDetailDataMultiplierNotFound with default headers values
func NewGetDetailDataMultiplierNotFound() *GetDetailDataMultiplierNotFound {

	return &GetDetailDataMultiplierNotFound{}
}

// WithPayload adds the payload to the get detail data multiplier not found response
func (o *GetDetailDataMultiplierNotFound) WithPayload(payload *models.Error) *GetDetailDataMultiplierNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get detail data multiplier not found response
func (o *GetDetailDataMultiplierNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDetailDataMultiplierNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDetailDataMultiplierInternalServerErrorCode is the HTTP code returned for type GetDetailDataMultiplierInternalServerError
const GetDetailDataMultiplierInternalServerErrorCode int = 500

/*
GetDetailDataMultiplierInternalServerError Internal server error

swagger:response getDetailDataMultiplierInternalServerError
*/
type GetDetailDataMultiplierInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDetailDataMultiplierInternalServerError creates GetDetailDataMultiplierInternalServerError with default headers values
func NewGetDetailDataMultiplierInternalServerError() *GetDetailDataMultiplierInternalServerError {

	return &GetDetailDataMultiplierInternalServerError{}
}

// WithPayload adds the payload to the get detail data multiplier internal server error response
func (o *GetDetailDataMultiplierInternalServerError) WithPayload(payload *models.Error) *GetDetailDataMultiplierInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get detail data multiplier internal server error response
func (o *GetDetailDataMultiplierInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDetailDataMultiplierInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetDetailDataMultiplierDefault error

swagger:response getDetailDataMultiplierDefault
*/
type GetDetailDataMultiplierDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDetailDataMultiplierDefault creates GetDetailDataMultiplierDefault with default headers values
func NewGetDetailDataMultiplierDefault(code int) *GetDetailDataMultiplierDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDetailDataMultiplierDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get detail data multiplier default response
func (o *GetDetailDataMultiplierDefault) WithStatusCode(code int) *GetDetailDataMultiplierDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get detail data multiplier default response
func (o *GetDetailDataMultiplierDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get detail data multiplier default response
func (o *GetDetailDataMultiplierDefault) WithPayload(payload *models.Error) *GetDetailDataMultiplierDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get detail data multiplier default response
func (o *GetDetailDataMultiplierDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDetailDataMultiplierDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
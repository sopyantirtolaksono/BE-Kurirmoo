// Code generated by go-swagger; DO NOT EDIT.

package update_trip_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"kurirmoo/gen/models"
)

// UpdateTripStatusReader is a Reader for the UpdateTripStatus structure.
type UpdateTripStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTripStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTripStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTripStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTripStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTripStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTripStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateTripStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTripStatusOK creates a UpdateTripStatusOK with default headers values
func NewUpdateTripStatusOK() *UpdateTripStatusOK {
	return &UpdateTripStatusOK{}
}

/*
UpdateTripStatusOK describes a response with status code 200, with default header values.

OK
*/
type UpdateTripStatusOK struct {
	Payload *models.Success
}

// IsSuccess returns true when this update trip status o k response has a 2xx status code
func (o *UpdateTripStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update trip status o k response has a 3xx status code
func (o *UpdateTripStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trip status o k response has a 4xx status code
func (o *UpdateTripStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trip status o k response has a 5xx status code
func (o *UpdateTripStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update trip status o k response a status code equal to that given
func (o *UpdateTripStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update trip status o k response
func (o *UpdateTripStatusOK) Code() int {
	return 200
}

func (o *UpdateTripStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusOK  %+v", 200, o.Payload)
}

func (o *UpdateTripStatusOK) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusOK  %+v", 200, o.Payload)
}

func (o *UpdateTripStatusOK) GetPayload() *models.Success {
	return o.Payload
}

func (o *UpdateTripStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Success)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTripStatusBadRequest creates a UpdateTripStatusBadRequest with default headers values
func NewUpdateTripStatusBadRequest() *UpdateTripStatusBadRequest {
	return &UpdateTripStatusBadRequest{}
}

/*
UpdateTripStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateTripStatusBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this update trip status bad request response has a 2xx status code
func (o *UpdateTripStatusBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trip status bad request response has a 3xx status code
func (o *UpdateTripStatusBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trip status bad request response has a 4xx status code
func (o *UpdateTripStatusBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trip status bad request response has a 5xx status code
func (o *UpdateTripStatusBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update trip status bad request response a status code equal to that given
func (o *UpdateTripStatusBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update trip status bad request response
func (o *UpdateTripStatusBadRequest) Code() int {
	return 400
}

func (o *UpdateTripStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTripStatusBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTripStatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTripStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTripStatusUnauthorized creates a UpdateTripStatusUnauthorized with default headers values
func NewUpdateTripStatusUnauthorized() *UpdateTripStatusUnauthorized {
	return &UpdateTripStatusUnauthorized{}
}

/*
UpdateTripStatusUnauthorized describes a response with status code 401, with default header values.

The specified resource was not found
*/
type UpdateTripStatusUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this update trip status unauthorized response has a 2xx status code
func (o *UpdateTripStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trip status unauthorized response has a 3xx status code
func (o *UpdateTripStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trip status unauthorized response has a 4xx status code
func (o *UpdateTripStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trip status unauthorized response has a 5xx status code
func (o *UpdateTripStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update trip status unauthorized response a status code equal to that given
func (o *UpdateTripStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update trip status unauthorized response
func (o *UpdateTripStatusUnauthorized) Code() int {
	return 401
}

func (o *UpdateTripStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTripStatusUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTripStatusUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTripStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTripStatusNotFound creates a UpdateTripStatusNotFound with default headers values
func NewUpdateTripStatusNotFound() *UpdateTripStatusNotFound {
	return &UpdateTripStatusNotFound{}
}

/*
UpdateTripStatusNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type UpdateTripStatusNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this update trip status not found response has a 2xx status code
func (o *UpdateTripStatusNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trip status not found response has a 3xx status code
func (o *UpdateTripStatusNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trip status not found response has a 4xx status code
func (o *UpdateTripStatusNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update trip status not found response has a 5xx status code
func (o *UpdateTripStatusNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update trip status not found response a status code equal to that given
func (o *UpdateTripStatusNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update trip status not found response
func (o *UpdateTripStatusNotFound) Code() int {
	return 404
}

func (o *UpdateTripStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTripStatusNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTripStatusNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTripStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTripStatusInternalServerError creates a UpdateTripStatusInternalServerError with default headers values
func NewUpdateTripStatusInternalServerError() *UpdateTripStatusInternalServerError {
	return &UpdateTripStatusInternalServerError{}
}

/*
UpdateTripStatusInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateTripStatusInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this update trip status internal server error response has a 2xx status code
func (o *UpdateTripStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update trip status internal server error response has a 3xx status code
func (o *UpdateTripStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update trip status internal server error response has a 4xx status code
func (o *UpdateTripStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update trip status internal server error response has a 5xx status code
func (o *UpdateTripStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update trip status internal server error response a status code equal to that given
func (o *UpdateTripStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update trip status internal server error response
func (o *UpdateTripStatusInternalServerError) Code() int {
	return 500
}

func (o *UpdateTripStatusInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTripStatusInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTripStatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTripStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTripStatusDefault creates a UpdateTripStatusDefault with default headers values
func NewUpdateTripStatusDefault(code int) *UpdateTripStatusDefault {
	return &UpdateTripStatusDefault{
		_statusCode: code,
	}
}

/*
UpdateTripStatusDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateTripStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update trip status default response has a 2xx status code
func (o *UpdateTripStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update trip status default response has a 3xx status code
func (o *UpdateTripStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update trip status default response has a 4xx status code
func (o *UpdateTripStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update trip status default response has a 5xx status code
func (o *UpdateTripStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update trip status default response a status code equal to that given
func (o *UpdateTripStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update trip status default response
func (o *UpdateTripStatusDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTripStatusDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatus default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTripStatusDefault) String() string {
	return fmt.Sprintf("[PUT /api/v1/drivers/{id}][%d] updateTripStatus default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTripStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateTripStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

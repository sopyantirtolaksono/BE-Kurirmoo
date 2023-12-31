// Code generated by go-swagger; DO NOT EDIT.

package city_by_name

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"kurirmoo/gen/models"
)

// GetCityByNameReader is a Reader for the GetCityByName structure.
type GetCityByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCityByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCityByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCityByNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCityByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCityByNameOK creates a GetCityByNameOK with default headers values
func NewGetCityByNameOK() *GetCityByNameOK {
	return &GetCityByNameOK{}
}

/*
GetCityByNameOK describes a response with status code 200, with default header values.

A JSON array of city's name
*/
type GetCityByNameOK struct {
	Payload *GetCityByNameOKBody
}

// IsSuccess returns true when this get city by name o k response has a 2xx status code
func (o *GetCityByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get city by name o k response has a 3xx status code
func (o *GetCityByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get city by name o k response has a 4xx status code
func (o *GetCityByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get city by name o k response has a 5xx status code
func (o *GetCityByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get city by name o k response a status code equal to that given
func (o *GetCityByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get city by name o k response
func (o *GetCityByNameOK) Code() int {
	return 200
}

func (o *GetCityByNameOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByNameOK  %+v", 200, o.Payload)
}

func (o *GetCityByNameOK) String() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByNameOK  %+v", 200, o.Payload)
}

func (o *GetCityByNameOK) GetPayload() *GetCityByNameOKBody {
	return o.Payload
}

func (o *GetCityByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetCityByNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCityByNameBadRequest creates a GetCityByNameBadRequest with default headers values
func NewGetCityByNameBadRequest() *GetCityByNameBadRequest {
	return &GetCityByNameBadRequest{}
}

/*
GetCityByNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCityByNameBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get city by name bad request response has a 2xx status code
func (o *GetCityByNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get city by name bad request response has a 3xx status code
func (o *GetCityByNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get city by name bad request response has a 4xx status code
func (o *GetCityByNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get city by name bad request response has a 5xx status code
func (o *GetCityByNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get city by name bad request response a status code equal to that given
func (o *GetCityByNameBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get city by name bad request response
func (o *GetCityByNameBadRequest) Code() int {
	return 400
}

func (o *GetCityByNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetCityByNameBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetCityByNameBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCityByNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCityByNameDefault creates a GetCityByNameDefault with default headers values
func NewGetCityByNameDefault(code int) *GetCityByNameDefault {
	return &GetCityByNameDefault{
		_statusCode: code,
	}
}

/*
GetCityByNameDefault describes a response with status code -1, with default header values.

error
*/
type GetCityByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get city by name default response has a 2xx status code
func (o *GetCityByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get city by name default response has a 3xx status code
func (o *GetCityByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get city by name default response has a 4xx status code
func (o *GetCityByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get city by name default response has a 5xx status code
func (o *GetCityByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get city by name default response a status code equal to that given
func (o *GetCityByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get city by name default response
func (o *GetCityByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetCityByNameDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCityByNameDefault) String() string {
	return fmt.Sprintf("[GET /api/v1/cities/{city_name}][%d] getCityByName default  %+v", o._statusCode, o.Payload)
}

func (o *GetCityByNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCityByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetCityByNameOKBody get city by name o k body
swagger:model GetCityByNameOKBody
*/
type GetCityByNameOKBody struct {

	// acronim
	Acronim string `json:"acronim,omitempty"`

	// city code
	CityCode string `json:"city_code,omitempty"`

	// city name
	CityName string `json:"city_name,omitempty"`
}

// Validate validates this get city by name o k body
func (o *GetCityByNameOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get city by name o k body based on context it is used
func (o *GetCityByNameOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetCityByNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetCityByNameOKBody) UnmarshalBinary(b []byte) error {
	var res GetCityByNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

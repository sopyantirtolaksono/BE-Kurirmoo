// Code generated by go-swagger; DO NOT EDIT.

package otp_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"kurirmoo/gen/models"
)

// PostRequestWhatsappReader is a Reader for the PostRequestWhatsapp structure.
type PostRequestWhatsappReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRequestWhatsappReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRequestWhatsappOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRequestWhatsappBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRequestWhatsappNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /chats/send] PostRequestWhatsapp", response, response.Code())
	}
}

// NewPostRequestWhatsappOK creates a PostRequestWhatsappOK with default headers values
func NewPostRequestWhatsappOK() *PostRequestWhatsappOK {
	return &PostRequestWhatsappOK{}
}

/*
PostRequestWhatsappOK describes a response with status code 200, with default header values.

success request whatsapp otp
*/
type PostRequestWhatsappOK struct {
	Payload *models.OtpWhatsapp
}

// IsSuccess returns true when this post request whatsapp o k response has a 2xx status code
func (o *PostRequestWhatsappOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post request whatsapp o k response has a 3xx status code
func (o *PostRequestWhatsappOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post request whatsapp o k response has a 4xx status code
func (o *PostRequestWhatsappOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post request whatsapp o k response has a 5xx status code
func (o *PostRequestWhatsappOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post request whatsapp o k response a status code equal to that given
func (o *PostRequestWhatsappOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post request whatsapp o k response
func (o *PostRequestWhatsappOK) Code() int {
	return 200
}

func (o *PostRequestWhatsappOK) Error() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappOK  %+v", 200, o.Payload)
}

func (o *PostRequestWhatsappOK) String() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappOK  %+v", 200, o.Payload)
}

func (o *PostRequestWhatsappOK) GetPayload() *models.OtpWhatsapp {
	return o.Payload
}

func (o *PostRequestWhatsappOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OtpWhatsapp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRequestWhatsappBadRequest creates a PostRequestWhatsappBadRequest with default headers values
func NewPostRequestWhatsappBadRequest() *PostRequestWhatsappBadRequest {
	return &PostRequestWhatsappBadRequest{}
}

/*
PostRequestWhatsappBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostRequestWhatsappBadRequest struct {
	Payload *models.OtpWhatsapp
}

// IsSuccess returns true when this post request whatsapp bad request response has a 2xx status code
func (o *PostRequestWhatsappBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post request whatsapp bad request response has a 3xx status code
func (o *PostRequestWhatsappBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post request whatsapp bad request response has a 4xx status code
func (o *PostRequestWhatsappBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post request whatsapp bad request response has a 5xx status code
func (o *PostRequestWhatsappBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post request whatsapp bad request response a status code equal to that given
func (o *PostRequestWhatsappBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post request whatsapp bad request response
func (o *PostRequestWhatsappBadRequest) Code() int {
	return 400
}

func (o *PostRequestWhatsappBadRequest) Error() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappBadRequest  %+v", 400, o.Payload)
}

func (o *PostRequestWhatsappBadRequest) String() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappBadRequest  %+v", 400, o.Payload)
}

func (o *PostRequestWhatsappBadRequest) GetPayload() *models.OtpWhatsapp {
	return o.Payload
}

func (o *PostRequestWhatsappBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OtpWhatsapp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRequestWhatsappNotFound creates a PostRequestWhatsappNotFound with default headers values
func NewPostRequestWhatsappNotFound() *PostRequestWhatsappNotFound {
	return &PostRequestWhatsappNotFound{}
}

/*
PostRequestWhatsappNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PostRequestWhatsappNotFound struct {
	Payload *models.OtpWhatsapp
}

// IsSuccess returns true when this post request whatsapp not found response has a 2xx status code
func (o *PostRequestWhatsappNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post request whatsapp not found response has a 3xx status code
func (o *PostRequestWhatsappNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post request whatsapp not found response has a 4xx status code
func (o *PostRequestWhatsappNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post request whatsapp not found response has a 5xx status code
func (o *PostRequestWhatsappNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post request whatsapp not found response a status code equal to that given
func (o *PostRequestWhatsappNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post request whatsapp not found response
func (o *PostRequestWhatsappNotFound) Code() int {
	return 404
}

func (o *PostRequestWhatsappNotFound) Error() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappNotFound  %+v", 404, o.Payload)
}

func (o *PostRequestWhatsappNotFound) String() string {
	return fmt.Sprintf("[POST /chats/send][%d] postRequestWhatsappNotFound  %+v", 404, o.Payload)
}

func (o *PostRequestWhatsappNotFound) GetPayload() *models.OtpWhatsapp {
	return o.Payload
}

func (o *PostRequestWhatsappNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OtpWhatsapp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostRequestWhatsappBody post request whatsapp body
swagger:model PostRequestWhatsappBody
*/
type PostRequestWhatsappBody struct {

	// message
	// Required: true
	Message *PostRequestWhatsappParamsBodyMessage `json:"message"`

	// Phone number string value with prefix format '62'
	// Example: 6281234567890
	// Required: true
	Receiver *string `json:"receiver"`
}

// Validate validates this post request whatsapp body
func (o *PostRequestWhatsappBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReceiver(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostRequestWhatsappBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	if o.Message != nil {
		if err := o.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "message")
			}
			return err
		}
	}

	return nil
}

func (o *PostRequestWhatsappBody) validateReceiver(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"receiver", "body", o.Receiver); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post request whatsapp body based on the context it is used
func (o *PostRequestWhatsappBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostRequestWhatsappBody) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if o.Message != nil {

		if err := o.Message.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "message")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostRequestWhatsappBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRequestWhatsappBody) UnmarshalBinary(b []byte) error {
	var res PostRequestWhatsappBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PostRequestWhatsappParamsBodyMessage Object message with text
swagger:model PostRequestWhatsappParamsBodyMessage
*/
type PostRequestWhatsappParamsBodyMessage struct {

	// text chat
	// Example: this is message from whatsapp
	Text string `json:"text,omitempty"`
}

// Validate validates this post request whatsapp params body message
func (o *PostRequestWhatsappParamsBodyMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post request whatsapp params body message based on context it is used
func (o *PostRequestWhatsappParamsBodyMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRequestWhatsappParamsBodyMessage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRequestWhatsappParamsBodyMessage) UnmarshalBinary(b []byte) error {
	var res PostRequestWhatsappParamsBodyMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

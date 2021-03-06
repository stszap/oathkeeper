// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/oathkeeper/sdk/go/oathkeeper/models"
)

// GetWellKnownReader is a Reader for the GetWellKnown structure.
type GetWellKnownReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWellKnownReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWellKnownOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWellKnownUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetWellKnownForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWellKnownOK creates a GetWellKnownOK with default headers values
func NewGetWellKnownOK() *GetWellKnownOK {
	return &GetWellKnownOK{}
}

/*GetWellKnownOK handles this case with default header values.

jsonWebKeySet
*/
type GetWellKnownOK struct {
	Payload *models.SwaggerJSONWebKeySet
}

func (o *GetWellKnownOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/jwks.json][%d] getWellKnownOK  %+v", 200, o.Payload)
}

func (o *GetWellKnownOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggerJSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWellKnownUnauthorized creates a GetWellKnownUnauthorized with default headers values
func NewGetWellKnownUnauthorized() *GetWellKnownUnauthorized {
	return &GetWellKnownUnauthorized{}
}

/*GetWellKnownUnauthorized handles this case with default header values.

The standard error format
*/
type GetWellKnownUnauthorized struct {
	Payload *GetWellKnownUnauthorizedBody
}

func (o *GetWellKnownUnauthorized) Error() string {
	return fmt.Sprintf("[GET /.well-known/jwks.json][%d] getWellKnownUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWellKnownUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWellKnownUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWellKnownForbidden creates a GetWellKnownForbidden with default headers values
func NewGetWellKnownForbidden() *GetWellKnownForbidden {
	return &GetWellKnownForbidden{}
}

/*GetWellKnownForbidden handles this case with default header values.

The standard error format
*/
type GetWellKnownForbidden struct {
	Payload *GetWellKnownForbiddenBody
}

func (o *GetWellKnownForbidden) Error() string {
	return fmt.Sprintf("[GET /.well-known/jwks.json][%d] getWellKnownForbidden  %+v", 403, o.Payload)
}

func (o *GetWellKnownForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWellKnownForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWellKnownForbiddenBody get well known forbidden body
swagger:model GetWellKnownForbiddenBody
*/
type GetWellKnownForbiddenBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get well known forbidden body
func (o *GetWellKnownForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWellKnownForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellKnownForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWellKnownForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWellKnownUnauthorizedBody get well known unauthorized body
swagger:model GetWellKnownUnauthorizedBody
*/
type GetWellKnownUnauthorizedBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// details
	Details []map[string]interface{} `json:"details"`

	// message
	Message string `json:"message,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`

	// request
	Request string `json:"request,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this get well known unauthorized body
func (o *GetWellKnownUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWellKnownUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWellKnownUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res GetWellKnownUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

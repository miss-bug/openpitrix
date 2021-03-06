// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ChangePasswordReader is a Reader for the ChangePassword structure.
type ChangePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangePasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangePasswordOK creates a ChangePasswordOK with default headers values
func NewChangePasswordOK() *ChangePasswordOK {
	return &ChangePasswordOK{}
}

/*ChangePasswordOK handles this case with default header values.

A successful response.
*/
type ChangePasswordOK struct {
	Payload *models.OpenpitrixChangePasswordResponse
}

func (o *ChangePasswordOK) Error() string {
	return fmt.Sprintf("[POST /v1/users/password:change][%d] changePasswordOK  %+v", 200, o.Payload)
}

func (o *ChangePasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixChangePasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

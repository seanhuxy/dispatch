// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/image-manager/gen/models"
)

// AddImageReader is a Reader for the AddImage structure.
type AddImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddImageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddImageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddImageCreated creates a AddImageCreated with default headers values
func NewAddImageCreated() *AddImageCreated {
	return &AddImageCreated{}
}

/*AddImageCreated handles this case with default header values.

created
*/
type AddImageCreated struct {
	Payload *models.Image
}

func (o *AddImageCreated) Error() string {
	return fmt.Sprintf("[POST /][%d] addImageCreated  %+v", 201, o.Payload)
}

func (o *AddImageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImageBadRequest creates a AddImageBadRequest with default headers values
func NewAddImageBadRequest() *AddImageBadRequest {
	return &AddImageBadRequest{}
}

/*AddImageBadRequest handles this case with default header values.

Invalid input
*/
type AddImageBadRequest struct {
	Payload *models.Error
}

func (o *AddImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /][%d] addImageBadRequest  %+v", 400, o.Payload)
}

func (o *AddImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImageDefault creates a AddImageDefault with default headers values
func NewAddImageDefault(code int) *AddImageDefault {
	return &AddImageDefault{
		_statusCode: code,
	}
}

/*AddImageDefault handles this case with default header values.

Generic error response
*/
type AddImageDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add image default response
func (o *AddImageDefault) Code() int {
	return o._statusCode
}

func (o *AddImageDefault) Error() string {
	return fmt.Sprintf("[POST /][%d] addImage default  %+v", o._statusCode, o.Payload)
}

func (o *AddImageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

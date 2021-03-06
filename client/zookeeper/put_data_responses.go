// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutDataReader is a Reader for the PutData structure.
type PutDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPutDataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPutDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDataOK creates a PutDataOK with default headers values
func NewPutDataOK() *PutDataOK {
	return &PutDataOK{}
}

/*PutDataOK handles this case with default header values.

Success
*/
type PutDataOK struct {
}

func (o *PutDataOK) Error() string {
	return fmt.Sprintf("[PUT /zk/put][%d] putDataOK ", 200)
}

func (o *PutDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDataNoContent creates a PutDataNoContent with default headers values
func NewPutDataNoContent() *PutDataNoContent {
	return &PutDataNoContent{}
}

/*PutDataNoContent handles this case with default header values.

No Content
*/
type PutDataNoContent struct {
}

func (o *PutDataNoContent) Error() string {
	return fmt.Sprintf("[PUT /zk/put][%d] putDataNoContent ", 204)
}

func (o *PutDataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDataNotFound creates a PutDataNotFound with default headers values
func NewPutDataNotFound() *PutDataNotFound {
	return &PutDataNotFound{}
}

/*PutDataNotFound handles this case with default header values.

ZK Path not found
*/
type PutDataNotFound struct {
}

func (o *PutDataNotFound) Error() string {
	return fmt.Sprintf("[PUT /zk/put][%d] putDataNotFound ", 404)
}

func (o *PutDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDataInternalServerError creates a PutDataInternalServerError with default headers values
func NewPutDataInternalServerError() *PutDataInternalServerError {
	return &PutDataInternalServerError{}
}

/*PutDataInternalServerError handles this case with default header values.

Internal server error
*/
type PutDataInternalServerError struct {
}

func (o *PutDataInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /zk/put][%d] putDataInternalServerError ", 500)
}

func (o *PutDataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

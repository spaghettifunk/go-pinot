// Code generated by go-swagger; DO NOT EDIT.

package zookeeper

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LsReader is a Reader for the Ls structure.
type LsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewLsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLsOK creates a LsOK with default headers values
func NewLsOK() *LsOK {
	return &LsOK{}
}

/*LsOK handles this case with default header values.

Success
*/
type LsOK struct {
}

func (o *LsOK) Error() string {
	return fmt.Sprintf("[GET /zk/ls][%d] lsOK ", 200)
}

func (o *LsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLsNotFound creates a LsNotFound with default headers values
func NewLsNotFound() *LsNotFound {
	return &LsNotFound{}
}

/*LsNotFound handles this case with default header values.

ZK Path not found
*/
type LsNotFound struct {
}

func (o *LsNotFound) Error() string {
	return fmt.Sprintf("[GET /zk/ls][%d] lsNotFound ", 404)
}

func (o *LsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLsInternalServerError creates a LsInternalServerError with default headers values
func NewLsInternalServerError() *LsInternalServerError {
	return &LsInternalServerError{}
}

/*LsInternalServerError handles this case with default header values.

Internal server error
*/
type LsInternalServerError struct {
}

func (o *LsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /zk/ls][%d] lsInternalServerError ", 500)
}

func (o *LsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package table

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// AddTableReader is a Reader for the AddTable structure.
type AddTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddTableOK creates a AddTableOK with default headers values
func NewAddTableOK() *AddTableOK {
	return &AddTableOK{}
}

/*AddTableOK handles this case with default header values.

successful operation
*/
type AddTableOK struct {
	Payload *models.SuccessResponse
}

func (o *AddTableOK) Error() string {
	return fmt.Sprintf("[POST /tables][%d] addTableOK  %+v", 200, o.Payload)
}

func (o *AddTableOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *AddTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

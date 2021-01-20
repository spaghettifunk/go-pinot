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

// AssignInstancesReader is a Reader for the AssignInstances structure.
type AssignInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAssignInstancesOK creates a AssignInstancesOK with default headers values
func NewAssignInstancesOK() *AssignInstancesOK {
	return &AssignInstancesOK{}
}

/*AssignInstancesOK handles this case with default header values.

successful operation
*/
type AssignInstancesOK struct {
	Payload map[string]models.InstancePartitions
}

func (o *AssignInstancesOK) Error() string {
	return fmt.Sprintf("[POST /tables/{tableName}/assignInstances][%d] assignInstancesOK  %+v", 200, o.Payload)
}

func (o *AssignInstancesOK) GetPayload() map[string]models.InstancePartitions {
	return o.Payload
}

func (o *AssignInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
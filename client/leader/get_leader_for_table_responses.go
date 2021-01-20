// Code generated by go-swagger; DO NOT EDIT.

package leader

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// GetLeaderForTableReader is a Reader for the GetLeaderForTable structure.
type GetLeaderForTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLeaderForTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLeaderForTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLeaderForTableOK creates a GetLeaderForTableOK with default headers values
func NewGetLeaderForTableOK() *GetLeaderForTableOK {
	return &GetLeaderForTableOK{}
}

/*GetLeaderForTableOK handles this case with default header values.

successful operation
*/
type GetLeaderForTableOK struct {
	Payload *models.LeadControllerResponse
}

func (o *GetLeaderForTableOK) Error() string {
	return fmt.Sprintf("[GET /leader/tables/{tableName}][%d] getLeaderForTableOK  %+v", 200, o.Payload)
}

func (o *GetLeaderForTableOK) GetPayload() *models.LeadControllerResponse {
	return o.Payload
}

func (o *GetLeaderForTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LeadControllerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
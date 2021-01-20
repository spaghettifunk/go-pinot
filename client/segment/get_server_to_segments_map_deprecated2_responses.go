// Code generated by go-swagger; DO NOT EDIT.

package segment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetServerToSegmentsMapDeprecated2Reader is a Reader for the GetServerToSegmentsMapDeprecated2 structure.
type GetServerToSegmentsMapDeprecated2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServerToSegmentsMapDeprecated2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServerToSegmentsMapDeprecated2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetServerToSegmentsMapDeprecated2OK creates a GetServerToSegmentsMapDeprecated2OK with default headers values
func NewGetServerToSegmentsMapDeprecated2OK() *GetServerToSegmentsMapDeprecated2OK {
	return &GetServerToSegmentsMapDeprecated2OK{}
}

/*GetServerToSegmentsMapDeprecated2OK handles this case with default header values.

successful operation
*/
type GetServerToSegmentsMapDeprecated2OK struct {
	Payload []map[string]string
}

func (o *GetServerToSegmentsMapDeprecated2OK) Error() string {
	return fmt.Sprintf("[GET /tables/{tableName}/segments/metadata][%d] getServerToSegmentsMapDeprecated2OK  %+v", 200, o.Payload)
}

func (o *GetServerToSegmentsMapDeprecated2OK) GetPayload() []map[string]string {
	return o.Payload
}

func (o *GetServerToSegmentsMapDeprecated2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
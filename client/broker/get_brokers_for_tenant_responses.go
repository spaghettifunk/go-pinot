// Code generated by go-swagger; DO NOT EDIT.

package broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBrokersForTenantReader is a Reader for the GetBrokersForTenant structure.
type GetBrokersForTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrokersForTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBrokersForTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBrokersForTenantOK creates a GetBrokersForTenantOK with default headers values
func NewGetBrokersForTenantOK() *GetBrokersForTenantOK {
	return &GetBrokersForTenantOK{}
}

/*GetBrokersForTenantOK handles this case with default header values.

successful operation
*/
type GetBrokersForTenantOK struct {
	Payload []string
}

func (o *GetBrokersForTenantOK) Error() string {
	return fmt.Sprintf("[GET /brokers/tenants/{tenantName}][%d] getBrokersForTenantOK  %+v", 200, o.Payload)
}

func (o *GetBrokersForTenantOK) GetPayload() []string {
	return o.Payload
}

func (o *GetBrokersForTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
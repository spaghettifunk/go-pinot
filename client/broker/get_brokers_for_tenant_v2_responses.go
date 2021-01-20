// Code generated by go-swagger; DO NOT EDIT.

package broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// GetBrokersForTenantV2Reader is a Reader for the GetBrokersForTenantV2 structure.
type GetBrokersForTenantV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrokersForTenantV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBrokersForTenantV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBrokersForTenantV2OK creates a GetBrokersForTenantV2OK with default headers values
func NewGetBrokersForTenantV2OK() *GetBrokersForTenantV2OK {
	return &GetBrokersForTenantV2OK{}
}

/*GetBrokersForTenantV2OK handles this case with default header values.

successful operation
*/
type GetBrokersForTenantV2OK struct {
	Payload []*models.InstanceInfo
}

func (o *GetBrokersForTenantV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/brokers/tenants/{tenantName}][%d] getBrokersForTenantV2OK  %+v", 200, o.Payload)
}

func (o *GetBrokersForTenantV2OK) GetPayload() []*models.InstanceInfo {
	return o.Payload
}

func (o *GetBrokersForTenantV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
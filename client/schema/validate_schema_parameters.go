// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spaghettifunk/pinot-go-client/models"
)

// NewValidateSchemaParams creates a new ValidateSchemaParams object
// with the default values initialized.
func NewValidateSchemaParams() *ValidateSchemaParams {
	var ()
	return &ValidateSchemaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateSchemaParamsWithTimeout creates a new ValidateSchemaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateSchemaParamsWithTimeout(timeout time.Duration) *ValidateSchemaParams {
	var ()
	return &ValidateSchemaParams{

		timeout: timeout,
	}
}

// NewValidateSchemaParamsWithContext creates a new ValidateSchemaParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateSchemaParamsWithContext(ctx context.Context) *ValidateSchemaParams {
	var ()
	return &ValidateSchemaParams{

		Context: ctx,
	}
}

// NewValidateSchemaParamsWithHTTPClient creates a new ValidateSchemaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateSchemaParamsWithHTTPClient(client *http.Client) *ValidateSchemaParams {
	var ()
	return &ValidateSchemaParams{
		HTTPClient: client,
	}
}

/*ValidateSchemaParams contains all the parameters to send to the API endpoint
for the validate schema operation typically these are written to a http.Request
*/
type ValidateSchemaParams struct {

	/*Body*/
	Body *models.FormDataMultiPart

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate schema params
func (o *ValidateSchemaParams) WithTimeout(timeout time.Duration) *ValidateSchemaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate schema params
func (o *ValidateSchemaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate schema params
func (o *ValidateSchemaParams) WithContext(ctx context.Context) *ValidateSchemaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate schema params
func (o *ValidateSchemaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate schema params
func (o *ValidateSchemaParams) WithHTTPClient(client *http.Client) *ValidateSchemaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate schema params
func (o *ValidateSchemaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate schema params
func (o *ValidateSchemaParams) WithBody(body *models.FormDataMultiPart) *ValidateSchemaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate schema params
func (o *ValidateSchemaParams) SetBody(body *models.FormDataMultiPart) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateSchemaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

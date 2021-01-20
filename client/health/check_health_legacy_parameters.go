// Code generated by go-swagger; DO NOT EDIT.

package health

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
)

// NewCheckHealthLegacyParams creates a new CheckHealthLegacyParams object
// with the default values initialized.
func NewCheckHealthLegacyParams() *CheckHealthLegacyParams {

	return &CheckHealthLegacyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckHealthLegacyParamsWithTimeout creates a new CheckHealthLegacyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckHealthLegacyParamsWithTimeout(timeout time.Duration) *CheckHealthLegacyParams {

	return &CheckHealthLegacyParams{

		timeout: timeout,
	}
}

// NewCheckHealthLegacyParamsWithContext creates a new CheckHealthLegacyParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckHealthLegacyParamsWithContext(ctx context.Context) *CheckHealthLegacyParams {

	return &CheckHealthLegacyParams{

		Context: ctx,
	}
}

// NewCheckHealthLegacyParamsWithHTTPClient creates a new CheckHealthLegacyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckHealthLegacyParamsWithHTTPClient(client *http.Client) *CheckHealthLegacyParams {

	return &CheckHealthLegacyParams{
		HTTPClient: client,
	}
}

/*CheckHealthLegacyParams contains all the parameters to send to the API endpoint
for the check health legacy operation typically these are written to a http.Request
*/
type CheckHealthLegacyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check health legacy params
func (o *CheckHealthLegacyParams) WithTimeout(timeout time.Duration) *CheckHealthLegacyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check health legacy params
func (o *CheckHealthLegacyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check health legacy params
func (o *CheckHealthLegacyParams) WithContext(ctx context.Context) *CheckHealthLegacyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check health legacy params
func (o *CheckHealthLegacyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check health legacy params
func (o *CheckHealthLegacyParams) WithHTTPClient(client *http.Client) *CheckHealthLegacyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check health legacy params
func (o *CheckHealthLegacyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CheckHealthLegacyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
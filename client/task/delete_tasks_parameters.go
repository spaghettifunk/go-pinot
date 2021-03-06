// Code generated by go-swagger; DO NOT EDIT.

package task

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
	"github.com/go-openapi/swag"
)

// NewDeleteTasksParams creates a new DeleteTasksParams object
// with the default values initialized.
func NewDeleteTasksParams() *DeleteTasksParams {
	var (
		forceDeleteDefault = bool(false)
	)
	return &DeleteTasksParams{
		ForceDelete: &forceDeleteDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTasksParamsWithTimeout creates a new DeleteTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTasksParamsWithTimeout(timeout time.Duration) *DeleteTasksParams {
	var (
		forceDeleteDefault = bool(false)
	)
	return &DeleteTasksParams{
		ForceDelete: &forceDeleteDefault,

		timeout: timeout,
	}
}

// NewDeleteTasksParamsWithContext creates a new DeleteTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTasksParamsWithContext(ctx context.Context) *DeleteTasksParams {
	var (
		forceDeleteDefault = bool(false)
	)
	return &DeleteTasksParams{
		ForceDelete: &forceDeleteDefault,

		Context: ctx,
	}
}

// NewDeleteTasksParamsWithHTTPClient creates a new DeleteTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTasksParamsWithHTTPClient(client *http.Client) *DeleteTasksParams {
	var (
		forceDeleteDefault = bool(false)
	)
	return &DeleteTasksParams{
		ForceDelete: &forceDeleteDefault,
		HTTPClient:  client,
	}
}

/*DeleteTasksParams contains all the parameters to send to the API endpoint
for the delete tasks operation typically these are written to a http.Request
*/
type DeleteTasksParams struct {

	/*ForceDelete
	  Whether to force deleting the tasks (expert only option, enable with cautious

	*/
	ForceDelete *bool
	/*TaskType
	  Task type

	*/
	TaskType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete tasks params
func (o *DeleteTasksParams) WithTimeout(timeout time.Duration) *DeleteTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tasks params
func (o *DeleteTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tasks params
func (o *DeleteTasksParams) WithContext(ctx context.Context) *DeleteTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tasks params
func (o *DeleteTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tasks params
func (o *DeleteTasksParams) WithHTTPClient(client *http.Client) *DeleteTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tasks params
func (o *DeleteTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceDelete adds the forceDelete to the delete tasks params
func (o *DeleteTasksParams) WithForceDelete(forceDelete *bool) *DeleteTasksParams {
	o.SetForceDelete(forceDelete)
	return o
}

// SetForceDelete adds the forceDelete to the delete tasks params
func (o *DeleteTasksParams) SetForceDelete(forceDelete *bool) {
	o.ForceDelete = forceDelete
}

// WithTaskType adds the taskType to the delete tasks params
func (o *DeleteTasksParams) WithTaskType(taskType string) *DeleteTasksParams {
	o.SetTaskType(taskType)
	return o
}

// SetTaskType adds the taskType to the delete tasks params
func (o *DeleteTasksParams) SetTaskType(taskType string) {
	o.TaskType = taskType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceDelete != nil {

		// query param forceDelete
		var qrForceDelete bool
		if o.ForceDelete != nil {
			qrForceDelete = *o.ForceDelete
		}
		qForceDelete := swag.FormatBool(qrForceDelete)
		if qForceDelete != "" {
			if err := r.SetQueryParam("forceDelete", qForceDelete); err != nil {
				return err
			}
		}

	}

	// path param taskType
	if err := r.SetPathParam("taskType", o.TaskType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

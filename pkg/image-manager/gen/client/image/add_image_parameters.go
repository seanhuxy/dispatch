// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/image-manager/gen/models"
)

// NewAddImageParams creates a new AddImageParams object
// with the default values initialized.
func NewAddImageParams() *AddImageParams {
	var ()
	return &AddImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddImageParamsWithTimeout creates a new AddImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddImageParamsWithTimeout(timeout time.Duration) *AddImageParams {
	var ()
	return &AddImageParams{

		timeout: timeout,
	}
}

// NewAddImageParamsWithContext creates a new AddImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddImageParamsWithContext(ctx context.Context) *AddImageParams {
	var ()
	return &AddImageParams{

		Context: ctx,
	}
}

// NewAddImageParamsWithHTTPClient creates a new AddImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddImageParamsWithHTTPClient(client *http.Client) *AddImageParams {
	var ()
	return &AddImageParams{
		HTTPClient: client,
	}
}

/*AddImageParams contains all the parameters to send to the API endpoint
for the add image operation typically these are written to a http.Request
*/
type AddImageParams struct {

	/*Body
	  Image object

	*/
	Body *models.Image

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add image params
func (o *AddImageParams) WithTimeout(timeout time.Duration) *AddImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add image params
func (o *AddImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add image params
func (o *AddImageParams) WithContext(ctx context.Context) *AddImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add image params
func (o *AddImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add image params
func (o *AddImageParams) WithHTTPClient(client *http.Client) *AddImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add image params
func (o *AddImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add image params
func (o *AddImageParams) WithBody(body *models.Image) *AddImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add image params
func (o *AddImageParams) SetBody(body *models.Image) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Image)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

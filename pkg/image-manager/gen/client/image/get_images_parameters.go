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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetImagesParams creates a new GetImagesParams object
// with the default values initialized.
func NewGetImagesParams() *GetImagesParams {
	var ()
	return &GetImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesParamsWithTimeout creates a new GetImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImagesParamsWithTimeout(timeout time.Duration) *GetImagesParams {
	var ()
	return &GetImagesParams{

		timeout: timeout,
	}
}

// NewGetImagesParamsWithContext creates a new GetImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImagesParamsWithContext(ctx context.Context) *GetImagesParams {
	var ()
	return &GetImagesParams{

		Context: ctx,
	}
}

// NewGetImagesParamsWithHTTPClient creates a new GetImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImagesParamsWithHTTPClient(client *http.Client) *GetImagesParams {
	var ()
	return &GetImagesParams{
		HTTPClient: client,
	}
}

/*GetImagesParams contains all the parameters to send to the API endpoint
for the get images operation typically these are written to a http.Request
*/
type GetImagesParams struct {

	/*Language
	  image runtime language

	*/
	Language *string
	/*Tags
	  Filter on image tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get images params
func (o *GetImagesParams) WithTimeout(timeout time.Duration) *GetImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images params
func (o *GetImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images params
func (o *GetImagesParams) WithContext(ctx context.Context) *GetImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images params
func (o *GetImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) WithHTTPClient(client *http.Client) *GetImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguage adds the language to the get images params
func (o *GetImagesParams) WithLanguage(language *string) *GetImagesParams {
	o.SetLanguage(language)
	return o
}

// SetLanguage adds the language to the get images params
func (o *GetImagesParams) SetLanguage(language *string) {
	o.Language = language
}

// WithTags adds the tags to the get images params
func (o *GetImagesParams) WithTags(tags []string) *GetImagesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get images params
func (o *GetImagesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Language != nil {

		// query param language
		var qrLanguage string
		if o.Language != nil {
			qrLanguage = *o.Language
		}
		qLanguage := qrLanguage
		if qLanguage != "" {
			if err := r.SetQueryParam("language", qLanguage); err != nil {
				return err
			}
		}

	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

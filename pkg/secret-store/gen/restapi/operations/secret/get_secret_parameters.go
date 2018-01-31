// Code generated by go-swagger; DO NOT EDIT.

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSecretParams creates a new GetSecretParams object
// with the default values initialized.
func NewGetSecretParams() GetSecretParams {
	var ()
	return GetSecretParams{}
}

// GetSecretParams contains all the bound params for the get secret operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSecret
type GetSecretParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*name of the secret to operate on
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	SecretName string
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetSecretParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rSecretName, rhkSecretName, _ := route.Params.GetOK("secretName")
	if err := o.bindSecretName(rSecretName, rhkSecretName, route.Formats); err != nil {
		res = append(res, err)
	}

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecretParams) bindSecretName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.SecretName = raw

	if err := o.validateSecretName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetSecretParams) validateSecretName(formats strfmt.Registry) error {

	if err := validate.Pattern("secretName", "path", o.SecretName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *GetSecretParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	tagsIC := rawData

	if len(tagsIC) == 0 {
		return nil
	}

	var tagsIR []string
	for _, tagsIV := range tagsIC {
		tagsI := tagsIV

		tagsIR = append(tagsIR, tagsI)
	}

	o.Tags = tagsIR

	return nil
}

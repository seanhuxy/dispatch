// Code generated by go-swagger; DO NOT EDIT.

package base_image

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

// NewGetBaseImageByNameParams creates a new GetBaseImageByNameParams object
// with the default values initialized.
func NewGetBaseImageByNameParams() GetBaseImageByNameParams {
	var ()
	return GetBaseImageByNameParams{}
}

// GetBaseImageByNameParams contains all the bound params for the get base image by name operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBaseImageByName
type GetBaseImageByNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Name of base image to return
	  Required: true
	  Pattern: ^[\w\d\-]+$
	  In: path
	*/
	BaseImageName string
	/*Filter based on tags
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetBaseImageByNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rBaseImageName, rhkBaseImageName, _ := route.Params.GetOK("baseImageName")
	if err := o.bindBaseImageName(rBaseImageName, rhkBaseImageName, route.Formats); err != nil {
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

func (o *GetBaseImageByNameParams) bindBaseImageName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.BaseImageName = raw

	if err := o.validateBaseImageName(formats); err != nil {
		return err
	}

	return nil
}

func (o *GetBaseImageByNameParams) validateBaseImageName(formats strfmt.Registry) error {

	if err := validate.Pattern("baseImageName", "path", o.BaseImageName, `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (o *GetBaseImageByNameParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

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

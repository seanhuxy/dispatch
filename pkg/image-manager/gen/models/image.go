///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Image image
// swagger:model Image

type Image struct {

	// base image name
	// Required: true
	// Pattern: ^[\w\d\-]+$
	BaseImageName *string `json:"baseImageName"`

	// created time
	CreatedTime int64 `json:"createdTime,omitempty"`

	// docker Url
	DockerURL string `json:"dockerUrl,omitempty"`

	// groups
	Groups []string `json:"groups"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// language
	Language Language `json:"language,omitempty"`

	// name
	// Required: true
	// Pattern: ^[\w\d\-]+$
	Name *string `json:"name"`

	// reason
	Reason []string `json:"reason"`

	// spec
	Spec Spec `json:"spec,omitempty"`

	// status
	Status Status `json:"status,omitempty"`

	// tags
	Tags ImageTags `json:"tags"`
}

/* polymorph Image baseImageName false */

/* polymorph Image createdTime false */

/* polymorph Image dockerUrl false */

/* polymorph Image groups false */

/* polymorph Image id false */

/* polymorph Image language false */

/* polymorph Image name false */

/* polymorph Image reason false */

/* polymorph Image spec false */

/* polymorph Image status false */

/* polymorph Image tags false */

// Validate validates this image
func (m *Image) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseImageName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Image) validateBaseImageName(formats strfmt.Registry) error {

	if err := validate.Required("baseImageName", "body", m.BaseImageName); err != nil {
		return err
	}

	if err := validate.Pattern("baseImageName", "body", string(*m.BaseImageName), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	return nil
}

func (m *Image) validateLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.Language) { // not required
		return nil
	}

	if err := m.Language.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("language")
		}
		return err
	}

	return nil
}

func (m *Image) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[\w\d\-]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Image) validateReason(formats strfmt.Registry) error {

	if swag.IsZero(m.Reason) { // not required
		return nil
	}

	return nil
}

func (m *Image) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if err := m.Spec.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("spec")
		}
		return err
	}

	return nil
}

func (m *Image) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Image) UnmarshalBinary(b []byte) error {
	var res Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncResult SyncResult holds the result of a sync with LDAP. This gives us information on which users were updated and how.
//
// swagger:model SyncResult
type SyncResult struct {

	// elapsed
	Elapsed Duration `json:"Elapsed,omitempty"`

	// failed users
	FailedUsers []*FailedUser `json:"FailedUsers"`

	// missing user ids
	MissingUserIds []int64 `json:"MissingUserIds"`

	// started
	// Format: date-time
	Started strfmt.DateTime `json:"Started,omitempty"`

	// updated user ids
	UpdatedUserIds []int64 `json:"UpdatedUserIds"`
}

// Validate validates this sync result
func (m *SyncResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElapsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncResult) validateElapsed(formats strfmt.Registry) error {
	if swag.IsZero(m.Elapsed) { // not required
		return nil
	}

	if err := m.Elapsed.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Elapsed")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Elapsed")
		}
		return err
	}

	return nil
}

func (m *SyncResult) validateFailedUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.FailedUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedUsers); i++ {
		if swag.IsZero(m.FailedUsers[i]) { // not required
			continue
		}

		if m.FailedUsers[i] != nil {
			if err := m.FailedUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FailedUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FailedUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SyncResult) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("Started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sync result based on the context it is used
func (m *SyncResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateElapsed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailedUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncResult) contextValidateElapsed(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Elapsed.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Elapsed")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("Elapsed")
		}
		return err
	}

	return nil
}

func (m *SyncResult) contextValidateFailedUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FailedUsers); i++ {

		if m.FailedUsers[i] != nil {
			if err := m.FailedUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FailedUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FailedUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncResult) UnmarshalBinary(b []byte) error {
	var res SyncResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

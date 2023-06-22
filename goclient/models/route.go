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
)

// Route A Route is a node that contains definitions of how to handle alerts. This is modified
// from the upstream alertmanager in that it adds the ObjectMatchers property.
//
// swagger:model Route
type Route struct {

	// continue
	Continue bool `json:"continue,omitempty"`

	// group by
	GroupBy []string `json:"group_by"`

	// group interval
	GroupInterval string `json:"group_interval,omitempty"`

	// group wait
	GroupWait string `json:"group_wait,omitempty"`

	// Deprecated. Remove before v1.0 release.
	Match map[string]string `json:"match,omitempty"`

	// match re
	MatchRe MatchRegexps `json:"match_re,omitempty"`

	// matchers
	Matchers Matchers `json:"matchers,omitempty"`

	// mute time intervals
	MuteTimeIntervals []string `json:"mute_time_intervals"`

	// object matchers
	ObjectMatchers ObjectMatchers `json:"object_matchers,omitempty"`

	// provenance
	Provenance Provenance `json:"provenance,omitempty"`

	// receiver
	Receiver string `json:"receiver,omitempty"`

	// repeat interval
	RepeatInterval string `json:"repeat_interval,omitempty"`

	// routes
	Routes []*Route `json:"routes"`
}

// Validate validates this route
func (m *Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchRe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatchers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectMatchers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvenance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Route) validateMatchRe(formats strfmt.Registry) error {
	if swag.IsZero(m.MatchRe) { // not required
		return nil
	}

	if m.MatchRe != nil {
		if err := m.MatchRe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match_re")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("match_re")
			}
			return err
		}
	}

	return nil
}

func (m *Route) validateMatchers(formats strfmt.Registry) error {
	if swag.IsZero(m.Matchers) { // not required
		return nil
	}

	if err := m.Matchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("matchers")
		}
		return err
	}

	return nil
}

func (m *Route) validateObjectMatchers(formats strfmt.Registry) error {
	if swag.IsZero(m.ObjectMatchers) { // not required
		return nil
	}

	if err := m.ObjectMatchers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("object_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("object_matchers")
		}
		return err
	}

	return nil
}

func (m *Route) validateProvenance(formats strfmt.Registry) error {
	if swag.IsZero(m.Provenance) { // not required
		return nil
	}

	if err := m.Provenance.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

func (m *Route) validateRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.Routes) { // not required
		return nil
	}

	for i := 0; i < len(m.Routes); i++ {
		if swag.IsZero(m.Routes[i]) { // not required
			continue
		}

		if m.Routes[i] != nil {
			if err := m.Routes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this route based on the context it is used
func (m *Route) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatchRe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMatchers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjectMatchers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProvenance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Route) contextValidateMatchRe(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MatchRe.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("match_re")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("match_re")
		}
		return err
	}

	return nil
}

func (m *Route) contextValidateMatchers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Matchers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("matchers")
		}
		return err
	}

	return nil
}

func (m *Route) contextValidateObjectMatchers(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ObjectMatchers.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("object_matchers")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("object_matchers")
		}
		return err
	}

	return nil
}

func (m *Route) contextValidateProvenance(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Provenance.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("provenance")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("provenance")
		}
		return err
	}

	return nil
}

func (m *Route) contextValidateRoutes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Routes); i++ {

		if m.Routes[i] != nil {
			if err := m.Routes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Route) UnmarshalBinary(b []byte) error {
	var res Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

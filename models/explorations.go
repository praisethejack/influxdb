package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Explorations explorations

swagger:model Explorations
*/
type Explorations struct {

	/* explorations
	 */
	Explorations []*Exploration `json:"explorations,omitempty"`
}

// Validate validates this explorations
func (m *Explorations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExplorations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Explorations) validateExplorations(formats strfmt.Registry) error {

	if swag.IsZero(m.Explorations) { // not required
		return nil
	}

	for i := 0; i < len(m.Explorations); i++ {

		if swag.IsZero(m.Explorations[i]) { // not required
			continue
		}

		if m.Explorations[i] != nil {

			if err := m.Explorations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

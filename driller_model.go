package client

import (
	"time"
)

// DrillerModel Driller response payload
type DrillerModel struct {
	*DefaultModelBase
	LicenseNumber          *string            `json:"licenseNumber,omitempty"`
	LicenseExpirationDate  *time.Time         `json:"licenseExpirationDate,omitempty"`
	LicenseIssuerTerritory *string            `json:"licenseIssuerTerritory,omitempty"`
	CompanyName            *string            `json:"companyName,omitempty"`
	FirstName              *string            `json:"firstName,omitempty"`
	LastName               *string            `json:"lastName,omitempty"`
	Email                  *string            `json:"email,omitempty"`
	StreetAddress1         *string            `json:"streetAddress1,omitempty"`
	StreetAddress2         *string            `json:"streetAddress2,omitempty"`
	City                   *string            `json:"city,omitempty"`
	State                  *string            `json:"state,omitempty"`
	PostalCode             *string            `json:"postalCode,omitempty"`
	PhoneNumbers           []PhoneNumberModel `json:"phoneNumbers"`

	_Update func(model *DrillerModel) (*DrillerModel, error)
	_Delete func() error
}

// Init Initializes spec and default backing functions for model instance
func (model *DrillerModel) Init(spec *ServiceSpec) *DrillerModel {
	model.Spec = spec

	model._Update = func(model *DrillerModel) (*DrillerModel, error) {
		return nil, nil
	}
	model._Delete = func() error {
		return nil
	}
	return model
}

// GetID getter for id attribute
func (model *DrillerModel) GetID() uint {
	return model.ID
}

// Update old model with new
func (model *DrillerModel) Update(updatedModel *DrillerModel) (*DrillerModel, error) {
	return model._Update(updatedModel)
}

// Delete model
func (model *DrillerModel) Delete() error {
	return model._Delete()
}

// DrillerPhoneNumberModel phone number model for driller association
type DrillerPhoneNumberModel struct {
	*PhoneNumberModel
}

package hydros

import (
	"errors"
	"gopkg.in/guregu/null.v3"
)

// DrillerModel Driller response payload
type DrillerModel struct {
	*DefaultModelBase
	LicenseNumber          null.String         `json:"licenseNumber"`
	LicenseExpirationDate  null.Time           `json:"licenseExpirationDate"`
	LicenseIssuerTerritory null.String         `json:"licenseIssuerTerritory"`
	CompanyName            null.String         `json:"companyName"`
	FirstName              null.String         `json:"firstName"`
	LastName               null.String         `json:"lastName"`
	Email                  null.String         `json:"email"`
	StreetAddress1         null.String         `json:"streetAddress1"`
	StreetAddress2         null.String         `json:"streetAddress2"`
	City                   null.String         `json:"city"`
	State                  null.String         `json:"state"`
	PostalCode             null.String         `json:"postalCode"`
	PhoneNumbers           []*PhoneNumberModel `json:"phoneNumbers"`

	_Save   func(model *DrillerModel) (*DrillerModel, error)
	_Delete func(model *DrillerModel) error
}

// Init Initializes spec and default backing functions for model instance
func (model *DrillerModel) Init(spec *ServiceSpec) *DrillerModel {
	model.Spec = spec

	if serviceMock, ok := spec.modelServiceCallMocks["Save"]; ok {
		model._Save = serviceMock.MockFunc.(func(model *DrillerModel) (*DrillerModel, error))
	} else {
		model._Save = func(model *DrillerModel) (*DrillerModel, error) {
			return nil, errors.New("not implemented")
		}
	}

	if serviceMock, ok := spec.modelServiceCallMocks["Delete"]; ok {
		model._Delete = serviceMock.MockFunc.(func(model *DrillerModel) error)
	} else {
		model._Delete = func(model *DrillerModel) error {
			return errors.New("not implemented")
		}
	}
	return model
}

// GetID getter for id attribute
func (model *DrillerModel) GetID() uint {
	return model.ID
}

// Save changed model
func (model *DrillerModel) Save() (*DrillerModel, error) {
	return model._Save(model)
}

// Delete model
func (model *DrillerModel) Delete() error {
	return model._Delete(model)
}

// DrillerPhoneNumberModel phone number model for driller association
type DrillerPhoneNumberModel struct {
	*PhoneNumberModel
}

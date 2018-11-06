package hydros

// ContactModel Contact response payload
type ContactModel struct {
	*DefaultModelBase
	CompanyID      uint                      `json:"companyId"`
	FirstName      string                    `json:"firstName"`
	LastName       string                    `json:"lastName"`
	CompanyName    string                    `json:"companyName"`
	Email          string                    `json:"email"`
	Address1       string                    `json:"address1"`
	Address2       string                    `json:"address2"`
	City           string                    `json:"city"`
	State          string                    `json:"state"`
	PostalCode     string                    `json:"postalCode"`
	Classification string                    `json:"classification"`
	PhoneNumbers   []ContactPhoneNumberModel `json:"phoneNumbers"`

	_Update func(model *ContactModel) (*ContactModel, error)
	_Delete func() error
}

// Init Initializes spec and default backing functions for model instance
func (model *ContactModel) Init(spec *ServiceSpec) *ContactModel {
	model.Spec = spec

	model._Update = func(model *ContactModel) (*ContactModel, error) {
		return nil, nil
	}
	model._Delete = func() error {
		return nil
	}
	return model
}

// GetID getter for id attribute
func (model *ContactModel) GetID() uint {
	return model.ID
}

// Update old model with new
func (model *ContactModel) Update(updatedModel *ContactModel) (*ContactModel, error) {
	return model._Update(updatedModel)
}

// Delete model
func (model *ContactModel) Delete() error {
	return model._Delete()
}

// ContactPhoneNumberModel phone number model for contact association
type ContactPhoneNumberModel struct {
	*PhoneNumberModel
}

package hydros

// WellModel Well response payload
type WellModel struct {
	*DefaultModelBase
	Serial string `json:"serial"`

	_Update func(model *WellModel) (*WellModel, error)
	_Delete func() error
}

// Init Initializes spec and default backing functions for model instance
func (model *WellModel) Init(spec *ServiceSpec) *WellModel {
	model.Spec = spec

	model._Update = func(model *WellModel) (*WellModel, error) {
		return nil, nil
	}
	model._Delete = func() error {
		return nil
	}
	return model
}

// GetID getter for id attribute
func (model *WellModel) GetID() uint {
	return model.ID
}

// Update old model with new
func (model *WellModel) Update(updatedModel *WellModel) (*WellModel, error) {
	return model._Update(updatedModel)
}

// Delete model
func (model *WellModel) Delete() error {
	return model._Delete()
}

// WellPhoneNumberModel phone number model for well association
type WellPhoneNumberModel struct {
	*PhoneNumberModel
}

package hydros

// HistoryModel History response payload
type HistoryModel struct {
	*DefaultModelBase
	UpdateID  string `json:"updateId"`
	CompanyID string `json:"companyId"`
	Type      string `json:"type"`
	Operation string `json:"operation"`
	Patch     string `json:"patch"`
	Snapshot  string `json:"snapshot"`
}

// Init Initialized spec and default backing functions for model instance
func (model *HistoryModel) Init(spec *ServiceSpec) *HistoryModel {
	model.Spec = spec

	return model
}

// GetID getter for id attribute
func (model *HistoryModel) GetID() uint {
	return model.ID
}

// GetUpdateID getter for Update ID attributes
func (model *HistoryModel) GetUpdateID() string {
	return model.UpdateID
}

package hydros

import "time"

// MeterModel Meter response payload
type MeterModel struct {
	*DefaultModelBase
	Name            string      `json:"name"`
	Make            string      `json:"make"`
	Model           string      `json:"model"`
	SerialNumber    string      `json:"serialNumber"`
	StartReading    int         `json:"startReading"`
	Unit            string      `json:"unit"`
	Active          bool        `json:"active"`
	DateInService   time.Time   `json:"dateInService"`
	DecomissionDate *time.Time   `json:"decomissionDate,omitempty"`
	Wells           []WellModel `json:"wells"`
}

// Init Initalized spec and default backing functions for model instance
func (model *MeterModel) Init(spec *ServiceSpec) *MeterModel {
	model.Spec = spec
	return model
}

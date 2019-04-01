package hydros

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/guregu/null.v3"
	"time"
)

// MeterReadingModel Meter Reading response payload
type MeterReadingModel struct {
	*DefaultModelBase
	MeterID     uint           `json:"meterId"`
	Reading     float64        `json:"reading"`
	ReadingDate *time.Time     `json:"readingDate"`
	Notes       null.String    `json:"notes"`
	ExtraData   postgres.Jsonb `json:"extraData"`
}

// ProductionModel Production response payload
type ProductionModel struct {
	MeterID   uint       `json:"meterId"`
	Volume    float64    `json:"volume"`
	FromDate  *time.Time `json:"fromDate"`
	ToDate    *time.Time `json:"toDate"`
	Estimated bool       `json:"estimated"`
}

// Init Initialized spec and default backing functions for model instance
func (model *MeterReadingModel) Init(spec *ServiceSpec) *MeterReadingModel {
	model.Spec = spec
	return model
}

// GetID getter for id attribute
func (model *MeterReadingModel) GetID() uint {
	return model.ID
}

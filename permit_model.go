package hydros

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

// PermitModel Permit response payload
type PermitModel struct {
	*DefaultModelBase
	CompanyID              string              `json:"companyId"`
	WellID                 uint                `json:"wellId"`
	HistoryUpdateID        string              `json:"historyUpdateId"`
	PermitTemplateID       uint                `json:"permitTemplateId"`
	PermitTemplate         PermitTemplateModel `json:"permitTemplate"`
	IssuedDate             *time.Time          `json:"issuedDate"`
	ExpirationDate         *time.Time          `json:"expirationDate"`
	TerminationDate        *time.Time          `json:"terminationDate"`
	Aggregate              bool                `json:"aggregate"`
	AggregatePermitID      null.Int            `json:"aggregatePermitId"`
	AggregatedPermits      []PermitModel       `json:"aggregatedPermits"`
	AmendedBy              null.Int            `json:"amendedBy"`
	OperatorFirstName      null.String         `json:"operatorFirstName"`
	OperatorLastName       null.String         `json:"operatorLastName"`
	OperatorCompanyName    null.String         `json:"operatorCompanyName"`
	OperatorEmail          null.String         `json:"operatorEmail"`
	OperatorPhoneNumber1   null.String         `json:"operatorPhoneNumber1"`
	OperatorPhoneNumber2   null.String         `json:"operatorPhoneNumber2"`
	OperatorStreetAddress1 null.String         `json:"operatorStreetAddress1"`
	OperatorStreetAddress2 null.String         `json:"operatorStreetAddress2"`
	OperatorCity           null.String         `json:"operatorCity"`
	OperatorState          null.String         `json:"operatorState"`
	OperatorPostalCode     null.String         `json:"operatorPostalCode"`
}

// PermitTemplateModel PermitTemplate response payload
type PermitTemplateModel struct {
	ID              uint   `json:"id"`
	CompanyID       uint   `json:"companyId"`
	PermitName      string `json:"permitName"`
	Condition       string `json:"condition"`
	RequiredFields  string `json:"requiredFields"`
	DurationDays    int    `json:"durationDays"`
	CanAggregate    bool   `json:"canAggregate"`
	AggregateFields string `json:"aggregateFields"`
}

// Init Initialized spec and default backing functions for model instance
func (model *PermitModel) Init(spec *ServiceSpec) *PermitModel {
	model.Spec = spec

	return model
}

// GetID getter for id attribute
func (model *PermitModel) GetID() uint {
	return model.ID
}

type AmendWellPermitsRequest struct {
	HistoryUpdateID string `json:"historyUpdateId"`
}

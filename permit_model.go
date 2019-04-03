package hydros

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// PermitModel Permit response payload
type PermitModel struct {
	*DefaultModelBase
	CompanyID              uint                `json:"companyId"`
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

	_Metrics func(model *PermitModel, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*PermitMetricsModel, error)
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

	if serviceMock, ok := spec.ModelServiceCallMocks["Metrics"]; ok {
		model._Metrics = serviceMock.MockFunc.(func(model *PermitModel, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*PermitMetricsModel, error))
	} else {
		model._Metrics = func(model *PermitModel, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*PermitMetricsModel, error) {

			uri := fmt.Sprintf("%s/%s/%d/metrics.json",
				model.Spec.Client.URL.String(), model.Spec.ServiceName, model.ID)

			baseURL, _ := url.Parse(uri)
			params := url.Values{}

			if fromDate != nil {
				params.Add("fromDate", fromDate.Format("2006-01-02T15:04:05-0700"))
			}
			if toDate != nil {
				params.Add("fromDate", fromDate.Format("2006-01-02T15:04:05-0700"))
			}
			params.Add("estimateBounds", fmt.Sprint(estimateBounds))

			baseURL.RawQuery = params.Encode()

			req, err := http.NewRequest("GET", baseURL.String(), nil)

			headers := model.Spec.Client.CreateHeadersFunc()
			for h := 0; h < len(headers); h++ {
				req.Header.Add(headers[h].Key, headers[h].Value)
			}

			resp, err := model.Spec.Client.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			if resp.StatusCode != 200 {
				var errorResponse ErrorResponse
				err = json.Unmarshal(bodyBytes, &errorResponse)
				if err == nil && errorResponse.Message != "" {
					return nil, fmt.Errorf("%s: %s", errorResponse.Message, errorResponse.Description)
				}
				return nil, fmt.Errorf("%d error: %s", resp.StatusCode, string(bodyBytes))
			}

			var metrics PermitMetricsModel
			err = json.Unmarshal(bodyBytes, &metrics)
			if err != nil {
				return nil, err
			}
			return &metrics, nil
		}
	}

	return model
}

// GetID getter for id attribute
func (model *PermitModel) GetID() uint {
	return model.ID
}

// Metrics get permit metrics
func (model *PermitModel) Metrics(fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*PermitMetricsModel, error) {
	return model._Metrics(model, fromDate, toDate, estimateBounds)
}

type AmendWellPermitsRequest struct {
	HistoryUpdateID string `json:"historyUpdateId"`
	Patch           string `json:"patch"`
}

// PermitTemplateModel PermitMetricsModel response payload
type PermitMetricsModel struct {
	PermitsCount                        int        `json:"permitsCount"`
	WellsCount                          int        `json:"wellsCount"`
	MetersCount                         int        `json:"metersCount"`
	OverPermittedProduction             bool       `json:"overPermittedProduction"`
	TotalEstimatedAnnualWaterProduction float32    `json:"totalEstimatedAnnualWaterProduction"`
	TotalVolumeProduced                 float32    `json:"totalVolumeProduced"`
	FromDate                            *time.Time `json:"fromDate"`
	ToDate                              *time.Time `json:"toDate"`
	Estimated                           bool       `json:"estimated"`
}

package hydros

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// NewMeterService creates & initialized new meter service
func NewMeterService(client *Client) MeterService {
	meterService := (&DefaultMeterService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "meters",
			Client:           client,
			PayloadModelType: reflect.TypeOf(MeterModel{}),
		})
	return meterService
}

// MeterService Meter service interface
type MeterService interface {
	Service

	Get(wellID uint, ID uint) (*MeterModel, error)
	ListByWellID(wellID uint) ([]MeterModel, error)
	Create(model *MeterModel) (*MeterModel, error)
	Update(model *MeterModel) (*MeterModel, error)
	Decommission(id uint, decommissionTime time.Time) (*MeterModel, error)
}

// DefaultMeterService default meter service struct that contains backing functions
type DefaultMeterService struct {
	*DefaultService
	GetFunc          func(wellID uint, ID uint) (*MeterModel, error)
	ListByWellIDFunc func(wellID uint) ([]MeterModel, error)
	CreateFunc       func(model *MeterModel) (*MeterModel, error)
	UpdateFunc       func(model *MeterModel) (*MeterModel, error)
	DecommissionFunc func(id uint, decommissionTime time.Time) (*MeterModel, error)
}

// Init initialized spec and default backing functions for service
func (service *DefaultMeterService) Init(spec *ServiceSpec) *DefaultMeterService {
	service.Spec = spec

	// Define Get backing function
	service.GetFunc = func(wellID uint, ID uint) (*MeterModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/%s/%d.json", service.Spec.Client.URL.String(), wellID, service.Spec.ServiceName, ID)
		req, err := http.NewRequest("GET", uri, nil)
		headers := service.Spec.Client.CreateHeadersFunc()
		for h := 0; h < len(headers); h++ {
			req.Header.Add(headers[h].Key, headers[h].Value)
		}

		resp, err := service.Spec.Client.HTTPClient.Do(req)
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

		var meter MeterModel
		err = json.Unmarshal(bodyBytes, &meter)
		if err != nil {
			return nil, err
		}
		return meter.Init(service.Spec), nil
	}

	// Define ListByWellID backing function
	service.ListByWellIDFunc = func(wellID uint) ([]MeterModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/meters.json", service.Spec.Client.URL.String(), wellID)
		req, err := http.NewRequest("GET", uri, nil)
		headers := service.Spec.Client.CreateHeadersFunc()
		for h := 0; h < len(headers); h++ {
			req.Header.Add(headers[h].Key, headers[h].Value)
		}

		resp, err := service.Spec.Client.HTTPClient.Do(req)
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

		var meters []MeterModel
		err = json.Unmarshal(bodyBytes, &meters)
		if err != nil {
			return nil, err
		}
		return meters, nil
	}

	// Define Create backing function
	service.CreateFunc = func(model *MeterModel) (*MeterModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define Update backing function
	service.UpdateFunc = func(model *MeterModel) (*MeterModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define Decommission backing function
	service.DecommissionFunc = func(id uint, decommissionTime time.Time) (*MeterModel, error) {
		return nil, errors.New("not implemented")
	}

	return service
}

// Get Get payload object by id
func (service *DefaultMeterService) Get(wellID uint, ID uint) (*MeterModel, error) {
	return service.GetFunc(wellID, ID)
}

// List Get list of meters for well
func (service *DefaultMeterService) ListByWellID(wellID uint) ([]MeterModel, error) {
	return service.ListByWellIDFunc(wellID)
}

// Create Create new
func (service *DefaultMeterService) Create(model *MeterModel) (*MeterModel, error) {
	return service.CreateFunc(model)
}

// Update Update model
func (service *DefaultMeterService) Update(model *MeterModel) (*MeterModel, error) {
	return service.UpdateFunc(model)
}

// Decommission Decommission model
func (service *DefaultMeterService) Decommission(id uint, decommissionDate time.Time) (*MeterModel, error) {
	return service.DecommissionFunc(id, decommissionDate)
}

package hydros

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// NewMeterReadingService creates initalized new meter reading service
func NewMeterReadingService(client *Client) MeterReadingService {
	meterReadingService := (&DefaultMeterReadingService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "meterReadings",
			Client:           client,
			PayloadModelType: reflect.TypeOf(MeterReadingModel{}),
		})
	return meterReadingService
}

// MeterReadingService Meter Reading service interface
type MeterReadingService interface {
	Service
	Get(wellID uint, meterID uint, ID uint) (*MeterReadingModel, error)
	CountByWell(wellID uint) (int, error)
	CountByWellAndMeter(wellID uint, meterID uint) (int, error)
	ListByWell(wellID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error)
	ListByWellAndMeter(wellID uint, meterID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error)
	GetProductionByWell(wellID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) ([]ProductionModel, error)
	GetProductionByWellAndMeter(wellID uint, meterID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*ProductionModel, error)
}

// DefaultMeterReadingService default meter reading service struct that contains backing functions
type DefaultMeterReadingService struct {
	*DefaultService
	GetFunc                         func(wellID uint, meterID uint, ID uint) (*MeterReadingModel, error)
	CountByWellFunc                 func(wellID uint) (int, error)
	CountByWellAndMeterFunc         func(wellID uint, meterID uint) (int, error)
	ListByWellFunc                  func(wellID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error)
	ListByWellAndMeterFunc          func(wellID uint, meterID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error)
	GetProductionByWellFunc         func(wellID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) ([]ProductionModel, error)
	GetProductionByWellAndMeterFunc func(wellID uint, meterID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*ProductionModel, error)
}

// Init initalized spec and default backing functions for service
func (service *DefaultMeterReadingService) Init(spec *ServiceSpec) *DefaultMeterReadingService {
	service.Spec = spec

	// Define Get backing function
	service.GetFunc = func(wellID uint, meterID uint, ID uint) (*MeterReadingModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/meters/%d/readings/%d.json", service.Spec.Client.URL.String(), wellID, meterID, ID)
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

		var meterReading MeterReadingModel
		err = json.Unmarshal(bodyBytes, &meterReading)
		if err != nil {
			return nil, err
		}
		return meterReading.Init(service.Spec), nil
	}

	// Define CountByWell backing function
	service.CountByWellFunc = func(wellID uint) (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define CountByWellAndMeter backing function
	service.CountByWellAndMeterFunc = func(wellID uint, meterID uint) (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define ListByWell backing function
	service.ListByWellFunc = func(wellID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define ListByWellAndMeter backing function
	service.ListByWellAndMeterFunc = func(wellID uint, meterID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define GetProductionByWell backing function
	service.GetProductionByWellFunc = func(wellID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) ([]ProductionModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/production.json", service.Spec.Client.URL.String(), wellID)
		req, err := http.NewRequest("GET", uri, nil)
		headers := service.Spec.Client.CreateHeadersFunc()
		for h := 0; h < len(headers); h++ {
			req.Header.Add(headers[h].Key, headers[h].Value)
		}

		q := req.URL.Query()
		if fromDate != nil {
			q.Add("fromDate", fromDate.Format("2006-01-02"))
		}

		if toDate != nil {
			q.Add("toDate", toDate.Format("2006-01-02"))
		}

		q.Add("estimateBounds", strconv.FormatBool(estimateBounds))

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

		var production []ProductionModel
		err = json.Unmarshal(bodyBytes, &production)
		if err != nil {
			return nil, err
		}
		return production, nil
	}

	// Define GetProductionByWellAndMeter backing function
	service.GetProductionByWellAndMeterFunc = func(wellID uint, meterID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*ProductionModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/meters/%d/production.json", service.Spec.Client.URL.String(), wellID, meterID)
		req, err := http.NewRequest("GET", uri, nil)
		headers := service.Spec.Client.CreateHeadersFunc()
		for h := 0; h < len(headers); h++ {
			req.Header.Add(headers[h].Key, headers[h].Value)
		}

		q := req.URL.Query()
		if fromDate != nil {
			q.Add("fromDate", fromDate.Format("2006-01-02"))
		}

		if toDate != nil {
			q.Add("toDate", toDate.Format("2006-01-02"))
		}

		q.Add("estimateBounds", strconv.FormatBool(estimateBounds))

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

		var production ProductionModel
		err = json.Unmarshal(bodyBytes, &production)
		if err != nil {
			return nil, err
		}
		return &production, nil
	}

	return service
}

// Get meter reading by id
func (service *DefaultMeterReadingService) Get(wellID uint, meterID uint, ID uint) (*MeterReadingModel, error) {
	return service.GetFunc(wellID, meterID, ID)
}

// Get meter reading count by well id
func (service *DefaultMeterReadingService) CountByWell(wellID uint) (int, error) {
	return service.CountByWellFunc(wellID)
}

// Get meter reading count by well id and meter id
func (service *DefaultMeterReadingService) CountByWellAndMeter(wellID uint, meterID uint) (int, error) {
	return service.CountByWellAndMeterFunc(wellID, meterID)
}

// Get meter readings by well id
func (service *DefaultMeterReadingService) ListByWell(wellID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error) {
	return service.ListByWellFunc(wellID, from, size, sort, startDate, endDate)
}

// Get meter readings by well id and meter id
func (service *DefaultMeterReadingService) ListByWellAndMeter(wellID uint, meterID uint, from int, size int, sort []Sort, startDate *time.Time, endDate *time.Time) ([]MeterReadingModel, error) {
	return service.ListByWellAndMeterFunc(wellID, meterID, from, size, sort, startDate, endDate)
}

// Get production by well id
func (service *DefaultMeterReadingService) GetProductionByWell(wellID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) ([]ProductionModel, error) {
	return service.GetProductionByWellFunc(wellID, fromDate, toDate, estimateBounds)
}

// Get production by well id and meter id
func (service *DefaultMeterReadingService) GetProductionByWellAndMeter(wellID uint, meterID uint, fromDate *time.Time, toDate *time.Time, estimateBounds bool) (*ProductionModel, error) {
	return service.GetProductionByWellAndMeterFunc(wellID, meterID, fromDate, toDate, estimateBounds)
}

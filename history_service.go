package hydros

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// NewHistoryService creates & initialized new history service
func NewHistoryService(client *Client) HistoryService {
	historyService := (&DefaultHistoryService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "history",
			Client:           client,
			PayloadModelType: reflect.TypeOf(HistoryModel{}),
		})
	return historyService
}

// HistoryService History service interface
type HistoryService interface {
	Service

	Get(updateID string) (*HistoryModel, error)
	Count() (int, error)
	List(from int, size int, sort []Sort, updateIds []string, modelType string) ([]*HistoryModel, error)
}

// DefaultHistoryService default history service struct that contains backing functions
type DefaultHistoryService struct {
	*DefaultService
	GetFunc   func(updateID string) (*HistoryModel, error)
	CountFunc func() (int, error)
	ListFunc  func(from int, size int, sort []Sort, updateIds []string, modelType string) ([]*HistoryModel, error)
}

// Init initialized spec and default backing functions for service
func (service *DefaultHistoryService) Init(spec *ServiceSpec) *DefaultHistoryService {
	service.Spec = spec

	// Define GetByUpdateID backing function
	service.GetFunc = func(updateID string) (*HistoryModel, error) {
		uri := fmt.Sprintf("%s/%s/%s.json", service.Spec.Client.URL.String(), service.Spec.ServiceName, updateID)
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

		var history HistoryModel
		err = json.Unmarshal(bodyBytes, &history)
		if err != nil {
			return nil, err
		}
		return history.Init(service.Spec), nil
	}

	// Define Count backing function
	service.CountFunc = func() (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define List backing function
	service.ListFunc = func(from int, size int, sort []Sort, updateIds []string, modelType string) ([]*HistoryModel, error) {
		return nil, errors.New("not implemented")
	}

	return service
}

// Get payload object by id
func (service *DefaultHistoryService) Get(updateID string) (*HistoryModel, error) {
	return service.GetFunc(updateID)
}

// List object for service
func (service *DefaultHistoryService) List(from int, size int, sort []Sort, updateIds []string, modelType string) ([]*HistoryModel, error) {
	return service.ListFunc(from, size, sort, updateIds, modelType)
}

// Count get a total number of objects
func (service *DefaultHistoryService) Count() (int, error) {
	return service.CountFunc()
}

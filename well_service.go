package hydros

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// NewWellService creates & initializes new well service
func NewWellService(client *Client) WellService {

	wellService := (&DefaultWellService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "wells",
			Client:           client,
			PayloadModelType: reflect.TypeOf(WellModel{}),
		})
	return wellService
}

// WellService Well service interface
type WellService interface {
	Service

	Get(ID uint) (*WellModel, error)
	Count() (int, error)
	List(from int, size int, sort []Sort, ids []int) ([]*WellModel, error)
	Create(model *WellModel) (*WellModel, error)
}

// DefaultWellService default well service struct that contains backing functions
type DefaultWellService struct {
	*DefaultService
	GetFunc    func(ID uint) (*WellModel, error)
	CountFunc  func() (int, error)
	ListFunc   func(from int, size int, sort []Sort, ids []int) ([]*WellModel, error)
	CreateFunc func(model *WellModel) (*WellModel, error)
}

// Init Initializes spec and default backing functions for service
func (service *DefaultWellService) Init(spec *ServiceSpec) *DefaultWellService {

	service.Spec = spec

	// Define Get backing function
	service.GetFunc = func(ID uint) (*WellModel, error) {
		uri := fmt.Sprintf("%s/%s/%d.json", service.Spec.Client.URL.String(), service.Spec.ServiceName, ID)
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

		var well WellModel
		err = json.Unmarshal(bodyBytes, &well)
		if err != nil {
			return nil, err
		}
		return &well, nil
	}

	// Define Count backing function
	service.CountFunc = func() (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define List backing function
	service.ListFunc = func(from int, size int, sort []Sort, ids []int) ([]*WellModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define Create backing function
	service.CreateFunc = func(*WellModel) (*WellModel, error) {
		return nil, errors.New("not implemented")
	}

	return service
}

// Get Get payload object by id
func (service *DefaultWellService) Get(ID uint) (*WellModel, error) {
	return service.GetFunc(ID)
}

// List List objects for service
func (service *DefaultWellService) List(from int, size int, sort []Sort, ids []int) ([]*WellModel, error) {
	return service.ListFunc(from, size, sort, ids)
}

// Count Get a total number of objects
func (service *DefaultWellService) Count() (int, error) {
	return service.CountFunc()
}

// Create Create new
func (service *DefaultWellService) Create(model *WellModel) (*WellModel, error) {
	return service.CreateFunc(model)
}

package hydros

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// NewDrillerService creates & initializes new driller service
func NewDrillerService(client *HydrosClient) DrillerService {

	drillerService := (&DefaultDrillerService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "drillers",
			HydrosClient:     client,
			PayloadModelType: reflect.TypeOf(DrillerModel{}),
		})
	return drillerService
}

// DrillerService Driller service interface
type DrillerService interface {
	HydrosService

	Get(ID uint) (*DrillerModel, error)
	Count() (int, error)
	List(from int, size int, sort []Sort, ids []int) ([]*DrillerModel, error)
	Create(model *DrillerModel) (*DrillerModel, error)
}

// DefaultDrillerService default driller service struct that contains backing functions
type DefaultDrillerService struct {
	*DefaultService
	_Get    func(ID uint) (*DrillerModel, error)
	_Count  func() (int, error)
	_List   func(from int, size int, sort []Sort, ids []int) ([]*DrillerModel, error)
	_Create func(model *DrillerModel) (*DrillerModel, error)
}

// Init Initializes spec and default backing functions for service
func (service *DefaultDrillerService) Init(spec *ServiceSpec) *DefaultDrillerService {

	service.Spec = spec

	// Define Get backing function
	service._Get = func(ID uint) (*DrillerModel, error) {
		uri := fmt.Sprintf("%s/%s/%d.json", service.Spec.HydrosClient.URL.String(), service.Spec.ServiceName, ID)
		req, err := http.NewRequest("GET", uri, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", service.Spec.HydrosClient.AccessToken))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")

		resp, err := service.Spec.HydrosClient.HTTPClient.Do(req)
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

		var driller DrillerModel
		err = json.Unmarshal(bodyBytes, &driller)
		if err != nil {
			return nil, err
		}
		return &driller, nil
	}

	// Define Count backing function
	service._Count = func() (int, error) {
		return 0, nil
	}

	// Define List backing function
	service._List = func(from int, size int, sort []Sort, ids []int) ([]*DrillerModel, error) {
		return nil, nil
	}

	// Define Create backing function
	service._Create = func(*DrillerModel) (*DrillerModel, error) {
		return nil, nil
	}

	return service
}

// Get Get payload object by id
func (service *DefaultDrillerService) Get(ID uint) (*DrillerModel, error) {
	return service._Get(ID)
}

// List List objects for service
func (service *DefaultDrillerService) List(from int, size int, sort []Sort, ids []int) ([]*DrillerModel, error) {
	return service._List(from, size, sort, ids)
}

// Count Get a total number of objects
func (service *DefaultDrillerService) Count() (int, error) {
	return service._Count()
}

// Create Create new
func (service *DefaultDrillerService) Create(model *DrillerModel) (*DrillerModel, error) {
	return service._Create(model)
}

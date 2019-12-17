package hydros

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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
	GetWellsByIDs(ids []uint) ([]WellModel, error)
	Count() (int, error)
	List(from int, size int, sort []Sort, ids []uint) ([]*WellModel, error)
	Search(query string, filters []string, from int, size int, sort []Sort) (*WellSearchResults, error)
	Create(model *WellModel) (*WellModel, error)
}

// DefaultWellService default well service struct that contains backing functions
type DefaultWellService struct {
	*DefaultService
	GetFunc           func(ID uint) (*WellModel, error)
	GetWellsByIDsFunc func(ids []uint) ([]WellModel, error)
	CountFunc         func() (int, error)
	ListFunc          func(from int, size int, sort []Sort, ids []uint) ([]*WellModel, error)
	SearchFunc        func(query string, filters []string, from int, size int, sort []Sort) (*WellSearchResults, error)
	CreateFunc        func(model *WellModel) (*WellModel, error)
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

		initializedWell := well.Init(spec)
		return initializedWell, nil
	}

	// Define GetWellsByIDs function
	service.GetWellsByIDsFunc = func(ids []uint) ([]WellModel, error) {
		wellIDsStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")

		jsonStr := []byte(fmt.Sprintf(`{"ids":[%s]}`, wellIDsStr))

		uri := fmt.Sprintf("%s/%s/wellsByIDs.json", service.Spec.Client.URL.String(), service.Spec.ServiceName)
		req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonStr))
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

		var wells []WellModel
		err = json.Unmarshal(bodyBytes, &wells)
		if err != nil {
			return nil, err
		}

		initializedWells := make([]WellModel, len(wells))
		for i := 0; i < len(wells); i++ {
			initializedWells[i] = *wells[i].Init(spec)
		}
		return initializedWells, nil
	}

	// Define Count backing function
	service.CountFunc = func() (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define List backing function
	service.ListFunc = func(from int, size int, sort []Sort, ids []uint) ([]*WellModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define Search backing function
	service.SearchFunc = func(query string, filters []string, from int, size int, sorts []Sort) (*WellSearchResults, error) {

		uri := fmt.Sprintf("%s/%s/search.json", service.Spec.Client.URL.String(), service.Spec.ServiceName)
		req, err := http.NewRequest("GET", uri, nil)
		headers := service.Spec.Client.CreateHeadersFunc()
		for h := 0; h < len(headers); h++ {
			req.Header.Add(headers[h].Key, headers[h].Value)
		}

		q := req.URL.Query()
		if sorts != nil && len(sorts) > 0 {
			var sortStr []string
			for _, sort := range sorts {
				sortStr = append(sortStr, fmt.Sprint(sort.Field, ":", sort.Direction))
			}
			q.Add("sort", strings.Join(sortStr, ","))
		}
		if filters != nil && len(filters) > 0 {
			q.Add("filters", strings.Join(filters, ","))
		}
		q.Add("from", fmt.Sprint(from))
		if size > 150 {
			return nil, errors.New("size parameter must not exceed 150")
		}
		q.Add("size", fmt.Sprint(size))
		req.URL.RawQuery = q.Encode()

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

		var wellSearchResults WellSearchResults
		err = json.Unmarshal(bodyBytes, &wellSearchResults)
		if err != nil {
			return nil, err
		}
		initializedWells := make([]*WellModel, len(wellSearchResults.Results))
		for i := 0; i < len(wellSearchResults.Results); i++ {
			initializedWells[i] = wellSearchResults.Results[i].Init(spec)
		}
		return &WellSearchResults{wellSearchResults.Total, initializedWells}, nil
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

// Get Get wells by ids
func (service *DefaultWellService) GetWellsByIDs(ids []uint) ([]WellModel, error) {
	return service.GetWellsByIDsFunc(ids)
}

// List List objects for service
func (service *DefaultWellService) List(from int, size int, sort []Sort, ids []uint) ([]*WellModel, error) {
	return service.ListFunc(from, size, sort, ids)
}

// Search wells
func (service *DefaultWellService) Search(query string, filters []string, from int, size int, sort []Sort) (*WellSearchResults, error) {
	return service.SearchFunc(query, filters, from, size, sort)
}

// Count Get a total number of objects
func (service *DefaultWellService) Count() (int, error) {
	return service.CountFunc()
}

// Create Create new
func (service *DefaultWellService) Create(model *WellModel) (*WellModel, error) {
	return service.CreateFunc(model)
}

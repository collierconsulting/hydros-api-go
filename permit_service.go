package hydros

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// NewPermitService creates * initialized new permit service
func NewPermitService(client *Client) PermitService {
	permitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).Init(
		&ServiceSpec{
			ServiceName:      "permits",
			Client:           client,
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})
	return permitService
}

// PermitService Permit service interface
type PermitService interface {
	Service
	Get(ID uint) (*PermitModel, error)
	Count() (int, error)
	List(from int, size int, sort []Sort, ids []uint, aggregate bool) ([]*PermitModel, error)
	AmendWellPermits(wellID uint, amendWellPermitsRequest AmendWellPermitsRequest) ([]PermitModel, error)
}

// DefaultPermitService default permit service struct that contains backing functions
type DefaultPermitService struct {
	*DefaultService
	GetFunc              func(ID uint) (*PermitModel, error)
	CountFunc            func() (int, error)
	ListFunc             func(from int, size int, sort []Sort, ids []uint, aggregate bool) ([]*PermitModel, error)
	AmendWellPermitsFunc func(wellID uint, amendWellPermitsRequest AmendWellPermitsRequest) ([]PermitModel, error)
}

// Init initialized spec and default backing functions for service
func (service *DefaultPermitService) Init(spec *ServiceSpec) *DefaultPermitService {
	service.Spec = spec

	// Define Get backing function
	service.GetFunc = func(ID uint) (*PermitModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define Count backing function
	service.CountFunc = func() (int, error) {
		return 0, errors.New("not implemented")
	}

	// Define List backing function
	service.ListFunc = func(from int, size int, sort []Sort, ids []uint, aggregate bool) ([]*PermitModel, error) {
		return nil, errors.New("not implemented")
	}

	// Define AmendWellPermits backing function
	service.AmendWellPermitsFunc = func(wellID uint, amendWellPermitsRequest AmendWellPermitsRequest) ([]PermitModel, error) {
		uri := fmt.Sprintf("%s/wells/%d/%s/amend.json", service.Spec.Client.URL.String(), wellID, service.Spec.ServiceName)
		jsonStr, err := json.Marshal(amendWellPermitsRequest)
		if err != nil {
			return nil, err
		}
		req, err := http.NewRequest("PATCH", uri, bytes.NewBuffer(jsonStr))
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

		if resp.StatusCode != http.StatusAccepted {
			var errorResponse ErrorResponse
			err = json.Unmarshal(bodyBytes, &errorResponse)
			if err == nil && errorResponse.Message != "" {
				return nil, fmt.Errorf("%s: %s", errorResponse.Message, errorResponse.Description)
			}
			return nil, fmt.Errorf("%d error: %s", resp.StatusCode, string(bodyBytes))
		}

		var amendedPermits []PermitModel
		err = json.Unmarshal(bodyBytes, &amendedPermits)
		if err != nil {
			return nil, err
		}
		return amendedPermits, nil
	}

	return service
}

// Get permit by id
func (service *DefaultPermitService) Get(ID uint) (*PermitModel, error) {
	return service.GetFunc(ID)
}

// AmendWellPermits amend well's permits
func (service *DefaultPermitService) List(from int, size int, sort []Sort, ids []uint, aggregate bool) ([]*PermitModel, error) {
	return service.ListFunc(from, size, sort, ids, aggregate)
}

// AmendWellPermits amend well's permits
func (service *DefaultPermitService) Count() (int, error) {
	return service.CountFunc()
}

// AmendWellPermits amend well's permits
func (service *DefaultPermitService) AmendWellPermits(wellID uint, amendWellPermitsRequest AmendWellPermitsRequest) ([]PermitModel, error) {
	return service.AmendWellPermitsFunc(wellID, amendWellPermitsRequest)
}

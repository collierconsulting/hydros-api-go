package hydros

import (
	"reflect"
)

// Service Base interface for services
type Service interface {
	_ServiceSpec() *ServiceSpec
	_SetModelServiceCallMock(string, *ModelServiceCallMock)
}

// SortDirection the sort direction of the query
type SortDirection string

// Sort direction constants
const (
	Asc  SortDirection = "asc"
	Desc SortDirection = "desc"
)

// Sort struct representing sort parameter of search query
type Sort struct {
	Field     string        `json:"field"`
	Direction SortDirection `json:"direction"`
}

// DefaultService Standard service base struct that implements Service
type DefaultService struct {
	Spec *ServiceSpec
}

// ServiceSpec individual service instance metadata
type ServiceSpec struct {
	ServiceName           string
	Client                *Client
	PayloadModelType      reflect.Type
	modelServiceCallMocks map[string]*ModelServiceCallMock
}

// _ServiceSpec Getter for service spec
func (service *DefaultService) _ServiceSpec() *ServiceSpec {
	return service.Spec
}

// _SetModelServiceCallMocks Getter for service spec
func (service *DefaultService) _SetModelServiceCallMock(targetMethodName string, modelServiceCallMocks *ModelServiceCallMock) {
	if service.Spec != nil {
		if service.Spec.modelServiceCallMocks == nil {
			service.Spec.modelServiceCallMocks = make(map[string]*ModelServiceCallMock)
		}
		service.Spec.modelServiceCallMocks[targetMethodName] = modelServiceCallMocks
	}
}

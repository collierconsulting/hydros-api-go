package hydros

import (
	"reflect"
)

// Service Base interface for services
type Service interface {
	_ServiceSpec() *ServiceSpec
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
	ServiceName      string
	Client           *Client
	PayloadModelType reflect.Type
}

// _ServiceSpec Getter for service spec
func (service *DefaultService) _ServiceSpec() *ServiceSpec {
	return service.Spec
}

package hydros

import (
	"errors"
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

	Get(id uint) (*MeterModel, error)
	Create(model *MeterModel) (*MeterModel, error)
	Update(model *MeterModel) (*MeterModel, error)
	Decommission(id uint, decommissionTime time.Time) (*MeterModel, error)
}

// DefaultMeterService default meter service struct that contains backing functions
type DefaultMeterService struct {
	*DefaultService
	GetFunc          func(id uint) (*MeterModel, error)
	CreateFunc       func(model *MeterModel) (*MeterModel, error)
	UpdateFunc       func(model *MeterModel) (*MeterModel, error)
	DecommissionFunc func(id uint, decommissionTime time.Time) (*MeterModel, error)
}

// Init initialized spec and default backing functions for service
func (service *DefaultMeterService) Init(spec *ServiceSpec) *DefaultMeterService {
	service.Spec = spec

	// Define Get backing function
	service.GetFunc = func(id uint) (*MeterModel, error) {
		return nil, errors.New("not implemented")
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
func (service *DefaultMeterService) Get(id uint) (*MeterModel, error) {
	return service.GetFunc(id)
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

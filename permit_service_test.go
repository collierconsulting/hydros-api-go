package hydros

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefaultPermitService_Init(t *testing.T) {
	defaultPermitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultPermitService, "Service should not be nil")
	assert.Equal(t, "test", defaultPermitService.Spec.ServiceName)

	assert.NotNil(t, defaultPermitService.AmendWellPermitsFunc, "AmendWellPermitsFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultPermitService.AmendWellPermitsFunc).Kind(), reflect.Func, "GetFunc should be func")
}

func TestDefaultPermitServiceGetFunc(t *testing.T) {
	id := uint(1)

	defaultPermitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultPermitService.GetFunc = func(ID uint) (*PermitModel, error) {
		return &PermitModel{DefaultModelBase: &DefaultModelBase{ID: ID}}, nil
	}
	returnedModel, err := defaultPermitService.Get(id)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, id, returnedModel.ID)
}

func TestDefaultPermitServiceCountFunc(t *testing.T) {

	defaultPermitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	defaultPermitService.CountFunc = func() (int, error) {
		return 42, nil
	}
	count, err := defaultPermitService.CountFunc()
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, 42, count)
}

func TestDefaultPermitServiceListFunc(t *testing.T) {

	defaultPermitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultPermitService.ListFunc = func(from int, size int, sort []Sort, ids []uint, aggregate bool) ([]*PermitModel, error) {
		list := make([]*PermitModel, 1)
		list[0] = &PermitModel{DefaultModelBase: &DefaultModelBase{ID: 356}}
		return list, nil
	}
	returnedModels, err := defaultPermitService.List(0, 1, nil, nil, false)
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(356), returnedModels[0].ID)
}

func TestDefaultPermitService_AmendWellPermits(t *testing.T) {
	defaultPermitService := (&DefaultPermitService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName: "test",
			PayloadModelType: reflect.TypeOf(PermitModel{}),
		})

	defaultPermitService.AmendWellPermitsFunc = func(wellID uint, amendWellPermitsRequest AmendWellPermitsRequest) ([]PermitModel, error) {
		var permits []PermitModel
		permits = append(permits, PermitModel{DefaultModelBase: &DefaultModelBase{ID: 3}})
		return permits, nil
	}
	returnedModels, err := defaultPermitService.AmendWellPermits(1, AmendWellPermitsRequest{HistoryUpdateID: "test"})
	assert.Nil(t, err, "Error should be nil.")
	assert.Len(t, returnedModels, 1)
	assert.Equal(t, uint(3), returnedModels[0].ID)
}
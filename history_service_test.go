package hydros

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestDefaultHistoryService_Init(t *testing.T) {
	defaultHistoryService := (&DefaultHistoryService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	assert.NotNil(t, defaultHistoryService, "Service should not be nil")
	assert.Equal(t, "test", defaultHistoryService.Spec.ServiceName)

	assert.NotNil(t, defaultHistoryService.GetFunc, "GetFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultHistoryService.GetFunc).Kind(), reflect.Func, "GetFunc should be func")
	assert.NotNil(t, defaultHistoryService.CountFunc, "CountFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultHistoryService.CountFunc).Kind(), reflect.Func, "CountFunc should be func")
	assert.NotNil(t, defaultHistoryService.ListFunc, "ListFunc should not be null")
	assert.Equal(t, reflect.TypeOf(defaultHistoryService.ListFunc).Kind(), reflect.Func, "ListFunc should be func")
}

func TestDefaultHistoryServiceCountFunc(t *testing.T) {

	defaultHistoryService := (&DefaultHistoryService{DefaultService: &DefaultService{}}).Init(&ServiceSpec{ServiceName: "test"})
	defaultHistoryService.CountFunc = func() (int, error) {
		return 42, nil
	}
	count, err := defaultHistoryService.CountFunc()
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, 42, count)
}

func TestDefaultHistoryServiceGetFunc(t *testing.T) {
	uid := uuid.NewV4()

	defaultHistoryService := (&DefaultHistoryService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(HistoryModel{}),
		})

	defaultHistoryService.GetFunc = func(updateID string) (*HistoryModel, error) {
		return &HistoryModel{DefaultModelBase: &DefaultModelBase{ID: uint(1)}, UpdateID: uid.String()}, nil
	}
	returnedModel, err := defaultHistoryService.Get(uid.String())
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(1), returnedModel.ID)
	assert.Equal(t, uid.String(), returnedModel.UpdateID)
}

func TestDefaultHistoryServiceListFunc(t *testing.T) {

	defaultHistoryService := (&DefaultHistoryService{DefaultService: &DefaultService{}}).
		Init(&ServiceSpec{
			ServiceName:      "test",
			PayloadModelType: reflect.TypeOf(HistoryModel{}),
		})

	defaultHistoryService.ListFunc = func(from int, size int, sort []Sort, updateIds []string, modelType string) ([]*HistoryModel, error) {
		list := make([]*HistoryModel, 1)
		list[0] = &HistoryModel{DefaultModelBase: &DefaultModelBase{ID: 356}}
		return list, nil
	}
	returnedModels, err := defaultHistoryService.List(0, 1, nil, nil, "com.test.type")
	assert.Nil(t, err, "Error should be nil.")
	assert.Equal(t, uint(356), returnedModels[0].ID)
}

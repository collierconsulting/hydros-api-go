package hydros

import "time"

// HydrosModel Base interface for service payload models
//type HydrosModel interface {
//	GetID() uint
//	Update(model HydrosModel) (HydrosModel, error)
//	Delete() error
//}

// DefaultModelBase Standard model base struct
type DefaultModelBase struct {
	Spec      *ServiceSpec
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

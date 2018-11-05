package hydros

import "time"

// DefaultModelBase Standard model base struct
type DefaultModelBase struct {
	Spec      *ServiceSpec
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

package hydros

import (
	"time"
)

// CompanyModel Company response payload
type CompanyModel struct {
	*DefaultModelBase
	ID                          uint           `json:"id"`
	CreatedAt                   time.Time      `json:"createdAt"`
	UpdatedAt                   time.Time      `json:"updatedAt"`
	Key                         string         `json:"key"`
	TierID                      uint           `json:"tierId"`
	Name                        string         `json:"name"`
	BrandUrl                    string         `json:"brandUrl"`
	LogoUrl                     string         `json:"logoUrl"`
	Slogan                      string         `json:"slogan"`
	ExportOutOfDistrictFee      float64        `json:"exportOutOfDistrictFee"`
	WaterUsageFee               float64        `json:"waterUsageFee"`
	DefaultProductionTimePeriod int            `json:"defaultProductionTimePeriod"`
	StandardUsersAllowed        uint           `json:"standardUsersAllowed"`
	SuperUsersAllowed           uint           `json:"superUsersAllowed"`
	DefaultLocale               string         `json:"defaultLocale"`
	Countries                   []CountryModel `json:"countries"`
	States                      []StateModel   `json:"states"`
	Counties                    []CountyModel  `json:"counties"`
	CompanyURLs                 []CompanyURL   `json:"urls"`
}

// CompanyURL model
type CompanyURL struct {
	URL     string `json:"url"`
	Default bool   `json:"default"`
}

package hydros

// StateModel model
type StateModel struct {
	ID        uint          `json:"id"`
	CountryID uint          `json:"countryId"`
	Fips      string        `json:"fips"`
	Iso31662  string        `json:"iso31662"`
	Name      string        `json:"name"`
	AltName   string        `json:"altName"`
	Region    string        `json:"region"`
	Postal    string        `json:"postal"`
	Latitude  float64       `json:"latitude"`
	Longitude float64       `json:"longitude"`
	RegionSub string        `json:"regionSub"`
	Counties  []CountyModel `json:"counties"`
}

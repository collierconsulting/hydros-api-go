package hydros

// type CountyModel struct { model
type CountyModel struct {
	ID   uint   `json:"id"`
	Fips string `json:"fips"`
	Name string `json:"name"`
}

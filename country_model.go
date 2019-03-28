package hydros

// CountryModel model
type CountryModel struct {
	ID             uint         `json:"id"`
	Fips           string       `json:"fips"`
	Iso3166Alpha2  string       `json:"iso3166Alpha2"`
	Iso3166Alpha3  string       `json:"iso3166Alpha3"`
	Iso3166Numeric string       `json:"iso3166Numeric3"`
	Name           string       `json:"name"`
	NameLong       string       `json:"nameLong"`
	Abbreviation   string       `json:"abbreviation"`
	Postal         string       `json:"postal"`
	FormalEN       string       `json:"formalEn"`
	Continent      string       `json:"continent"`
	RegionUN       string       `json:"regionUn"`
	SubRegion      string       `json:"subRegion"`
	States         []StateModel `json:"states"`
}

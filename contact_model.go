package hydros

import "gopkg.in/guregu/null.v3"

// ContactModel Contact response payload
type ContactModel struct {
	*DefaultModelBase
	FirstName      null.String                `json:"firstName"`
	LastName       null.String                `json:"lastName"`
	CompanyName    null.String                `json:"companyName"`
	Email          null.String                `json:"email"`
	Address1       null.String                `json:"address1"`
	Address2       null.String                `json:"address2"`
	City           null.String                `json:"city"`
	State          null.String                `json:"state"`
	PostalCode     null.String                `json:"postalCode"`
	Classification null.String                `json:"classification"`
	PhoneNumbers   []*ContactPhoneNumberModel `json:"phoneNumbers,omitempty"`
}

// ContactPhoneNumberModel phone number model for contact association
type ContactPhoneNumberModel struct {
	*PhoneNumberModel
}

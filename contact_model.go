package hydros

// ContactModel Contact response payload
type ContactModel struct {
	*DefaultModelBase
	CompanyID      uint                       `json:"companyId"`
	FirstName      *string                    `json:"firstName,omitempty"`
	LastName       *string                    `json:"lastName,omitempty"`
	CompanyName    *string                    `json:"companyName,omitempty"`
	Email          *string                    `json:"email,omitempty"`
	Address1       *string                    `json:"address1,omitempty"`
	Address2       *string                    `json:"address2,omitempty"`
	City           *string                    `json:"city,omitempty"`
	State          *string                    `json:"state,omitempty"`
	PostalCode     *string                    `json:"postalCode,omitempty"`
	Classification *string                    `json:"classification,omitempty"`
	PhoneNumbers   []*ContactPhoneNumberModel `json:"phoneNumbers,omitempty"`
}

// ContactPhoneNumberModel phone number model for contact association
type ContactPhoneNumberModel struct {
	*PhoneNumberModel
}

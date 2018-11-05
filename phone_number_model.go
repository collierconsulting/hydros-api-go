package client

// PhoneNumberModel phone number payload model
type PhoneNumberModel struct {
	*DefaultModelBase
	Type        string    `json:"numberType"`
	PhoneNumber string    `json:"phoneNumber"`
	Primary     bool      `json:"isPrimary"`
}

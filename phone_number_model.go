package hydros

import "gopkg.in/guregu/null.v3"

// PhoneNumberModel phone number payload model
type PhoneNumberModel struct {
	*DefaultModelBase
	Type        null.String `json:"numberType"`
	PhoneNumber null.String `json:"phoneNumber"`
	Primary     null.Bool   `json:"isPrimary"`
}

package data

// OTPData represents the data structure for sending an OTP
type OTPData struct {
	PhoneNumber string `json:"phoneNumber,omitempty" validate:"required"`
	// PhoneNumber: the recipient's phone number. Marked as required and mapped to the JSON field "phoneNumber".
}

// VerifyData represents the data structure for verifying an OTP
type VerifyData struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	// User: a reference to OTPData that includes phone number details. Marked as required and mapped to the JSON field "user".

	Code string `json:"code,omitempty" validate:"required"`
	// Code: the OTP code entered by the user for verification. Marked as required and mapped to the JSON field "code".
}

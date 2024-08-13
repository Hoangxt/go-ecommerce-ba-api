package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken      = 30001 // Invalid token
	ErrCodeUserHasExists = 50002 // User has exist
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid token",

	ErrCodeUserHasExists: "User has already registered",
}

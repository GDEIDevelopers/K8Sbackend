package errhandle

type ErrCode int

//go:generate stringer -type=ErrCode
const (
	NoError ErrCode = iota
	InnerError
	TokenError
	ParamsError
	UserExists
	UserNonExists
	PasswordInvalid
	PermissionDenied
)

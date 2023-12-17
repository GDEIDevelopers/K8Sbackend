// Code generated by "stringer -type=ErrCode"; DO NOT EDIT.

package errhandle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoError-0]
	_ = x[InnerError-1]
	_ = x[TokenError-2]
	_ = x[UserExists-3]
	_ = x[UserNonExists-4]
	_ = x[PasswordInvalid-5]
	_ = x[PermissionDenied-6]
}

const _ErrCode_name = "NoErrorInnerErrorTokenErrorUserExistsUserNonExistsPasswordInvalidPermissionDenied"

var _ErrCode_index = [...]uint8{0, 7, 17, 27, 37, 50, 65, 81}

func (i ErrCode) String() string {
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}

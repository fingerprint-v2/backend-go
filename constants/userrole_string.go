// Code generated by "stringer -type=UserRole"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SUPERADMIN-1]
	_ = x[ADMIN-2]
	_ = x[USER-3]
}

const _UserRole_name = "SUPERADMINADMINUSER"

var _UserRole_index = [...]uint8{0, 10, 15, 19}

func (i UserRole) String() string {
	i -= 1
	if i < 0 || i >= UserRole(len(_UserRole_index)-1) {
		return "UserRole(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _UserRole_name[_UserRole_index[i]:_UserRole_index[i+1]]
}

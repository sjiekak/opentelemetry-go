// Code generated by "stringer -type=ValueKind"; DO NOT EDIT.

package metric

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Int64ValueKind-0]
	_ = x[Float64ValueKind-1]
}

const _ValueKind_name = "Int64ValueKindFloat64ValueKind"

var _ValueKind_index = [...]uint8{0, 14, 30}

func (i ValueKind) String() string {
	if i < 0 || i >= ValueKind(len(_ValueKind_index)-1) {
		return "ValueKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ValueKind_name[_ValueKind_index[i]:_ValueKind_index[i+1]]
}
// Code generated by "stringer -type=FormatVersion"; DO NOT EDIT

package sql

import "fmt"

const _FormatVersion_name = "BaseFormatVersion"

var _FormatVersion_index = [...]uint8{0, 17}

func (i FormatVersion) String() string {
	i -= 1
	if i >= FormatVersion(len(_FormatVersion_index)-1) {
		return fmt.Sprintf("FormatVersion(%d)", i+1)
	}
	return _FormatVersion_name[_FormatVersion_index[i]:_FormatVersion_index[i+1]]
}

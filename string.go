package cloudformation

import (
	"encoding/json"
	"strings"
)

const rawMarker = "\xef\xbb\xbf"

type String string

func (s String) MarshalJSON() ([]byte, error) {
	if strings.HasPrefix(string(s), rawMarker) {
		return []byte(strings.TrimPrefix(string(s), rawMarker)), nil
	}
	return json.Marshal(string(s))
}

func rawString(v interface{}) String {
	buf, _ := json.Marshal(&v)
	return String(rawMarker + String(buf))
}

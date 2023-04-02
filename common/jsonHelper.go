package common

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Scan for scanner helper
func Scan(data interface{}, value interface{}) error {
	if value == nil {
		return nil
	}

	switch value.(type) {
	case []byte:
		return json.Unmarshal(value.([]byte), &data)
	case string:
		return json.Unmarshal([]byte(value.(string)), &data)
	default:
		return fmt.Errorf("val type is valid, is %+v", value)
	}
}

// Value for valuer helper
func Value(data interface{}) (interface{}, error) {
	vi := reflect.ValueOf(data)
	// 判断是否为 0 值
	if vi.IsZero() {
		return nil, nil
	}

	return json.Marshal(data)
}

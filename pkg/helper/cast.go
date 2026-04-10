package helper

import (
	"strconv"
	"time"
)

func ToUint(v any) uint {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int32:
		return uint(val)
	case uint32:
		return uint(val)
	case int64:
		return uint(val)
	case uint64:
		return uint(val)
	case float64:
		return uint(val)
	case []byte:
		n, _ := strconv.ParseUint(string(val), 10, 64)
		return uint(n)
	case string:
		n, _ := strconv.ParseUint(val, 10, 64)
		return uint(n)
	}
	return 0
}

func ToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	}
	return ""
}

func ToStringPtr(v any) *string {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case string:
		return &val
	case []byte:
		s := string(val)
		return &s
	}
	return nil
}

func ToIntPtr(v any) *int {
	if v == nil {
		return nil
	}
	switch val := v.(type) {
	case int32:
		i := int(val)
		return &i
	case uint32:
		i := int(val)
		return &i
	case int64:
		i := int(val)
		return &i
	case float64:
		i := int(val)
		return &i
	case []byte:
		n, err := strconv.Atoi(string(val))
		if err != nil {
			return nil
		}
		return &n
	case string:
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil
		}
		return &n
	}
	return nil
}

func ToTimePtr(v any) *time.Time {
	if v == nil {
		return nil
	}
	t, ok := v.(time.Time)
	if !ok {
		return nil
	}
	return &t
}

func ToBool(v any) bool {
	if v == nil {
		return false
	}
	switch val := v.(type) {
	case int32:
		return val == 1
	case int64:
		return val == 1
	case bool:
		return val
	case []byte:
		return string(val) == "1"
	case string:
		return val == "1"
	}
	return false
}

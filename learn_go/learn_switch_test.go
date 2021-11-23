package learn_go

import (
	"strconv"
	"testing"
)

func Test_Switch(t *testing.T) {
	var v interface{}
	ToStringAny(v)
	ToInt32Any(v)
	ToInt64Any(v)
	ToFloat64Any(v)
	ToBoolAny(v)

	var str string
	ToInt(str)
	ToInt32(str)
	ToInt64(str)
	ToBool(str)
	ToFloat64(str)
}

func ToStringAny(e interface{}) string {
	res := ""
	if e == nil {
		return res
	}
	switch v := e.(type) {
	case string:
		res = e.(string)
		break
	case int:
		res = strconv.FormatInt(int64(v), 10)
		break
	case int32:
		res = strconv.FormatInt(int64(v), 10)
		break
	case int64:
		res = strconv.FormatInt(v, 10)
		break
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', -1, 64)
		break
	case float64:
		res = strconv.FormatFloat(v, 'f', -1, 64)
		break
	}
	return res
}

// ToInt32Any converts an interface{} to a base 10 32-bit integer or return 0
func ToInt32Any(v interface{}) int32 {
	if v, ok := v.(int); ok {
		return int32(v)
	}
	if v, ok := v.(int32); ok {
		return v
	}
	if v, ok := v.(int64); ok {
		return int32(v)
	}
	if v, ok := v.(float32); ok {
		return int32(v)
	}
	if v, ok := v.(float64); ok {
		return int32(v)
	}
	if v, ok := v.(*int); ok {
		if v != nil {
			return int32(*v)
		}
	}
	if v, ok := v.(*int32); ok {
		if v != nil {
			return *v
		}
	}
	if v, ok := v.(*int64); ok {
		if v != nil {
			return int32(*v)
		}
	}
	if v, ok := v.(string); ok {
		return ToInt32(v)
	}
	if v, ok := v.(*string); ok {
		if v != nil {
			return ToInt32(*v)
		}
	}
	return 0
}

// ToInt64Any converts an interface{} to a base 10 64-bit integer or return 0
func ToInt64Any(v interface{}) int64 {
	if v, ok := v.(int); ok {
		return int64(v)
	}
	if v, ok := v.(int32); ok {
		return int64(v)
	}
	if v, ok := v.(int64); ok {
		return v
	}
	if v, ok := v.(float32); ok {
		return int64(v)
	}
	if v, ok := v.(float64); ok {
		return int64(v)
	}
	if v, ok := v.(*int); ok {
		if v != nil {
			return int64(*v)
		}
	}
	if v, ok := v.(*int32); ok {
		if v != nil {
			return int64(*v)
		}
	}
	if v, ok := v.(*int64); ok {
		if v != nil {
			return int64(*v)
		}
	}
	if v, ok := v.(string); ok {
		return ToInt64(v)
	}
	if v, ok := v.(*string); ok {
		if v != nil {
			return ToInt64(*v)
		}
	}
	return 0
}

// ToFloat64Any converts an interface{} to a 64-bit float or return 0
func ToFloat64Any(v interface{}) float64 {
	if v, ok := v.(int); ok {
		return float64(v)
	}
	if v, ok := v.(int32); ok {
		return float64(v)
	}
	if v, ok := v.(int64); ok {
		return float64(v)
	}
	if v, ok := v.(float32); ok {
		return float64(v)
	}
	if v, ok := v.(float64); ok {
		return v
	}
	if v, ok := v.(*int); ok {
		if v != nil {
			return float64(*v)
		}
	}
	if v, ok := v.(*int32); ok {
		if v != nil {
			return float64(*v)
		}
	}
	if v, ok := v.(*int64); ok {
		if v != nil {
			return float64(*v)
		}
	}
	if v, ok := v.(string); ok {
		return ToFloat64(v)
	}
	if v, ok := v.(*string); ok {
		if v != nil {
			return ToFloat64(*v)
		}
	}
	return 0.0
}

// ToBoolAny returns bool for interface{} and if parse error, return false
func ToBoolAny(v interface{}) bool {
	if v, ok := v.(bool); ok {
		return v
	}
	if v, ok := v.(*bool); ok {
		if v != nil {
			return *v
		}
		return false
	}
	if v, ok := v.(int); ok {
		return v > 0
	}
	if v, ok := v.(int32); ok {
		return v > 0
	}
	if v, ok := v.(int64); ok {
		return v > 0
	}
	if v, ok := v.(float32); ok {
		return v > 0
	}
	if v, ok := v.(float64); ok {
		return v > 0
	}
	if v, ok := v.(*int); ok {
		if v != nil {
			return *v > 0
		}
	}
	if v, ok := v.(*int32); ok {
		if v != nil {
			return *v > 0
		}
	}
	if v, ok := v.(*int64); ok {
		if v != nil {
			return *v > 0
		}
	}
	if v, ok := v.(string); ok {
		return ToBool(v)
	}
	if v, ok := v.(*string); ok {
		if v != nil {
			return ToBool(*v)
		}
	}
	return false
}

// ToInt32 converts a string to a base 10 32-bit integer or return 0
func ToInt32(v string) int32 {
	if v, err := strconv.ParseInt(v, 10, 32); err == nil {
		return int32(v)
	}
	return 0
}

func ToInt(v string) int {
	if v, err := strconv.ParseInt(v, 10, 32); err == nil {
		return int(v)
	}
	return 0
}

// ToInt64 converts a string to a base 10 64-bit integer or return 0
func ToInt64(v string) int64 {
	if v, err := strconv.ParseInt(v, 10, 64); err == nil {
		return v
	}
	return 0
}

// ToBool returns bool for string and if parse error, return false
func ToBool(v string) bool {
	if ok, err := strconv.ParseBool(v); err == nil {
		return ok
	}
	return false
}

// ToFloat64 converts a string to a 64-bit float or return 0
func ToFloat64(v string) float64 {
	if v, err := strconv.ParseFloat(v, 64); err == nil {
		return v
	}
	return 0
}
package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Merge merge second into self,if the key is the same then the new value replaces the old value.
func (cli *MapStr) Merge(second MapStr) {
	for key, val := range second {
		(*cli)[key] = val
	}
}

// ToJSON convert to json string
func (cli *MapStr) ToJSON() []byte {
	js, _ := json.Marshal(cli)
	return js
}

// Bool get the value as bool
func (cli *MapStr) Bool(key string) bool {
	switch t := (*cli)[key].(type) {
	case nil:
		return false
	default:
		return false
	case bool:
		return t
	}
}

// Int return the value by the key
func (cli *MapStr) Int(key string) (int, error) {

	switch t := (*cli)[key].(type) {
	default:
		return 0, errors.New("invalid num")
	case nil:
		return 0, errors.New("invalid key, not found value")
	case int:
		return t, nil
	case int16:
		return int(t), nil
	case int32:
		return int(t), nil
	case int64:
		return int(t), nil
	case float32:
		return int(t), nil
	case float64:
		return int(t), nil
	case json.Number:
		num, err := t.Int64()
		return int(num), err
	case string:
		return strconv.Atoi(t)
	}
}

// Float get the value as float64
func (cli *MapStr) Float(key string) (float64, error) {
	switch t := (*cli)[key].(type) {
	default:
		return 0, errors.New("invalid num")
	case nil:
		return 0, errors.New("invalid key, not found value")
	case int:
		return float64(t), nil
	case int16:
		return float64(t), nil
	case int32:
		return float64(t), nil
	case int64:
		return float64(t), nil
	case float32:
		return float64(t), nil
	case float64:
		return t, nil
	case json.Number:
		return t.Float64()
	}
}

// String get the value as string
func (cli *MapStr) String(key string) string {
	switch t := (*cli)[key].(type) {
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", t)
	case map[string]interface{}, []interface{}, interface{}:
		rest, _ := json.Marshal(t)
		return string(rest)
	case json.Number:
		return t.String()
	case string:
		return t
	}
}

// Time get the value as time.Time
func (cli *MapStr) Time(key string) (*time.Time, error) {
	switch t := (*cli)[key].(type) {
	default:
		return nil, errors.New("invalid time value")
	case nil:
		return nil, errors.New("invalid key")
	case time.Time:
		return &t, nil
	case *time.Time:
		return t, nil
	case string:
		if tm, tmErr := time.Parse(time.RFC1123, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC1123Z, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC3339, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC3339Nano, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC822, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC822Z, t); nil == tmErr {
			return &tm, nil
		}

		if tm, tmErr := time.Parse(time.RFC850, t); nil == tmErr {
			return &tm, nil
		}

		return nil, errors.New("can not parse the datetime")
	}
}

// MapStr get the MapStr object
func (cli *MapStr) MapStr(key string) (MapStr, error) {

	switch t := (*cli)[key].(type) {
	default:
		return nil, errors.New("the data is not a map[string]interface{} type")
	case nil:
		return nil, errors.New("the key is invalid")
	case map[string]interface{}:
		return MapStr(t), nil
	}

}

// MapStrArray get the MapStr object array
func (cli *MapStr) MapStrArray(key string) ([]MapStr, error) {

	switch t := (*cli)[key].(type) {
	default:
		return nil, errors.New("the data is not a map[string]interface{} type")
	case nil:
		return nil, errors.New("the key is invalid")
	case []map[string]interface{}:
		items := make([]MapStr, 0)
		for _, item := range t {
			items = append(items, item)
		}
		return items, nil
	case []interface{}:
		items := make([]MapStr, 0)
		for _, item := range t {
			switch childType := item.(type) {
			case map[string]interface{}:
				items = append(items, childType)
			}
		}
		return items, nil
	}

}
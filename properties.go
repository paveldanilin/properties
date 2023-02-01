package properties

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrPropertyNotFound = errors.New("property not found")
)

type Properties struct {
	props map[string]string
}

// Creates a new properties struct from key/value pairs
func New(pairs ...string) *Properties {
	pairsNum := len(pairs)
	if pairsNum > 0 && pairsNum%2 != 0 {
		// mismatch key-value pairs
		pairs = append(pairs, "")
	}

	pairsNum = len(pairs)
	props := make(map[string]string)
	for i := 0; i < pairsNum; i = i + 2 {
		key := pairs[i]
		value := pairs[i+1]
		props[key] = value
	}

	return &Properties{
		props: props,
	}
}

func (props *Properties) SetProperty(key string, value string) {
	props.props[key] = value
}

func (props *Properties) GetProperty(key string) (string, error) {
	if v, ok := props.props[key]; ok {
		return v, nil
	}
	return "", ErrPropertyNotFound
}

func (props *Properties) Property(key string, def string) string {
	if props.HasProperty(key) {
		return props.props[key]
	}
	return def
}

func (props *Properties) MustString(key string) string {
	return props.Property(key, "")
}

func (props *Properties) String(key string, def string) string {
	return props.Property(key, def)
}

func (props *Properties) GetBool(key string) (bool, error) {
	v, err := props.GetProperty(key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(v)
}

func (props *Properties) MustBool(key string) bool {
	v, err := props.GetBool(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Bool(key string, def bool) bool {
	v, err := props.GetBool(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetInt(key string) (int, error) {
	v, err := props.GetProperty(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(v)
}

func (props *Properties) MustInt(key string) int {
	v, err := props.GetInt(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Int(key string, def int) int {
	v, err := props.GetInt(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetFloat(key string) (float64, error) {
	v, err := props.GetProperty(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(v, 64)
}

func (props *Properties) MustFloat(key string) float64 {
	v, err := props.GetFloat(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Float(key string, def float64) float64 {
	v, err := props.GetFloat(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetUint(key string) (uint64, error) {
	v, err := props.GetProperty(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(v, 10, 64)
}

func (props *Properties) MustUint(key string) uint64 {
	v, err := props.GetUint(key)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Uint(key string, def uint64) uint64 {
	v, err := props.GetUint(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) HasProperty(key string) bool {
	_, ok := props.props[key]
	return ok
}

// Keys returns all property keys
func (props *Properties) Keys() []string {
	keys := make([]string, len(props.props))
	for k := range props.props {
		keys = append(keys, k)
	}
	return keys
}

// KeysWithPrefix returns an array of property keys that start with prefix
func (props *Properties) KeysWithPrefix(prefix string) []string {
	keys := []string{}
	for k := range props.props {
		if strings.HasPrefix(k, prefix) {
			keys = append(keys, k)
		}
	}
	return keys
}

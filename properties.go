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

// Creates a new properties struct from a list of key/value pairs
func New(pairs ...string) *Properties {
	pairsNum := len(pairs)
	if pairsNum > 0 && pairsNum%2 != 0 {
		// mismatch key-value pairs, just add empty value
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

// Creates a new properties struct from a map
func NewFromMap(props map[string]string) *Properties {
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

func (props *Properties) Bool(key string, def bool) (bool, error) {
	v, err := props.GetBool(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def, nil
		}
		return false, err
	}
	return v, nil
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

func (props *Properties) Int(key string, def int) (int, error) {
	v, err := props.GetInt(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def, nil
		}
		return 0, err
	}
	return v, nil
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

func (props *Properties) Float(key string, def float64) (float64, error) {
	v, err := props.GetFloat(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def, nil
		}
		return 0, err
	}
	return v, nil
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

func (props *Properties) Uint(key string, def uint64) (uint64, error) {
	v, err := props.GetUint(key)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def, nil
		}
		return 0, err
	}
	return v, nil
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

func (props *Properties) All() map[string]string {
	return props.props
}

// GetWithPrefix returns a new property struct that holds all properties with a specified prefix
func (props *Properties) GetWithPrefix(prefix string) *Properties {
	keys := props.KeysWithPrefix(prefix)
	p := New()
	for _, k := range keys {
		p.SetProperty(k, props.props[k])
	}
	return p
}

func (props *Properties) Size() int {
	return len(props.props)
}

func (props *Properties) IsEmpty() bool {
	return props.Size() == 0
}

func (props *Properties) IsNotEmpty() bool {
	return props.Size() != 0
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

// RemoveProperty removes property by a given key
func (props *Properties) RemoveProperty(key string) bool {
	if props.HasProperty(key) {
		delete(props.props, key)
		return true
	}
	return false
}

// Contains returns TRUE if all keys are present, otherwise returns FALSE
// If keys size is 0, returns FALSE
func (props *Properties) Contains(keys []string) bool {
	if len(keys) == 0 {
		return false
	}
	for _, k := range keys {
		if !props.HasProperty(k) {
			return false
		}
	}
	return true
}

// ContainsAny returns TRUE if at least one key is present, otherwise returns FALSE
func (props *Properties) ContainsAny(keys []string) bool {
	if len(keys) == 0 {
		return false
	}
	for _, k := range keys {
		if props.HasProperty(k) {
			return true
		}
	}
	return false
}

// Merge merges two property structs into a new one
func (props *Properties) Merge(properties *Properties, overwriteSameKeys bool) *Properties {
	newProps := NewFromMap(props.All())
	for k, v := range properties.All() {
		if newProps.HasProperty(k) {
			if overwriteSameKeys {
				newProps.SetProperty(k, v)
			}
		} else {
			newProps.SetProperty(k, v)
		}
	}
	return newProps
}

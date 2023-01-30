package properties

import (
	"errors"
	"strconv"
)

var (
	ErrPropertyNotFound = errors.New("property not found")
)

type Properties struct {
	props map[string]string
}

func New(pairs ...string) *Properties {
	pairsNum := len(pairs)
	if pairsNum > 0 && pairsNum%2 != 0 {
		// mismatch key-value pairs
		pairs = append(pairs, "")
	}

	pairsNum = len(pairs)
	props := make(map[string]string)
	for i := 0; i < pairsNum; i = i + 2 {
		name := pairs[i]
		value := pairs[i+1]
		props[name] = value
	}

	return &Properties{
		props: props,
	}
}

func (props *Properties) SetProperty(name string, value string) {
	props.props[name] = value
}

func (props *Properties) GetProperty(name string) (string, error) {
	if v, ok := props.props[name]; ok {
		return v, nil
	}
	return "", ErrPropertyNotFound
}

func (props *Properties) Property(name string, def string) string {
	if props.HasProperty(name) {
		return props.props[name]
	}
	return def
}

func (props *Properties) MustString(name string) string {
	return props.Property(name, "")
}

func (props *Properties) String(name string, def string) string {
	return props.Property(name, def)
}

func (props *Properties) GetBool(name string) (bool, error) {
	v, err := props.GetProperty(name)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(v)
}

func (props *Properties) MustBool(name string) bool {
	v, err := props.GetBool(name)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Bool(name string, def bool) bool {
	v, err := props.GetBool(name)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetInt(name string) (int, error) {
	v, err := props.GetProperty(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(v)
}

func (props *Properties) MustInt(name string) int {
	v, err := props.GetInt(name)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Int(name string, def int) int {
	v, err := props.GetInt(name)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetFloat(name string) (float64, error) {
	v, err := props.GetProperty(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(v, 64)
}

func (props *Properties) MustFloat(name string) float64 {
	v, err := props.GetFloat(name)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Float(name string, def float64) float64 {
	v, err := props.GetFloat(name)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) GetUint(name string) (uint64, error) {
	v, err := props.GetProperty(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(v, 10, 64)
}

func (props *Properties) MustUint(name string) uint64 {
	v, err := props.GetUint(name)
	if err != nil {
		panic(err)
	}
	return v
}

func (props *Properties) Uint(name string, def uint64) uint64 {
	v, err := props.GetUint(name)
	if err != nil {
		if errors.Is(err, ErrPropertyNotFound) {
			return def
		}
		panic(err)
	}
	return v
}

func (props *Properties) HasProperty(name string) bool {
	_, ok := props.props[name]
	return ok
}

func (props *Properties) Keys() []string {
	keys := make([]string, len(props.props))
	for k := range props.props {
		keys = append(keys, k)
	}
	return keys
}

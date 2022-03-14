package json

import (
	"encoding/json"
	"io"
	"reflect"
)

type DiffType int64

const (
	Match DiffType = iota
	TypeMismatch
	PrimitiveValueMismatch
	ObjectMismatch
)

func (dt DiffType) String() string {
	switch dt {
	case Match:
		return "Match"
	case TypeMismatch:
		return "TypeMismatch"
	case PrimitiveValueMismatch:
		return "PrimitiveValueMismatch"
	case ObjectMismatch:
		return "ObjectMismatch"
	}
	return "unknown"
}

func compare(a, b interface{}) (result DiffType) {
	typeA := reflect.TypeOf(a)
	typeB := reflect.TypeOf(b)
	if typeA != typeB {
		result = TypeMismatch
		return
	}
	switch typeA.Kind() {
	case reflect.Float64:
	case reflect.String:
		if a != b {
			result = PrimitiveValueMismatch
			return
		}
	case reflect.Map:
		for k1, v1 := range a.(map[string]interface{}) {
			v2 := b.(map[string]interface{})[k1]
			compare(v1, v2)
		}
		return
	}
	return
}

func Compare(a, b io.Reader) (result DiffType) {
	resultA, err := decodeOne(a)
	if err != nil {
		panic(err)
	}
	resultB, err := decodeOne(b)
	if err != nil {
		panic(err)
	}
	result = compare(resultA, resultB)
	return
}

func decodeOne(r io.Reader) (result interface{}, err error) {
	dec := json.NewDecoder(r)
	err = dec.Decode(&result)
	return
}

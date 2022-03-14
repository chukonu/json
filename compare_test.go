package json

import (
	"strings"
	"testing"
)

func TestCompareWithMismatchingTypes(t *testing.T) {
	result := Compare(strings.NewReader(`1`), strings.NewReader(`"1"`))
	if result != TypeMismatch {
		t.Fatalf(`Compare 1 with "1" = %q, want %#q`, result, TypeMismatch)
	}
}

func TestCompareWithDifferentStrings(t *testing.T) {
	result := Compare(strings.NewReader(`"a"`), strings.NewReader(`"b"`))
	if result != PrimitiveValueMismatch {
		t.Fatalf(`Compare "a" with "b" = %q, want %#q`, result, PrimitiveValueMismatch)
	}
}

func TestCompareWithIdenticalStrings(t *testing.T) {
	result := Compare(strings.NewReader(`"a"`), strings.NewReader(`"a"`))
	if result != Match {
		t.Fatalf(`Compare "a" with "a" = %q, want %#q`, result, Match)
	}
}

func TestCompareWithIdenticalObjects(t *testing.T) {
	result := Compare(strings.NewReader(`{"a":0}`), strings.NewReader(`{"a":0}`))
	if result != Match {
		t.Fatalf(`Compare {"a":0} with {"a":0} = %q, want %#q`, result, Match)
	}
}

func TestCompareWithDifferentObjects(t *testing.T) {
	result := Compare(strings.NewReader(`{"a":0}`), strings.NewReader(`{"a":1}`))
	if result != ObjectMismatch {
		t.Fatalf(`Compare {"a":0} with {"a":1} = %q, want %#q`, result, ObjectMismatch)
	}
}
